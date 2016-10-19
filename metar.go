// JL Lacroix 2016

// metar is a go program to fetch the aviation METAR's and TAF's for a given list of airports in console (terminal) mode.
//
// Usage:
// Retrieve messages for a list of stations (IATA or ICAO codes):
//		$ metar lhr jfk bru uudd
// Find the IATA/ICAO airport code for an airport
//		$ metar -s munich
//		$ metar -s "new york"
// Help screen:
//		$ metar -h
// Bug reports:
// https://github.com/esperlu/metar/issues
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"metar/data"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	baseUrl  = "http://aviationweather.gov/adds/dataserver_current/httpparam?dataSource=%s&requestType=retrieve&format=csv&stationString=%s"
	urlMetar = baseUrl + "&hoursBeforeNow=%d"
	urlTaf   = baseUrl + "&hoursBeforeNow=%.1f&mostRecentForEachStation=postfilter"
)

type Metar struct {
	Raw  string
	Id   string
	Temp float64
	Dew  float64
	Wind float64
	Gust float64
}

type Taf struct {
	Raw string
	Id  string
}

func main() {
	startTotal := time.Now()

	// Get init variables from pkg myfunctions
	adList, Help := data.InitVariables()

	// parse the command line options
	searchFlagBool := flag.Bool("s", false, "Search IATA/ICAO code for an airport")
	rawFlagBool := flag.Bool("r", false, "Print raw data w/o the additional factors")
	numberFlagStr := flag.String("n", "4", "Set number of Metars to print per station. N 1 to 10.")

	var Usage = func() {
		fmt.Println(Help)
	}
	flag.Usage = Usage
	flag.Parse()

	if len(flag.Args()) == 0 {
		Usage()
		os.Exit(1)
	}

	// validate flag numberFlagStr. Format int[,.]int
	re := regexp.MustCompile("^(10|[1-9])$")
	aMatches := re.FindStringSubmatch(*numberFlagStr)
	if len(aMatches) == 0 {
		fmt.Printf(
			"\n\t%s\n\t%s\n\n",
			"Invalid value for option -n.",
			"Minimum 1, maximum 10",
		)
		os.Exit(1)
	}

	// assign values to maxNbrMetars & maxNbrTafs
	maxNbrMetars, _ := strconv.Atoi(aMatches[1])

	// Search ICAO/IATA airport code
	if *searchFlagBool && len(flag.Args()) == 0 {
		fmt.Printf("\n\tNo search pattern given. Quitting...\n\tTry: metar -s munich\n\n")
		os.Exit(1)
	}

	if *searchFlagBool {
		SearchAirport(adList, flag.Args()[0])
		totalTime := time.Since(startTotal)
		fmt.Printf("\nProcessed in: %.3f sec.\n",
			totalTime.Seconds(),
		)
		os.Exit(0)
	}

	// Construct two maps [iata]=>(icao) and [icao]=>(airport)+(details)
	mAirportIcao4 := make(map[string]string)
	mAirportIata3 := make(map[string]string)
	for _, line := range adList {
		lineSplit := strings.Split(string(line), ";")
		mAirportIata3[lineSplit[0]] = lineSplit[1]
		mAirportIcao4[lineSplit[1]] = fmt.Sprintf("(%s) %s, %s", lineSplit[0], lineSplit[2], lineSplit[3])
	}

	// Otherwise parse airport list
	var sStations []string
	for _, v := range flag.Args() {
		V := strings.ToUpper(v)
		fmtNotFound := "\n\t\"%s\" not found. Try to run: metar -s %[1]s\n"
		if strings.Contains(v, " ") {
			fmtNotFound = "\n\t\"%s\" not found. Try to run: metar -s \"%[1]s\"\n"
		}

		switch len(V) {
		// if IATA airport code (3 characters), lookup the mAirportIata3 map
		case 3:
			if _, ok := mAirportIata3[V]; ok {
				sStations = append(sStations, mAirportIata3[V])
			} else {
				fmt.Printf(fmtNotFound, v)
			}
		// if ICAO airport code (4 characters), lookup the mAirportIcao4 map
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
		fmt.Println("\n\tNothing to fetch. Quitting...\n")
		os.Exit(1)
	}

	// prepare the station list string for the URL
	stationList := strings.Join(sStations, "%20")

	// start goroutines and store result in chanels
	// METARS
	startDownload := time.Now()
	chanMetars := make(chan string)
	url := fmt.Sprintf(urlMetar, "metars", stationList, maxNbrMetars)
	go Wget(url, 2, chanMetars)

	// TAFS
	chanTafs := make(chan string)
	url = fmt.Sprintf(urlTaf, "tafs", stationList, 0.3)
	go Wget(url, 2, chanTafs)

	// Read chanels
	metars := <-chanMetars
	tafs := <-chanTafs

	downloadTime := time.Since(startDownload)

	// store every line in a slice
	aM := strings.Split(metars, "\n")
	aT := strings.Split(tafs, "\n")

	// check for errors in weather server response
	if aM[0] != "No errors" || aT[0] != "No errors" {
		fmt.Println("Error in response from weather server")
		os.Exit(1)
	}

	// initialize map for metars and tafs
	mMetars := make(map[string][]string)
	mTafs := make(map[string][]string)

	// add METARS. Skip the first 6 lines
	var factors string
	var m Metar

	for _, aVal := range aM[6:] {
		fields := strings.Split(aVal, ",")
		id := fields[1]
		if len(mMetars[id]) >= maxNbrMetars { // max number of METARS per station
			continue
		}
		raw := fields[0]
		if *rawFlagBool {
			m = Metar{Raw: raw, Id: id}
		} else {
			temp, _ := strconv.ParseFloat(fields[5], 64)
			dew, _ := strconv.ParseFloat(fields[6], 64)
			wind, _ := strconv.ParseFloat(fields[8], 64)
			m = Metar{Raw: raw, Id: id, Temp: temp, Dew: dew, Wind: wind}
			if !*rawFlagBool {
				wc, hf, rh := Factors(wind, temp, dew)
				factors = fmt.Sprintf(" [%.0f %.0f %.0f%%]", wc, hf, rh)
			}
		}
		mMetars[m.Id] = append(mMetars[m.Id], m.Raw+factors)
	}

	// add TAFS. Skip the first 6 lines and remove trailing ,,,
	for _, aVal := range aT[6:] {
		fields := strings.Split(aVal, ",")
		id := fields[1]
		if len(mTafs[id]) >= 1 { // max number of TAFS per station
			continue
		}
		raw := fields[0]
		tafHeader := ""
		if raw[:3] != "TAF" {
			tafHeader = "TAF "
		}
		t := Taf{Id: id, Raw: raw}
		mTafs[t.Id] = append(mTafs[t.Id], tafHeader+t.Raw)
	}

	// final print
	for _, v := range sStations {
		// Airport title or separator (raw option)
		if *rawFlagBool {
			fmt.Println("=")
		} else {
			fmt.Printf("\n%s %s\n", v, mAirportIcao4[v])
		}
		// If no METARS and no TAFS
		if len(mMetars[v]) == 0 && len(mTafs[v]) == 0 {
			fmt.Println("No reports for this station")

			// else print all METARS and ...
		} else {
			for _, vv := range mMetars[v] {
				fmt.Printf("%s\n", vv)
			}

			// ... print the most recent TAF
			for k, vv := range mTafs[v] {
				if k == 2 {
					break
				}
				fmt.Printf("%s\n", vv)
			}
		}
	}
	// print timing
	if !*rawFlagBool {
		totalTimeS := time.Since(startTotal).Seconds()
		downloadTimeS := downloadTime.Seconds()
		fmt.Printf("\nDownload: %.3f sec. | Process: %.3f sec. | Total: %.3f sec.\n",
			downloadTimeS,
			totalTimeS-downloadTimeS,
			totalTimeS,
		)
	}
}

// Functions

// Wget HTTP fetches URL content
func Wget(url string, wgetTimeout time.Duration, ch chan<- string) {
	timeout := time.Duration(wgetTimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	// Get page and check for error (timeout, http ...)
	res, err := client.Get(url)
	if err != nil {
		log.Fatal(
			fmt.Errorf(
				"\n\n\t%s\n\t%s\n\n",
				"No response could be received from the weather server.",
				"Check your internet connection or try again later.",
			),
		)
	}
	defer res.Body.Close()

	// if not HTTP 200 OK in response header
	if res.StatusCode != http.StatusOK {
		log.Fatal(
			fmt.Errorf("\n\n\t%s\n\n", "Page not found"),
		)
	}

	wgetAnswer, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// send outpput to ch (after removing trailing \n)
	ch <- fmt.Sprintf("%s", wgetAnswer[:len(wgetAnswer)-1])
}

// WindChillHeatFactorRelativeHumidity Extract wind, temp and dew point
// to calculate wind chill, heat factors and relative humidity
func Factors(wind float64, temp float64, dew float64) (float64, float64, float64) {

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

// SearchAirport searches airport in airport data list
func SearchAirport(adFile []string, searchText string) {
	list := ""
	searchText = strings.ToUpper(searchText)
	for _, line := range adFile {
		if strings.Contains(strings.ToUpper(line), searchText) {
			list += fmt.Sprintf("%s\n", strings.Replace(line, ";", " ", -1))
		}
	}
	if list != "" {
		fmt.Println("\n" + list)
	} else {
		fmt.Println("\nNothing found\n")
	}
}
