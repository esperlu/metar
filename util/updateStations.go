// This program updates the METAR stations list.
//
// It compiles lists from NOAA and ourairports.com web sites.
// Change `dataFile` assigment here below to the actual path where your `ad_list.go` file lives.
// Once updated, you will need to recompile or run metar.go to hardcode the updated stations into the metar binary.
//
// Warning: do not change the var declaration `var AdList` in ad_list.go as it works as a marker for this program.
//
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

type st struct {
	Name, Icao, Iata, Country string
}

const (
	dataFile       string = "/home/jeanluc/golang/src/jeanluc/metarDEV/data/ad_list.go"
	noaaURL        string = "https://www.aviationweather.gov/docs/metar/stations.txt"
	ourairportsURL string = "https://ourairports.com/data/airports.csv"
)

func main() {

	stationsNOAA := make(map[string]st)
	stationsLIST := make(map[string]st)
	var wg sync.WaitGroup

	start := time.Now()

	// Get and process NOAA file
	wg.Add(1)
	go func() {
		defer wg.Done()
		s, err := wget(noaaURL, 5)
		if err != nil {
			fmt.Printf("\n file %s %v\n\n", noaaURL, err)
			os.Exit(1)
		}

		lines := strings.Split(s, "\n")
		// Traverse lines (skip 42 lines)
		for _, l := range lines {
			// Skip titles
			if len(l) != 83 {
				continue
			}
			// Skip stations without METAR reports and stations witout ICAO code
			// Note: iata code is in fact the FAA code. Not usable.
			if string(l[62]) == "X" && strings.TrimSpace(l[20:24]) != "" {
				stationsNOAA[strings.TrimSpace(l[20:24])] = st{
					Name:    strings.TrimSpace(l[3:20]),
					Icao:    strings.TrimSpace(l[20:24]),
					Country: l[81:83],
				}
			}

		}
		if len(stationsNOAA) == 0 {
			fmt.Printf("\n %s file not valid.\n No valid record found.\n\n", noaaURL)
			os.Exit(1)
		}
	}()

	// get and process ourairport file
	wg.Add(1)
	go func() {
		defer wg.Done()
		s, err := wget(ourairportsURL, 5)
		if err != nil {
			fmt.Printf("\n file %s %v\n\n", ourairportsURL, err)
			os.Exit(1)
		}

		// make a new reader from string `s`
		r := csv.NewReader(strings.NewReader(s))
		// required number of fields
		r.FieldsPerRecord = 18
		// parse all csv lines
		lines, err := r.ReadAll()
		if err != nil {
			fmt.Printf("\n %s file not valid. Expecting 18 fields\n %v\n\n", ourairportsURL, err)
			os.Exit(1)
		}
		for _, l := range lines {
			// remove possible remaining `"` in airport name
			l[3] = strings.ReplaceAll(l[3], "\"", "")
			stationsLIST[l[1]] = st{
				Name:    l[3] + " (" + l[10] + ")",
				Icao:    l[1],
				Iata:    l[13],
				Country: l[8],
			}
		}
	}()

	wg.Wait()

	// print stations from NOAA with details taken in LIST if exists (else use NOAA names and codes)
	var final []st
	for _, vNOAA := range stationsNOAA {
		if vLIST, ok := stationsLIST[vNOAA.Icao]; ok {
			final = append(final, st{Iata: vLIST.Iata, Icao: vLIST.Icao, Name: vLIST.Name, Country: vLIST.Country})
		} else {
			final = append(final, st{Icao: vNOAA.Icao, Name: vNOAA.Name, Country: vNOAA.Country})
		}

	}

	// Sort `final` on Icao code
	sort.Slice(final, func(i, j int) bool { return final[i].Icao < final[j].Icao })

	// Add final to data file
	// Append to metarDEV/data/ad_list.go
	f, err := os.OpenFile(dataFile, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// delete lines after `var AdList = []string{`
	scanner := bufio.NewScanner(f)
	var bytesRead int
	// find line containing `var AdList`
	for scanner.Scan() {
		// bytes read so far (+1 for \n)
		bytesRead += len(scanner.Text()) + 1
		if strings.Contains(scanner.Text(), "var AdList") {
			break
		}
	}
	f.Truncate(int64(bytesRead))

	// store all records in a file
	for _, v := range final {
		if _, err := f.WriteString(fmt.Sprintf("\t\"%s;%s;%s;%s\",\n", v.Icao, v.Iata, v.Name, v.Country)); err != nil {
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
	FUNC
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

	// return string(wgetAnswer), nil
	// return output (after removing trailing \n)
	return string(wgetAnswer[:len(wgetAnswer)-1]), nil
}
