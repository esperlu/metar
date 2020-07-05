// metar is a go program to fetch the aviation METAR's and TAF's for a given list of airports in console (terminal) mode.
//
// Usage:
// Retrieve messages for a list of stations (IATA or ICAO codes):
//		$ metar lhr jfk bru uudd
// Find the IATA/ICAO airport code for an airport
//		$ metar -s munich
//		$ metar -s new york
// Help screen for other options:
//		$ metar -h
// Bug reports:
// https://github.com/esperlu/metar/issues

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/esperlu/metar/data"
)

// Constants to fetch Weather reports from aviationweather.com
const (
	URL         = "http://aviationweather.gov/adds/dataserver_current/httpparam?dataSource=%s&requestType=retrieve&format=csv&stationString=%s"
	urlMETARfmt = URL + "&hoursBeforeNow=%.1f"
	urlTAFfmt   = URL + "&hoursBeforeNow=%.1f&mostRecentForEachStation=true&Fields=raw_text"
	maxNbMETAR  = 70
	maxTIMEOUT  = 10
)

// flag variables
var numberMetarFlagInt = flag.Int("n", 4, "Set number of Metars to print per station. N 1 to 30.")
var searchFlagBool = flag.Bool("s", false, "Search IATA/ICAO code for an airport")
var timeoutFlagInt = flag.Int("t", 2, "Change the default timeout of 2 sec. to a maximum of 10 sec.")
var rawFlagBool = flag.Bool("r", false, "Print raw data w/o the additional factors")
var metarOnlyFlagBool = flag.Bool("m", false, "METARs only")

// init flags to make them available to func's
func init() {
	flag.Usage = func() {
		fmt.Println(data.Help)
	}
	flag.Parse()
}

func main() {

	startTotal := time.Now()

	// If option search but no search pattern
	if *searchFlagBool && len(flag.Args()) == 0 {
		fmt.Printf("\n\tAsked for search, but no search pattern given. Quitting...\n\tTry: metar -s munich or metar -s mun\n\n")
		return

		// No args -> help screen
	} else if len(flag.Args()) == 0 {
		fmt.Printf("\n\tNo airport(s) given. Quitting...\n\tTry: metar MUC ebbr JFK (more examples below)\n\n")
		fmt.Println(data.Help)
		return
	}

	// validate number of reports (INT)
	if *numberMetarFlagInt <= 0 || *numberMetarFlagInt > maxNbMETAR {
		fmt.Printf(
			"\n\t%s\n\t%s\n\n",
			"Invalid value for option -n (number of METARs).",
			"Minimum 1, maximum 70",
		)
		return
	}

	// validate timeout flag (INT)
	if *timeoutFlagInt < 1 || *timeoutFlagInt > maxTIMEOUT {
		fmt.Printf(
			"\n\t%s\n\t%s\n\n",
			"Invalid value for option -t (connection timeout).",
			"Minimum 1, maximum 10",
		)
		return
	}

	// search option -s
	if *searchFlagBool {
		fmt.Printf("\n%s\n", searchAirport(data.AdList, strings.Join(flag.Args(), " ")))
		fmt.Printf("\tFound in: %.3f ms.\n\n", time.Since(startTotal).Seconds()*1000)
		return
	}

	// Construct two maps key-values [iata]=>(icao) and [icao]=>(airport)+(details)
	mIata2Icao := make(map[string]string)
	mIcao2AirportInfos := make(map[string]string)
	for _, line := range data.AdList {
		lineSplit := strings.Split(string(line), ";")
		mIata2Icao[lineSplit[0]] = lineSplit[1]
		mIcao2AirportInfos[lineSplit[1]] = fmt.Sprintf("(%s) %s, %s", lineSplit[0], lineSplit[2], lineSplit[3])
	}

	// Parse airport list and convert IATA code (3 char.) to ICAO code (4 char.)
	var sStations []string
	for _, v := range flag.Args() {
		v = strings.ToUpper(v)
		fmtNotFound := "\n\t\"%s\" not found. Try to run: metar -s %[1]s\n"

		switch len(v) {
		// if IATA airport code (3 char.), lookup the mIata2Icao map
		case 3:
			if _, ok := mIata2Icao[v]; ok {
				sStations = append(sStations, mIata2Icao[v])
			} else {
				fmt.Printf(fmtNotFound, v)
			}
		// if ICAO airport code (4 char.), lookup the mIcao2AirportInfos map
		case 4:
			if _, ok := mIcao2AirportInfos[v]; ok {
				sStations = append(sStations, v)
			} else {
				fmt.Printf(fmtNotFound, v)
			}
		default:
			fmt.Printf(fmtNotFound, v)
		}
	}

	// if no station to process --> exit
	if len(sStations) == 0 {
		fmt.Print("\n\tNothing to fetch. Quitting...\n\n")
		return
	}

	// prepare the station list string for the URL
	stationList := strings.Join(sStations, "%20")

	// Start download timer
	startDownload := time.Now()

	var wg sync.WaitGroup
	var metarResponse, tafResponse string
	var errM, errT error

	// get METARs routine (arg[4]--> 2 METARs per hour + 30 minutes)
	urlM := fmt.Sprintf(urlMETARfmt, "metars", stationList, float32(*numberMetarFlagInt)/2+0.5)
	wg.Add(1)
	go func() {
		defer wg.Done()
		metarResponse, errM = wget(urlM, *timeoutFlagInt)
	}()

	// get TAFS routine
	urlT := fmt.Sprintf(urlTAFfmt, "tafs", stationList, 0.3)
	wg.Add(1)
	go func() {
		defer wg.Done()
		tafResponse, errT = wget(urlT, *timeoutFlagInt)
	}()

	// Wait for all go routines to finish
	wg.Wait()

	if errM != nil {
		fmt.Printf("\n%s\n\n", checkWgetErr(errM))
		return
	}
	if errT != nil && !*metarOnlyFlagBool {
		fmt.Printf("\n%s\n\n", checkWgetErr(errT))
		return
	}

	// parse responses
	mMetars, err := parseMetarNOAA(metarResponse)
	if err != nil {
		fmt.Println("\n", err)
	}

	var mTafs map[string][]string
	if !*metarOnlyFlagBool {
		mTafs, err = parseTafNOAA(tafResponse)
		if err != nil {
			fmt.Println("\n", err)
		}
	}

	// stop download timer
	downloadTime := time.Since(startDownload).Seconds()

	// final print to terminal
	for _, v := range sStations {

		// Airport title or separator (raw option)
		if !*rawFlagBool {
			fmt.Printf("\n%s %s\n", v, mIcao2AirportInfos[v])
		}

		// print METARs
		if len(mMetars[v]) != 0 {
			for _, vv := range mMetars[v] {
				fmt.Printf("%s\n", vv)
			}
		} else {
			fmt.Println("No METAR received for this station")
		}

		if !*metarOnlyFlagBool {
			// print TAFs
			if len(mTafs[v]) != 0 {
				fmt.Printf("%s\n", mTafs[v][0])
			} else {
				fmt.Println("No TAF received for this station")
			}
		}
	}

	// print timing (not for raw output)
	if !*rawFlagBool {
		totalTime := time.Since(startTotal).Seconds()
		fmt.Printf("\nv2.3 | Download: %.3f sec. | Process: %.3f sec. | Total: %.3f sec.\n",
			downloadTime,
			totalTime-downloadTime,
			totalTime,
		)
	}
}

// Functions

// wget HTTP fetches URL content
func wget(urlString string, wgetTimeout int) (string, error) {

	timeout := time.Duration(wgetTimeout) * time.Second
	client := http.Client{Timeout: timeout}

	// Get page and check for error (timeout, http ...)
	res, err := client.Get(urlString)

	// unwrap url.error and check error and its type
	if e, ok := err.(*url.Error); ok && e.Timeout() {
		return "", fmt.Errorf("Timeout")
	} else if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// if not HTTP 200 OK in response header
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Page not found")
	}

	wgetAnswer, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// return output (after removing trailing \n)
	return string(wgetAnswer[:len(wgetAnswer)-1]), nil
}

// checkWgetErr check connection error sent by wget
func checkWgetErr(err error) string {

	// Change error string if timeout
	if err.Error() == "Timeout" {
		return "Connection Timeout\nTry again later or increase timeout with the -t option (in sec.)"
	}
	return err.Error()
}

// computeFactors Extracts wind, temp and dew point
// to calculate wind chill, heat factors and relative humidity
func computeFactors(wind float64, temp float64, dew float64) (float64, float64, float64) {

	// wind in kmh
	wind *= 1.852

	// Wind Chill (if within limits)
	var wc float64
	if (wind > 5) && (temp < 10) {
		wc = 13.2 + 0.6215*temp + (0.3965*temp-11.37)*math.Pow(wind, 0.16)
	} else {
		wc = temp
	}

	// Relative Humidity (rh)
	tf := temp*1.8 + 32
	rh := 100 * math.Pow((112-0.1*temp+dew)/(112+0.9*temp), 8)

	// Heat Factor (if within limits)
	var hf float64
	if temp >= 27 && dew >= 12 && rh >= 40 {
		const (
			c1 = -42.379
			c2 = 2.04901523
			c3 = 10.14333127
			c4 = -0.22475541
			c5 = -6.83783 / 1e3
			c6 = -5.481717 / 1e2
			c7 = 1.22874 / 1e3
			c8 = 8.5282 / 1e4
			c9 = -1.99 / 1e6
		)

		hff := c1 + c2*tf + c3*rh + c4*tf*rh + c5*math.Pow(tf, 2) + c6*math.Pow(rh, 2) +
			c7*math.Pow(tf, 2)*rh + c8*tf*math.Pow(rh, 2) + c9*math.Pow(tf, 2)*math.Pow(rh, 2)
		hf = (hff - 32) / 1.8
	} else {
		hf = temp
	}

	return wc, hf, rh

}

// searchAirport searches airport in airport data list
func searchAirport(airports []string, searchText string) string {
	list := ""
	searchText = strings.ToUpper(searchText)
	for _, line := range airports {
		if strings.Contains(strings.ToUpper(line), searchText) {
			list += fmt.Sprintf("\t%s\n", strings.Replace(line, ";", " ", -1))
		}
	}
	if list == "" {
		return "\tNothing found\n"
	}
	return list
}

// parseMetarNOAA parse METARs received from NOAA server
func parseMetarNOAA(metarResponse string) (map[string][]string, error) {
	aM := strings.Split(metarResponse, "\n")
	mMetars := make(map[string][]string)
	var factors string

	// add METARS.
	if aM[0] != "No errors" {
		return mMetars, fmt.Errorf("Error METAR: NOAA weather server error: %s", aM[0])
	}

	// Skip the first 6 lines
	for _, aVal := range aM[6:] {

		// Split fields
		fields := strings.Split(aVal, ",")

		// Store raw METAR and ICAO airport ID
		raw := fields[0]
		id := fields[1]

		// loop if maximum number of METAR si reached for one station
		if len(mMetars[id]) >= *numberMetarFlagInt {
			continue
		}

		// If raw not requested, compute wind chill factor, heat factor and relative humidity
		if !*rawFlagBool {
			temp, _ := strconv.ParseFloat(fields[5], 64)
			dew, _ := strconv.ParseFloat(fields[6], 64)
			wind, _ := strconv.ParseFloat(fields[8], 64)
			wc, hf, rh := computeFactors(wind, temp, dew)
			factors = fmt.Sprintf(" [%.0f %.0f %.0f%%]", wc, hf, rh)
			raw += factors
		} else {
			raw = "M:" + raw
		}
		mMetars[id] = append(mMetars[id], raw)
	}

	return mMetars, nil
}

// parseTafNOAA parse TAFs received from NOAA server
func parseTafNOAA(tafResponse string) (map[string][]string, error) {
	aT := strings.Split(tafResponse, "\n")
	mTafs := make(map[string][]string)

	if aT[0] != "No errors" {
		return mTafs, fmt.Errorf("Error TAF: NOAA weather server error: %s", aT[0])
	}
	// Skip the first 6 lines
	for _, aVal := range aT[6:] {

		// Remove trailing ,,, (only interested in the first field)
		fields := aVal[:strings.Index(aVal, ",")]

		// Remove leading "TAF COR" or "TAF" if present
		if fields[:8] == "TAF COR " {
			fields = fields[8:]
		} else if fields[:4] == "TAF " {
			fields = fields[4:]
		}

		// put ICAO ident into id
		id := fields[:4]

		// If there is already a TAF for that station --> loop (only 1 TAF per station)
		if len(mTafs[id]) == 1 {
			continue
		}

		// Store TAF in mTafs map
		raw := fields[5:]
		header := "TAF "
		if *rawFlagBool {
			header = "T:" + id + " "
		}
		mTafs[id] = append(mTafs[id], header+raw)
	}

	return mTafs, nil
}
