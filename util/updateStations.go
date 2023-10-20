// This program updates the METAR stations list.

// It compiles lists from NOAA and ourairports.com web sites.
// Change `dataFile` assigment here below to the actual path where your `ad_list.go` file lives.
// Once updated, you will need to recompile or run metar.go to hardcode the updated stations into the metar binary.

// Warning: do not change the var declaration `var AdList` in ad_list.go as it works as a marker for this program.
package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// station stores METAR station details
type station struct {
	name, icao, iata, country string
	lat, long                 float64
}

const (
	// CAUTION ! Change dataFile path to where this ad_list.go lives
	dataFile       string = "/home/jeanluc/golang/src/jeanluc/metar/data/ad_list.go"
	noaaURL        string = "https://aviationweather.gov/data/cache/metars.cache.csv.gz"
	ourairportsURL string = "https://ourairports.com/data/airports.csv"
)

func main() {

	stationsNOAA := make(map[string]station)
	stationsLIST := make(map[string]station)
	var wg sync.WaitGroup

	start := time.Now()

	// Get and process NOAA gziped file
	wg.Add(1)
	go func() {
		defer wg.Done()

		// get noaa station list (gz file)
		s, err := wget(noaaURL, 5)
		if err != nil {
			log.Fatalf("\nError getting %s\n  %s\n\n", noaaURL, err)
		}

		// decompress gz file
		s, err = gunzipGz(s)
		if err != nil {
			log.Fatalf("\nError decompressing data from %s\n  %s\n", noaaURL, err)
		}

		// make a new reader from string `s`
		r := csv.NewReader(strings.NewReader(s))
		r.LazyQuotes = true
		r.FieldsPerRecord = -1

		// get all records
		records, err := r.ReadAll()
		if err != nil {
			log.Fatalf("\nError reading records from %s\n  %s\n", noaaURL, err)
		}

		// Process noaa file skipping first 6 lines of file
		for _, l := range records[6:] {
			lt, _ := strconv.ParseFloat(l[3], 64)
			lg, _ := strconv.ParseFloat(l[4], 64)
			stationsNOAA[l[1]] = station{
				icao: l[1],
				lat:  lt,
				long: lg,
			}
		}
	}()

	// get and process ourairport csv file
	wg.Add(1)
	go func() {
		defer wg.Done()
		s, err := wget(ourairportsURL, 5)
		if err != nil {
			log.Fatalf("\nError getting %s\n  %s\n\n", ourairportsURL, err)
		}

		// make a new reader from string `s`
		r := csv.NewReader(strings.NewReader(s))

		// required number of fields
		r.FieldsPerRecord = 18

		// parse all csv lines
		lines, err := r.ReadAll()
		if err != nil {
			log.Fatalf("\n%s file not valid. Expecting 18 fields\n %v\n\n", ourairportsURL, err)
		}

		// Process parsed ourairport records
		for _, l := range lines[1:] {
			// remove possible remaining `"` in airport name
			l[3] = strings.ReplaceAll(l[3], "\"", "")
			l[3] = strings.ReplaceAll(l[3], "\\", "/")

			// parse and conv coord. set lat and long to 999.0 if err
			lt, errLt := strconv.ParseFloat(l[4], 64)
			lg, errLg := strconv.ParseFloat(l[5], 64)
			if errLt != nil || errLg != nil {
				lt, lg = 999.0, 999.0
			}
			// if `municipality` not empty
			if l[10] != "" {
				l[10] = strings.ReplaceAll(l[10], "\"", "")
				l[10] = strings.ReplaceAll(l[10], "\\", "/")
				l[10] = fmt.Sprintf(" (%s)", l[10])
			}
			stationsLIST[l[1]] = station{
				name:    l[3] + l[10],
				icao:    l[1],
				iata:    l[13],
				country: l[8],
				lat:     lt,
				long:    lg,
			}
		}
	}()

	// wait for all go routines to finish
	wg.Wait()

	// print ICAO codes from stationsNOAA with details taken in LIST only if exists in stationsLIST
	// (some ICAO codes in NOAA are incorrect)
	var final []station
	for _, vNOAA := range stationsNOAA {
		if vLIST, ok := stationsLIST[vNOAA.icao]; ok {
			// no valid lat-long in vLIST --> take NOAA lat-long
			if vLIST.lat == 999.0 {
				vLIST.lat = vNOAA.lat
				vLIST.long = vNOAA.long
			}
			final = append(final, vLIST)
		}
	}

	// Sort `final` on icao code
	sort.Slice(final, func(i, j int) bool { return final[i].icao < final[j].icao })

	// Add `final` to data file `ad_list.go`
	f, err := os.OpenFile(dataFile, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("\nError laoding %s\n %s", dataFile, err)
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

	err = f.Truncate(int64(bytesRead))
	if err != nil {
		log.Fatalf("Error truncating file %s\n %s\n\n", dataFile, err)
	}

	// store all records in file (option: add lat-long)
	for _, v := range final {
		if _, err := f.WriteString(fmt.Sprintf("\t\"%s;%s;%s;%s;%.3f;%.3f\",\n", v.icao, v.iata, v.name, v.country, v.lat, v.long)); err != nil {
			log.Fatalf("Error writing records into %s\n %s\n\n", dataFile, err)
		}
	}
	if _, err := f.WriteString("}\n"); err != nil {
		log.Fatalf("Error writing final '}\n' into %s\n %s\n\n", dataFile, err)
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
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// if not HTTP 200 OK in response header
	if res.StatusCode != http.StatusOK {
		err := new(url.Error)
		*err = url.Error{
			Op:  "Get",
			URL: urlString,
			Err: fmt.Errorf("HTTP error: %s", http.StatusText(res.StatusCode)),
		}
		return "", err
	}

	// return res.Body, nil
	wgetAnswer, err := io.ReadAll(res.Body)
	if err != nil {
		err := new(url.Error)
		*err = url.Error{
			Op:  "Get",
			URL: urlString,
			Err: fmt.Errorf("Error reading response body"),
		}
	}

	// return output (after removing trailing \n)
	return string(wgetAnswer[:len(wgetAnswer)-1]), nil
}

// httpError format (unwrap) url.Error returned by func wget
func httpError(err error) string {
	return fmt.Sprintf("\n\n Error: %s %s\n %v\n\n",
		err.(*url.Error).Op,
		err.(*url.Error).URL,
		err.(*url.Error).Err,
	)
}

// uncompress gziped string
// return guziped string and error
func gunzipGz(gzString string) (string, error) {

	reader := bytes.NewReader([]byte(gzString))
	gzReader, err := gzip.NewReader(reader)
	if err != nil {
		return "", err
	}

	// err != io.ErrUnexpectedEOF to get rid of the "unexpected EOF" normal error
	output, err := io.ReadAll(gzReader)
	if err != nil && err != io.ErrUnexpectedEOF {
		return "", err
	}

	return string(output), nil
}
