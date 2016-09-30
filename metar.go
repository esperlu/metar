/*
 * JL 2016-09-20
 * JL 2016-09-30
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"metar/myfunctions"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)


func main() {

	/* Get init variables from pkg myfunctions */
	AdList, Help := myfunctions.InitVariables()

	/* Read  and parse AdList*/
	lines := strings.Split(string(AdList), "\n")
	/* delete last empty element from slice lines due to trailing \n */

	mAirportIcao4 := make(map[string]string)
	mAirportIata3 := make(map[string]string)
	for _, line := range lines {
		lineSplit := strings.Split(line, ";")
		mAirportIcao4[lineSplit[1]] = fmt.Sprintf("(%s) %s, %s", lineSplit[0], lineSplit[2], lineSplit[3])
		mAirportIata3[lineSplit[0]] = lineSplit[1]
	}

	/* read args from command line (w/o program name) */
	args := os.Args[1:]

	/* If option -h or none --> help page */
	if len(args) == 0 || args[0] == "-h" {
		fmt.Println(Help)
		os.Exit(0)
	}

	/* If option -s search string in airport lines and exit */
	if args[0] == "-s" {
		searchAirport(lines, args[1])
		os.Exit(0)
	}

	/* Otherwise parse airport list */
	var sStations []string
	for _, v := range args {
		V := strings.ToUpper(v)
		switch len(V) {
		case 3:
			if _, ok := mAirportIata3[V]; ok {
				sStations = append(sStations, mAirportIata3[V])
			} else {
				fmt.Printf("\n%s not found. Try to run: metar -s %s\n", V, V)
			}
		case 4:
			if _, ok := mAirportIcao4[V]; ok {
				sStations = append(sStations, V)
			} else {
				fmt.Printf("\n%s not found. Try to run: metar -s %s\n", V, V)
			}
		default:
			fmt.Printf("\n%s not found. Try to run: metar -s %s\n", V, V)
		}
	}

	/* if no station to process --> exit */
	if len(sStations) == 0 {
		fmt.Println("\nNothing to fetch. Quitting...\n")
		os.Exit(1)
	}

	/* prepare the stations string for the URL */
	stations := strings.Join(sStations, "%20")
	urlFormat := "http://aviationweather.gov/adds/dataserver_current/httpparam?dataSource=%s&requestType=retrieve&format=csv&stationString=%s&hoursBeforeNow=%.1f&fields=raw_text&mostRecentForEachStation=false"

	/* start goroutines and store result in chanels */
	/* METARS */
	chanMetars := make(chan string)
	url := fmt.Sprintf(urlFormat, "metars", stations, 3.0)
	go wget(url, 2, chanMetars)

	/* TAFS */
	chanTafs := make(chan string)
	url = fmt.Sprintf(urlFormat, "tafs", stations, 0.5)
	go wget(url, 2, chanTafs)

	/* Read chanels */
	metars := <-chanMetars
	tafs := <-chanTafs

	/* store every line in a slice and remove empty element from slice end due to trailing \n */
	aM := strings.Split(metars, "\n")
	aT := strings.Split(tafs, "\n")
	aM = aM[:len(aM)-1]
	aT = aT[:len(aT)-1]

	/* initialize map metar */
	mMetars := make(map[string][]string)
	mTafs := make(map[string][]string)


	/* skip the first 6 lines and remove trailings ,,,,,, */
	for _, aVal := range aM[6:] {
		str := aVal[:strings.Index(aVal, ",")]
		station := strings.Split(str, " ")[0]
		mMetars[station] = append(mMetars[station], str)
	}

	/* add TAF  */
	for _, aVal := range aT[6:] {
		str := aVal[:strings.Index(aVal, ",")]
		station := strings.Split(str, " ")[1]
		mTafs[station] = append(mTafs[station], str)
		//mTafs[station] = str
	}

	/* if no reports found for valid stations ... */
	if len(mMetars) == 0 && len(mTafs) == 0 {
		for _, v := range sStations {
			fmt.Printf("\n%s %s\n", v, mAirportIcao4[v])
			fmt.Println("No reports for this station")
		}
		fmt.Println("")
		os.Exit(0)
	}

	/* final print */
	for _, v := range sStations {
		fmt.Printf("\n%s %s\n", v, mAirportIcao4[v])

		// If no METARS and no TAFS
		if len(mMetars[v]) == 0 && len(mTafs[v]) == 0 {
			fmt.Println("No reports for this station")

		// else print all METARS and ...
		} else {
			for _, vv := range mMetars[v] {
				wc, hf, rh := WindChillHeatFactorRelativeHumidity(vv)
				fmt.Printf("%s [%.0f %.0f %.0f%%] \n", vv, wc, hf, rh)
			}


			// ... print the last 2 TAFS
			for _, vv := range mTafs[v] {
				fmt.Printf("%s\n", vv)
			}

		}
	}

	fmt.Println("")

}

/*
 * Functions
 */

/* HTTP get */
func wget(url string, wgetTimeout time.Duration, ch chan<- string) {
	timeout := time.Duration(wgetTimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	// Get page and check for error (timeout, http ...)
	res, err := client.Get(url)
	if err != nil {
		sErr := err.Error()
		switch {
		case strings.Index(sErr, "no such host") != -1 :
			sErr = "Host not found (Internet connected?)"
		case strings.Index(sErr, "Timeout")  != -1 :
			//log.Fatal(fmt.Errorf("Timeout"))
			sErr = fmt.Sprintf("Timeout (%s): no response received from host. Retry later.", timeout)
		}
		log.Fatal(fmt.Errorf(sErr))
	}
	// if not HTTP 200 OK in response header
	if res.Status != "200 OK" {
		log.Fatal(fmt.Errorf("Page not found"))
	}

	wgetAnswer, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	ch <- fmt.Sprintf("%s", wgetAnswer)
}

/* Extract wind, temp and dew point to calculate wind chill & heat factors and relative humidity */
func WindChillHeatFactorRelativeHumidity(metar string) (float64, float64, float64) {

	re := regexp.MustCompile("^.+([0-9]{2})(KT|MPS).+ (?:M)?([0-9]{2})/(?://|M)?([0-9]{2})")

	w, _ := strconv.ParseFloat((re.FindStringSubmatch(metar)[1]), 64)
	wUnit := re.FindStringSubmatch(metar)[2]
	tc, _ := strconv.ParseFloat(re.FindStringSubmatch(metar)[3], 64)
	dc, _ := strconv.ParseFloat(re.FindStringSubmatch(metar)[4], 64)

	// conversion en km/h
	if wUnit == "MPS" {
		w = w * 3.6
	} else {
		w = w * 1.852
	}

	// calcul du wind chill (if within limits)
	var wc float64
	if (w < 5) || (tc > 10) {
		wc = tc
	} else {
		wc = 13.2 + 0.6215*tc + (0.3965 * tc - 11.37) * math.Pow(w, 0.16)
	}

	// Relative humidity (rh)
	tf := tc*1.8 + 32
	rh := 100 * math.Pow((112 - 0.1 * tc + dc) / (112 + 0.9 * tc), 8)

	// Heat factor (if within limits)
	var hf float64
	if tc >= 27 && dc >= 12 && rh >= 40 {
		const (
			c1 = -42.379
			c2 = 2.04901523
			c3 = 10.14333127
			c4 = -0.22475541
			c5 = -6.83783 / 1000
			c6 = -5.481717 / 100
			c7 = 1.22874 / 1000
			c8 = 8.5282 / 10000
			c9 = -1.99 / 1000000
		)

		hff := c1 + c2*tf + c3*rh + c4*tf*rh + c5*math.Pow(tf, 2) + c6*math.Pow(rh, 2) +
			c7*math.Pow(tf, 2)*rh + c8*tf*math.Pow(rh, 2) + c9*math.Pow(tf, 2)*math.Pow(rh, 2)
		hf = (hff - 32) / 1.8
	} else {
		hf = tc
	}

	return wc, hf, rh

}

/* Search airport in airport data list */
func searchAirport(adFile []string, searchText string) {
	list := ""
	for _, line := range adFile {
		if strings.Index(strings.ToUpper(line), strings.ToUpper(searchText)) >= 0 {
			list += fmt.Sprintf("%s\n", strings.Replace(line, ";", " ", -1))
		}
	}
	if list != "" {
		fmt.Println(list)
	} else {
		fmt.Println("\nNothing found\n")
	}
}
