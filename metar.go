// metar is a go program to fetch the aviation METAR's and TAF's for a given list of airports in console (terminal) mode.
//
// Usage:
// Retrieve messages for a list of stations (IATA or ICAO codes):
//		$ metar lhr jfk bru uudd
// Find the IATA/ICAO airport code for an airport
//		$ metar -s munich
//		$ metar -s new york
// Help screen:
//		$ metar -h
// Bug reports:
// https://github.com/esperlu/metar/issues

package main

import (
	"./data"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Typical URL:
// https://www.aviationweather.gov/dataserver/example?datatype=metar

// Constants to fetch Weather reports from aviationweather.com
const (
	URL         = "http://aviationweather.gov/adds/dataserver_current/httpparam?dataSource=%s&requestType=retrieve&format=csv&stationString=%s"
	urlMETARfmt = URL + "&hoursBeforeNow=%d"
	urlTAFfmt   = URL + "&hoursBeforeNow=%.1f&mostRecentForEachStation=true&Fields=raw_text"
	maxNbMETAR  = 70
	maxTIMEOUT  = 10
)

// Metar struct for METAR variables
type Metar struct {
	Raw  string
	ID   string
	Temp float64
	Dew  float64
	Wind float64
	Gust float64
}

// Taf struct for TAF variables
type Taf struct {
	Raw string
	ID  string
}

func main() {
	startTotal := time.Now()

	// Get init variables from data package
	adList, Help := data.InitVariables()

	// parse the command line options
	searchFlagBool := flag.Bool("s", false, "Search IATA/ICAO code for an airport")
	rawFlagBool := flag.Bool("r", false, "Print raw data w/o the additional factors")
	numberMetarFlagInt := flag.Int("n", 4, "Set number of Metars to print per station. N 1 to 30.")
	timeoutFlagInt := flag.Int("t", 2, "Change the default timeout of 2 sec. to a maximum of 10 sec.")

	var Usage = func() {
		fmt.Println(Help)
	}
	flag.Usage = Usage
	flag.Parse()

	// If option search but no search pattern
	if *searchFlagBool && len(flag.Args()) == 0 {
		fmt.Printf("\n\tAsked for search, but no search pattern given. Quitting...\n\tTry: metar -s munich or metar -s mun\n\n")
		return

		// If option raw but no airport
	} else if *rawFlagBool && len(flag.Args()) == 0 {
		fmt.Printf("\n\tRaw output option requested, but no airport(s) given. Quitting...\n\tTry: metar -r MUC ebbr JFK\n\n")
		return

		// No args -> help screen
	} else if len(flag.Args()) == 0 {
		Usage()
		return
	}

	// validate number of reports (INT)
	if *numberMetarFlagInt <= 0 || *numberMetarFlagInt > maxNbMETAR {
		fmt.Printf(
			"\n\t%s\n\t%s\n\n",
			"Invalid value for option -n (number of METARS).",
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

	// search option
	if *searchFlagBool {
		fmt.Printf("\n%s\n", searchAirport(adList, strings.Join(flag.Args(), " ")))
		fmt.Printf("\tFound in: %.3f ms.\n\n", time.Since(startTotal).Seconds()*1000)
		return
	}

	// Construct two maps [iata]=>(icao) and [icao]=>(airport)+(details)
	mAirportIcao4 := make(map[string]string)
	mAirportIata3 := make(map[string]string)
	for _, line := range adList {
		lineSplit := strings.Split(string(line), ";")
		mAirportIata3[lineSplit[0]] = lineSplit[1]
		mAirportIcao4[lineSplit[1]] = fmt.Sprintf("(%s) %s, %s", lineSplit[0], lineSplit[2], lineSplit[3])
	}

	// Parse airport list
	var sStations []string
	for _, v := range flag.Args() {
		V := strings.ToUpper(v)
		fmtNotFound := "\n\t\"%s\" not found. Try to run: metar -s %[1]s\n"

		switch len(V) {
		// if IATA airport code (3 char.), lookup the mAirportIata3 map
		case 3:
			if _, ok := mAirportIata3[V]; ok {
				sStations = append(sStations, mAirportIata3[V])
			} else {
				fmt.Printf(fmtNotFound, v)
			}
		// if ICAO airport code (4 char.), lookup the mAirportIcao4 map
		case 4:
			if _, ok := mAirportIcao4[V]; ok {
				sStations = append(sStations, V)
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
		os.Exit(1)
	}

	// prepare the station list string for the URL
	stationList := strings.Join(sStations, "%20")

	startDownload := time.Now()
	var wg sync.WaitGroup
	var metars, tafs string

	// get METARS
	url := fmt.Sprintf(urlMETARfmt, "metars", stationList, *numberMetarFlagInt)
	wg.Add(1)
	go func(urlM string) {
		metars = wget(urlM, *timeoutFlagInt)
		wg.Done()
	}(url)

	// get TAFS
	url = fmt.Sprintf(urlTAFfmt, "tafs", stationList, 0.3)
	wg.Add(1)
	go func(urlT string) {
		tafs = wget(urlT, *timeoutFlagInt)
		wg.Done()
	}(url)

	// Wait for all go routines to finish
	wg.Wait()

	downloadTime := time.Since(startDownload).Seconds()

	// store each line in a slice of strings
	aM := strings.Split(metars, "\n")
	aT := strings.Split(tafs, "\n")

	// initialize map for metars and tafs
	mMetars := make(map[string][]string)
	mTafs := make(map[string][]string)

	var factors string
	var m Metar

	// add METARS.
	if aM[0] == "No errors" {

		// Skip the first 6 lines
		for _, aVal := range aM[6:] {

			// Split fields
			fields := strings.Split(aVal, ",")

			// Store ICAO airport ID
			id := fields[0][:4]

			// Stop the for loop if maximum number of METAR si reached
			if len(mMetars[id]) >= *numberMetarFlagInt {
				continue
			}

			raw := fields[0]
			// If only raw requested, don't compute wind chill factor, heat factor and relative humidity
			if *rawFlagBool {
				m = Metar{Raw: raw, ID: id}
			} else {
				temp, _ := strconv.ParseFloat(fields[5], 64)
				dew, _ := strconv.ParseFloat(fields[6], 64)
				wind, _ := strconv.ParseFloat(fields[8], 64)
				m = Metar{Raw: raw, ID: id, Temp: temp, Dew: dew, Wind: wind}
				if !*rawFlagBool {
					wc, hf, rh := computeFactors(wind, temp, dew)
					factors = fmt.Sprintf(" [%.0f %.0f %.0f%%]", wc, hf, rh)
				}
			}
			mMetars[m.ID] = append(mMetars[m.ID], m.Raw+factors)
		}

	} else {

		fmt.Printf("\n\tMETAR: weather server error:\n\t%s\n", aM[0])

	}

	// add TAFS.
	if aT[0] == "No errors" {

		// Skip the first 6 lines
		for _, aVal := range aT[6:] {

			// Remove trailing ,,,
			fields := aVal[:strings.Index(aVal, ",")]

			// Remove leading "TAF COR" or "TAF" if present
			if fields[:8] == "TAF COR " {
				fields = fields[8:]
			} else if fields[:4] == "TAF " {
				fields = fields[4:]
			}

			// put ICAO ident into id
			id := fields[:4]

			// only 1 TAF per station
			if len(mTafs[id]) >= 1 {
				break
			}

			// Store TAF in mTafs map
			raw := fields[5:]
			t := Taf{ID: id, Raw: raw}
			mTafs[t.ID] = append(mTafs[t.ID], "TAF "+t.Raw)
		}

	} else {

		fmt.Printf("\n\tTAF: weather server error:\n\t%s\n", aT[0])

	}

	// final print

	for _, v := range sStations {
		// Airport title or separator (raw option)
		if *rawFlagBool {
			fmt.Println("")
		} else {
			fmt.Printf("\n%s %s\n", v, mAirportIcao4[v])
		}

		// print METARS
		if len(mMetars[v]) != 0 {
			for _, vv := range mMetars[v] {
				fmt.Printf("%s\n", vv)
			}
		} else {
			fmt.Println("No METAR received for this station")
		}

		// print most recent TAFS
		if len(mTafs[v]) != 0 {
			for k, vv := range mTafs[v] {
				if k == 2 {
					break
				}
				fmt.Printf("%s\n", vv)
			}
		} else {
			fmt.Println("No TAF received for this station")
		}
	}
	// print timing (not for raw output)
	if !*rawFlagBool {
		totalTime := time.Since(startTotal).Seconds()
		fmt.Printf("\nv2.1 | Download: %.3f sec. | Process: %.3f | Total: %.3f sec.\n",
			downloadTime,
			totalTime-downloadTime,
			totalTime,
		)
	}
}

// Functions

// wget HTTP fetches URL content
func wget(url string, wgetTimeout int) string {
	timeout := time.Duration(wgetTimeout) * time.Second
	client := http.Client{
		Timeout: timeout,
	}

	// Get page and check for error (timeout, http ...)
	res, err := client.Get(url)
	if err != nil {
		if strings.Index(err.Error(), "Client.Timeout exceeded") > 0 {
			errText := fmt.Sprintf("\n\n\tConnection timeout (%.0f sec.)\n\tCheck your internet connection,"+
				" try again later\n\tor increase timeout with the -t option (in sec.)\n\n", timeout.Seconds())
			log.Fatal(fmt.Errorf(errText))
		}
	}
	defer res.Body.Close()

	// if not HTTP 200 OK in response header
	if res.StatusCode != http.StatusOK {
		log.Fatal(
			fmt.Errorf("\n\n\t%s", "Page not found\n\n"),
		)
	}

	wgetAnswer, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// send outpput to ch (after removing trailing \n)
	return fmt.Sprintf("%s", wgetAnswer[:len(wgetAnswer)-1])
}

// computeFactors WindChill HeatFactor Relative Humidity - Extract wind, temp and dew point
// to calculate wind chill, heat factors and relative humidity
func computeFactors(wind float64, temp float64, dew float64) (float64, float64, float64) {

	wind = wind * 1.852 // wind in kmh

	// Wind Chill (if within limits)
	var wc float64
	if (wind < 5) || (temp > 10) {
		wc = temp
	} else {
		wc = 13.2 + 0.6215*temp + (0.3965*temp-11.37)*math.Pow(wind, 0.16)
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
func searchAirport(adFile []string, searchText string) string {
	list := ""
	searchText = strings.ToUpper(searchText)
	for _, line := range adFile {
		if strings.Contains(strings.ToUpper(line), searchText) {
			list += fmt.Sprintf("\t%s\n", strings.Replace(line, ";", " ", -1))
		}
	}
	if list == "" {
		return "\tNothing found\n"
	}
	return list
}
