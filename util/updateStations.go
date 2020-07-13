// This program updates the METAR stations list.

// It compiles lists from NOAA and ourairports.com web sites.
// Change `dataFile` assigment here below to the actual path where your `ad_list.go` file lives.
// Once updated, you will need to recompile or run metar.go to hardcode the updated stations into the metar binary.

// Warning: do not change the var declaration `var AdList` in ad_list.go as it works as a marker for this program.
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

// station stores METAR station details
type station struct {
	name, icao, iata, country string
}

const (
	dataFile       string = "/home/jeanluc/golang/src/jeanluc/metar/data/ad_list.go"
	noaaURL        string = "https://www.aviationweather.gov/docs/metar/stations.txt"
	ourairportsURL string = "https://ourairports.com/data/airports.csv"
)

func main() {

	stationsNOAA := make(map[string]station)
	stationsLIST := make(map[string]station)
	var wg sync.WaitGroup

	start := time.Now()

	// Get and process NOAA file
	wg.Add(1)
	go func() {
		defer wg.Done()
		s, err := wget(noaaURL, 5)
		if err != nil {
			log.Fatalf("\n\n file %s\n %v\n\n", noaaURL, err)
		}

		lines := strings.Split(s, "\n")
		// Traverse lines
		for _, l := range lines {
			// Skip titles, non METAR stations and stations w/o icao code
			if len(l) != 83 || string(l[62]) != "X" || strings.TrimSpace(l[20:24]) == "" {
				continue
			}
			// Store station ident in map[icao]station
			// Note: iata code is in fact the FAA code. Not usable.
			stationsNOAA[strings.TrimSpace(l[20:24])] = station{
				name:    strings.TrimSpace(l[3:20]),
				icao:    strings.TrimSpace(l[20:24]),
				country: l[81:83],
			}
		}
		if len(stationsNOAA) == 0 {
			log.Fatalf("\n\n %s file not valid.\n No valid record found.\n\n", noaaURL)
		}
	}()

	// get and process ourairport csv file
	wg.Add(1)
	go func() {
		defer wg.Done()
		s, err := wget(ourairportsURL, 5)
		if err != nil {
			log.Fatalf("\n\n file %s\n %v\n\n", ourairportsURL, err)
		}

		// make a new reader from string `s`
		r := csv.NewReader(strings.NewReader(s))
		// required number of fields
		r.FieldsPerRecord = 18
		// parse all csv lines
		lines, err := r.ReadAll()
		if err != nil {
			log.Fatalf("\n\n %s file not valid. Expecting 18 fields\n %v\n\n", ourairportsURL, err)
		}
		for _, l := range lines {
			// remove possible remaining `"` in airport name
			l[3] = strings.ReplaceAll(l[3], "\"", "")
			stationsLIST[l[1]] = station{
				name:    l[3] + " (" + l[10] + ")",
				icao:    l[1],
				iata:    l[13],
				country: l[8],
			}
		}
	}()

	// wait for all go routines to finish
	wg.Wait()

	// print stations from NOAA with details taken in LIST if exists (else use NOAA names and codes)
	var final []station
	for _, vNOAA := range stationsNOAA {
		if vLIST, ok := stationsLIST[vNOAA.icao]; ok {
			final = append(final, station{iata: vLIST.iata, icao: vLIST.icao, name: vLIST.name, country: vLIST.country})
		} else {
			final = append(final, station{icao: vNOAA.icao, name: vNOAA.name, country: vNOAA.country})
		}
	}

	// Sort `final` on icao code
	sort.Slice(final, func(i, j int) bool { return final[i].icao < final[j].icao })

	// Add `final` to data file `ad_list.go`
	f, err := os.OpenFile(dataFile, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// delete lines after `var AdList = []string{`
	scanner := bufio.NewScanner(f)
	var bytesRead int
	// find line containing `var AdList` and count bytes read (+1 for \n)
	for scanner.Scan() {
		t := scanner.Text()
		bytesRead += len(t) + 1
		if strings.Contains(t, "var AdList") {
			break
		}
	}
	f.Truncate(int64(bytesRead))

	// store all records in file
	for _, v := range final {
		if _, err := f.WriteString(fmt.Sprintf("\t\"%s;%s;%s;%s\",\n", v.icao, v.iata, v.name, v.country)); err != nil {
			log.Fatal(err)
		}
	}
	if _, err := f.WriteString("}\n"); err != nil {
		log.Fatal(err)
	}

	fmt.Printf(
		"\n %d records updated in %.3f sec.\n you can now recompile metar.go with the updated stations.\n\n",
		len(final),
		time.Since(start).Seconds(),
	)

}

/*
	FUNC's
*/

func wget(urlString string, wgetTimeout int) (string, error) {

	timeout := time.Duration(wgetTimeout) * time.Second
	client := http.Client{Timeout: timeout}

	// Get page and check for error (timeout, http ...)
	res, err := client.Get(urlString)

	// unwrap url.error and check error and its type
	if e, ok := err.(*url.Error); ok && e.Timeout() {
		return "", e.Unwrap()
	} else if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// if not HTTP 200 OK in response header
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf(" Page not found")
	}

	// return res.Body, nil

	wgetAnswer, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// return output (after removing trailing \n)
	return string(wgetAnswer[:len(wgetAnswer)-1]), nil
}
