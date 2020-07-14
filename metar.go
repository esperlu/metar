// metar is a go program to fetch the aviation METAR's and TAF's for a given list of airports in console (terminal) mode.

// Usage:
// Retrieve messages for a list of stations (IATA or ICAO codes):
// 		$ metar lhr jfk bru uudd
// Find the IATA/ICAO airport code for an airport
// 		$ metar -s munich
// 		$ metar -s new york
// Help screen for other options:
// 		$ metar -h
// Bug reports:
// https://github.com/esperlu/metar/issues
package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	// "jeanluc/metarDEV/data"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"./data"
)

// Constants to fetch Weather reports from aviationweather.com
const (
	URLfmt      = "http://aviationweather.gov/adds/dataserver_current/httpparam?dataSource=%s&requestType=retrieve&format=csv&stationString=%s"
	urlMETARfmt = URLfmt + "&hoursBeforeNow=%.1f"
	urlTAFfmt   = URLfmt + "&hoursBeforeNow=%.1f&mostRecentForEachStation=true&Fields=raw_text"
	maxNbMETAR  = 70
	maxTIMEOUT  = 10
	ver         = "2.4.1"
)

// Initialize and parse flags
var (
	numberMetarFlag   = flag.Int("n", 4, "Set number of Metars to print per station. N 1 to 30.")
	searchFlag        = flag.Bool("s", false, "Search IATA/ICAO code for an airport")
	listCountriesFlag = flag.Bool("lc", false, "List all countries with ISO code")
	listAirportsFlag  = flag.Bool("la", false, "List all for one country (ISO code)")
	timeoutFlag       = flag.Int("t", 2, "Change the default timeout of 2 sec. to a maximum of 10 sec.")
	rawFlag           = flag.Bool("r", false, "Print raw data w/o the additional factors")
	metarOnlyFlag     = flag.Bool("m", false, "METARs only")
	tafOnlyFlag       = flag.Bool("f", false, "TAF (forecast) only")
)

func init() {
	flag.Usage = func() {
		fmt.Println(data.Help)
	}
	flag.Parse()
}

func main() {

	startTotal := time.Now()

	// If option search but no search pattern
	if *searchFlag && len(flag.Args()) == 0 {
		fmt.Printf("\n Asked for search, but no search pattern given. Quitting...\n  Try: `metar -s munich` or `metar -s mun`\n\n")
		return
	}

	// No args -> help screen
	if !*listCountriesFlag && !*listAirportsFlag && len(flag.Args()) == 0 {
		fmt.Printf("\n  No airport(s) given. Quitting...\n  Try: `metar MUC ebbr JFK` (more examples below)\n\n")
		fmt.Println(data.Help)
		return
	}

	// validate number of reports (INT)
	if *numberMetarFlag <= 0 || *numberMetarFlag > maxNbMETAR {
		fmt.Printf(
			"\n  %s\n  %s\n\n",
			"Invalid value for option -n (number of METARs).",
			"Minimum 1, maximum 70",
		)
		return
	}

	// validate timeout flag (INT)
	if *timeoutFlag < 1 || *timeoutFlag > maxTIMEOUT {
		fmt.Printf(
			"\n  %s\n  %s\n\n",
			"Invalid value for option -t (connection timeout).",
			"Minimum 1, maximum 10",
		)
		return
	}

	// options -m and -f mutually exclusive
	if *metarOnlyFlag && *tafOnlyFlag {
		fmt.Printf("\n options -m (metar only) and -f (taf only) are mutually exclusive. Nothing to print.\n\n")
		return
	}

	// Option -la (list airports) needs arguments.
	if *listAirportsFlag && len(flag.Args()) == 0 {
		fmt.Printf("\n %s\n %s\n\n",
			"option -la (list country airports) needs at least one argument (country code)",
			"Try: `metar -la BE` or for multiple countries: `metar FR CH PT`",
		)
		return
	}

	// Build key-values tables
	// Table map[countryCode] -> []string{Country name, continent code}
	code2country := make(map[string][]string)
	for _, line := range data.CountryList {
		lineSplit := strings.Split(line, ";")
		code2country[lineSplit[1]] = []string{lineSplit[0], lineSplit[2]}

	}
	// for k, v := range code2country {
	// 	fmt.Printf("[%s] %s\n", k, v)
	// }
	// return

	// Table map[iata]=>(icao) and map[icao]=>(airport)+(details)
	iata2icao := make(map[string]string)
	icao2airportInfos := make(map[string]string)
	for _, line := range data.AdList {
		// fmt.Println(line)
		lineSplit := strings.Split(string(line), ";")
		iata2icao[lineSplit[1]] = lineSplit[0]
		// empty iata code
		if lineSplit[1] == "" {
			lineSplit[1] = "---"
		}
		icao2airportInfos[lineSplit[0]] = fmt.Sprintf(
			"%s (%s) %s, %s %s (%s)",
			lineSplit[0],
			lineSplit[1],
			lineSplit[2],
			code2country[lineSplit[3]][0],
			lineSplit[3],
			code2country[lineSplit[3]][1],
		)
	}

	// for k, v := range icao2airportInfos {
	// 	fmt.Printf("[%s] %s\n", k, v)
	// }
	// return

	// search option -s
	if *searchFlag {
		fmt.Printf("\n%s\n", searchAirport(icao2airportInfos, strings.Join(flag.Args(), " ")))
		fmt.Printf("  Processed in %.3f ms.\n\n", time.Since(startTotal).Seconds()*1000)
		return
	}

	// List country ISO codes
	if *listCountriesFlag {
		fmt.Printf("\n%s\n", listCountries(strings.Join(flag.Args(), " ")))
		fmt.Printf("  Processed in %.3f ms.\n\n", time.Since(startTotal).Seconds()*1000)
		return
	}

	// List country airports
	if *listAirportsFlag {
		if len(flag.Args()[0]) != 2 {
			fmt.Printf(
				"\n  %s\n  %s\n  %s\n\n",
				"Invalid value for option -la (list country's airports).",
				"Expecting: 2 characters of a ISO country code.",
				"Use option -lc <STRING> to list ISO country codes.",
			)
			return
		}
		fmt.Printf("\n%s\n", listAirports(flag.Args(), code2country))
		fmt.Printf("  Processed in %.3f ms.\n\n", time.Since(startTotal).Seconds()*1000)
		return
	}

	// Parse airport list and convert IATA code (3 char.) to ICAO code (4 char.)
	var stations []string
	for _, v := range flag.Args() {
		v = strings.ToUpper(v)
		fmtNotFound := "\n  \"%s\" not found. Try to run: metar -s %[1]s\n"

		switch len(v) {
		// if IATA airport code (3 char.), lookup the iata2icao map
		case 3:
			if _, ok := iata2icao[v]; ok {
				stations = append(stations, iata2icao[v])
			} else {
				fmt.Printf(fmtNotFound, v)
			}
		// if ICAO airport code (4 char.), lookup the icao2airportInfos map
		case 4:
			if _, ok := icao2airportInfos[v]; ok {
				stations = append(stations, v)
			} else {
				fmt.Printf(fmtNotFound, v)
			}
		default:
			fmt.Printf(fmtNotFound, v)
		}
	}

	// if no station to process --> exit
	if len(stations) == 0 {
		fmt.Print("\n  Nothing to fetch. Quitting...\n\n")
		return
	}

	// prepare the station list string for the URL
	stationList := strings.Join(stations, "%20")

	// Start download timer
	startDownload := time.Now()

	var wg sync.WaitGroup
	var metarResponse, tafResponse string
	var errM, errT error

	// get METARs routine (arg[4]--> 2 METARs per hour + 30 minutes)
	urlM := fmt.Sprintf(urlMETARfmt, "metars", stationList, float32(*numberMetarFlag)/2+0.5)
	wg.Add(1)
	go func() {
		defer wg.Done()
		metarResponse, errM = wget(urlM, *timeoutFlag)
	}()

	// get TAFS routine
	urlT := fmt.Sprintf(urlTAFfmt, "tafs", stationList, 0.3)
	wg.Add(1)
	go func() {
		defer wg.Done()
		tafResponse, errT = wget(urlT, *timeoutFlag)
	}()

	// Wait for all go routines to finish
	wg.Wait()

	if errM != nil {
		fmt.Printf("\n%s\n\n", checkWgetErr(errM))
		return
	}
	if errT != nil && !*metarOnlyFlag {
		fmt.Printf("\n%s\n\n", checkWgetErr(errT))
		return
	}

	// parse responses
	metars, err := parseMetarNOAA(metarResponse)
	if err != nil {
		fmt.Printf("\n %s.\n\n", err)
		return
	}

	var tafs map[string][]string
	if !*metarOnlyFlag {
		tafs, err = parseTafNOAA(tafResponse)
		if err != nil {
			fmt.Printf("\n %s.\n\n", err)
		}
	}

	// stop download timer
	downloadTime := time.Since(startDownload).Seconds()

	// final print to terminal
	for _, v := range stations {

		// Airport title or separator (raw option)
		if !*rawFlag {
			fmt.Printf("\n%s\n", icao2airportInfos[v])
		}

		// print METARs
		if !*tafOnlyFlag {
			if len(metars[v]) != 0 {
				for _, vv := range metars[v] {
					fmt.Printf("%s\n", vv)
				}
			} else {
				fmt.Println("No METAR received for this station")
			}
		}

		if !*metarOnlyFlag {
			// print TAFs
			if len(tafs[v]) != 0 {
				fmt.Printf("%s\n", tafs[v][0])
			} else {
				fmt.Println("No TAF received for this station")
			}
		}
	}

	// print timing (not for raw output)
	if !*rawFlag {
		totalTime := time.Since(startTotal).Seconds()
		fmt.Printf("\n%s | Download: %.3f sec. | Process: %.3f sec. | Total: %.3f sec.\n",
			ver,
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
	// return string(wgetAnswer), nil
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
	rh := 100 * math.Pow((-0.1*temp+dew+112)/(0.9*temp+112), 8)

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
func searchAirport(icao2airportInfos map[string]string, countryCodes string) string {
	var list string
	countryCodes = strings.ToUpper(countryCodes)
	for _, line := range icao2airportInfos {
		if strings.Contains(strings.ToUpper(line), countryCodes) {
			list += fmt.Sprintf("  %s\n", line)
		}
	}
	if list == "" {
		return fmt.Sprintf("  Nothing found for: '%s'\n", countryCodes)
	}
	return list
}

// listCountries List all countries and ISO code
func listCountries(countryCodes string) string {
	var list string
	countryCodes = strings.ToUpper(countryCodes)
	for _, line := range data.CountryList {
		if strings.Contains(strings.ToUpper(line), countryCodes) {
			splitLine := strings.Split(line, ";")
			list += fmt.Sprintf("  %s %s (%s)\n", splitLine[1], splitLine[0], splitLine[2])
		}
	}
	if list == "" {
		return "  " + countryCodes + ": Nothing found in country list.\n"
	}
	return list
}

// listAirports list all airports for one or more countries (ISO country code)
func listAirports(countryCodes []string, code2country map[string][]string) string {
	var list string
	for _, code := range countryCodes {
		code = strings.ToUpper(code)
		// Check if country code exists
		if _, ok := code2country[code]; !ok {
			return fmt.Sprintf("\n  Unkown country code: %s.\n  Use option -lc [search text] to list county codes.\n", code)
		}
		// Country title line
		header := fmt.Sprintf("\n  %s (%s)\n", code2country[code][0], code2country[code][1])
		countryAdList := ""
		// Loop through all airports for that country code
		for _, line := range data.AdList {
			splitLine := strings.Split(line, ";")
			if code == splitLine[3] {
				countryAdList += fmt.Sprintf("  %s %-3s %s %s\n", splitLine[0], splitLine[1], splitLine[3], splitLine[2])
			}
		}
		if countryAdList != "" {
			list += header + countryAdList
		}
	}
	if list == "" {
		return "  Nothing found.\n"
	}
	return list

}

// parseMetarNOAA parse METARs received from NOAA server
func parseMetarNOAA(metarResponse string) (map[string][]string, error) {
	metars := make(map[string][]string)
	var factors string

	// Split response on EOL and store fields value in map
	aM := strings.Split(metarResponse, "\n")
	if aM[0] != "No errors" {
		return metars, fmt.Errorf("Error METAR: NOAA weather server error: %s", aM[0])
	}

	// Skip the first 6 lines
	for _, aVal := range aM[6:] {

		// Split fields. Retrun error if missing fields.
		fields := strings.Split(aVal, ",")
		if len(fields) != 44 {
			return metars, fmt.Errorf("Error METAR: NOAA format non compliant. Missing fields in response")
		}

		// Store raw METAR and ICAO airport ID
		raw := fields[0]
		id := fields[1]

		// loop if maximum number of METAR si reached for one station
		if len(metars[id]) >= *numberMetarFlag {
			continue
		}

		// If raw not requested, compute wind chill factor, heat factor and relative humidity
		if !*rawFlag && fields[5] != "" && fields[6] != "" && fields[8] != "" {
			temp, _ := strconv.ParseFloat(fields[5], 64)
			dew, _ := strconv.ParseFloat(fields[6], 64)
			wind, _ := strconv.ParseFloat(fields[8], 64)
			wc, hf, rh := computeFactors(wind, temp, dew)
			factors = fmt.Sprintf(" [%.0f %.0f %.0f%%]", wc, hf, rh)
			raw += factors
		} else {
			raw = "M:" + raw
		}
		metars[id] = append(metars[id], raw)
	}
	return metars, nil
}

// parseTafNOAA parse TAFs received from NOAA server
func parseTafNOAA(tafResponse string) (map[string][]string, error) {
	mTafs := make(map[string][]string)

	// Split response on EOL and store fields value in map
	aT := strings.Split(tafResponse, "\n")
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
		if *rawFlag {
			header = "T:" + id + " "
		}
		mTafs[id] = append(mTafs[id], header+raw)
	}

	return mTafs, nil
}
