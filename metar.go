// metar is a go program to retrieve aviation METAR's and TAF's for a given list of airports in console (terminal) mode.

// Usage:
// Retrieve messages for a list of stations (IATA or ICAO codes):
//
//	$ metar lhr jfk bru uudd
//
// Find the IATA/ICAO airport code for an airport
//
//	$ metar -s munich
//	$ metar -s new york
//
// Help screen for other options:
//
//	$ metar -h
//
// Bug reports:
// https://github.com/esperlu/metar/issues

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math"
	"metar/data"
	"net/http"
	"net/url"
	"strings"
	"time"
	// "github.com/esperlu/metar/data"
)

// Constants to fetch Weather reports from aviationweather.com's new API
const (
	URLfmt      = "https://aviationweather.gov/api/data/%s?ids=%s&format=json&taf=true"
	urlMETARfmt = URLfmt + "&hours=%.1f"
	maxNbMETAR  = 70
	maxTIMEOUT  = 10
	ver         = "2.5.0"
)

// Initialize and parse flags
var (
	numberMetarFlag   = flag.Int("n", 4, "Set number of Metars to print per station. N 1 to 70.")
	searchFlag        = flag.Bool("s", false, "Search IATA/ICAO code for an airport")
	listCountriesFlag = flag.Bool("lc", false, "List all countries with ISO code")
	listAirportsFlag  = flag.Bool("la", false, "List all for one country (ISO code)")
	timeoutFlag       = flag.Int("t", 4, "Change the default timeout of 4 sec. to a maximum of 10 sec.")
	rawFlag           = flag.Bool("r", false, "Print raw data w/o the additional factors")
	metarOnlyFlag     = flag.Bool("m", false, "METARs only")
	tafOnlyFlag       = flag.Bool("f", false, "TAF (forecast) only")
)

// Custom usage function
func init() {
	flag.Usage = func() {
		fmt.Println(data.Help)
	}
	flag.Parse()
}

// Airport structure to hold airport data
type Airport struct {
	adIcaoCode, adIataCode, adName, adCountryCode, adCountryName, adLat, adLong string
}

// Metar structure to parse JSON response
type Metar struct {
	IcaoID string  `json:"icaoId"`
	Temp   float64 `json:"temp"`
	Dewp   float64 `json:"dewp"`
	Wspd   int     `json:"wspd"`
	RawOb  string  `json:"rawOb"`
	RawTaf string  `json:"rawTaf"`
	Name   string  `json:"name"`
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

	// Continents
	continents := map[string]string{}
	for _, line := range data.ContinentList {
		lineSplit := strings.Split(string(line), ";")
		continents[lineSplit[0]] = lineSplit[1]
	}

	// Countries
	countries := map[string][]string{}
	for _, line := range data.CountryList {
		lineSplit := strings.Split(string(line), ";")
		countries[lineSplit[0]] = []string{lineSplit[1], lineSplit[2]}
	}

	// Build airport info map
	// Table map[iata]=>(icao) and map[icao]=>(airport)+(details)
	iata2icao := map[string]string{}
	airportInfos := map[string]string{}
	for _, line := range data.AdList {
		lineSplit := strings.Split(string(line), ";")

		var airport Airport
		airport.adIcaoCode = lineSplit[0]
		airport.adIataCode = lineSplit[1]
		airport.adName = lineSplit[2]
		airport.adCountryCode = lineSplit[3]
		airport.adCountryName = countries[airport.adCountryCode][0]
		airport.adLat = lineSplit[4]
		airport.adLong = lineSplit[5]

		// empty iata code
		if airport.adIataCode == "" {
			airport.adIataCode = "---"
		}

		// store airportInfo in map
		airportInfos[airport.adIcaoCode] = fmt.Sprintf(
			"%s (%s) %s, %s %s (%s) %s %s",
			airport.adIcaoCode,
			airport.adIataCode,
			airport.adName,
			airport.adCountryName,
			airport.adCountryCode,
			countries[airport.adCountryCode][1],
			airport.adLat,
			airport.adLong,
		)

		// store IATA to ICAO codes in map
		iata2icao[airport.adIataCode] = airport.adIcaoCode
	}

	// search option -s
	if *searchFlag {
		fmt.Printf("\n%s\n", searchAirport(airportInfos, strings.Join(flag.Args(), " ")))
		fmt.Printf("  Processed in %.1f ms.\n\n", time.Since(startTotal).Seconds()*1000)
		return
	}

	// List country ISO codes
	if *listCountriesFlag {
		fmt.Printf("\n%s\n", listCountries(strings.Join(flag.Args(), " ")))
		fmt.Printf("  Processed in %.1f ms.\n\n", time.Since(startTotal).Seconds()*1000)
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
		fmt.Printf("\n%s\n", listAirports(flag.Args(), countries))
		fmt.Printf("  Processed in %.1f ms.\n\n", time.Since(startTotal).Seconds()*1000)
		return
	}

	// Parse airport list and convert IATA code (3 char.) to ICAO code (4 char.)
	var stations []string
	for _, v := range flag.Args() {
		v = strings.ToUpper(v)
		fmtNotFound := "\n\"%s\" not in the airport list. Try to run: metar -s %[1]s\n"

		switch {
		// check number of char. in arg
		case len(v) == 3:
			//  If found in icao2airportInfos map (some ICAO codes have only 3 characters).
			if _, ok := airportInfos[v]; ok {
				stations = append(stations, v)
				break
			}
			// if 3 char. and found in iata2icao map[iata]icao
			if _, ok := iata2icao[v]; ok {
				stations = append(stations, iata2icao[v])
			} else {
				fmt.Printf(fmtNotFound, v)
			}

			// if ICAO airport code (4 char.), lookup the icao2airportInfos map
		case len(v) >= 4:
			if _, ok := airportInfos[v]; ok {
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
		fmt.Print("\nNo airport given. Quitting...\n\n")
		return
	}

	// prepare the station list string for the URL
	stationList := strings.Join(stations, ",")

	// Start download timer
	startDownload := time.Now()

	// var wg sync.WaitGroup
	var metarResponse []byte

	// // get METARs routine (arg[4]--> 2 METARs per hour + 30 minutes)
	urlMETAR := fmt.Sprintf(urlMETARfmt, "metar", stationList, float32(*numberMetarFlag)/2+0.5)

	var errM error
	metarResponse, errM = wget(urlMETAR, *timeoutFlag)
	if errM != nil {
		fmt.Printf("\n  %s\n\n", checkWgetErr(errM))
		return
	}

	// Parse METAR response
	var reports []Metar
	if err := json.Unmarshal(metarResponse, &reports); err != nil {
		panic(err)
	}

	// Process METAR reports
	metars := map[string][]string{}
	taf := map[string]string{}
	for _, m := range reports {
		// If raw (-r) requested, don't compute wind chill factor, heat factor and relative humidity
		if *rawFlag {
			metars[m.IcaoID] = append(metars[m.IcaoID], m.RawOb)
			taf[m.IcaoID] = m.RawTaf
			continue
		}
		wc, hf, rh := computeFactors(float64(m.Wspd), m.Temp, m.Dewp)
		// Append computed factors to METAR report
		metars[m.IcaoID] = append(metars[m.IcaoID], m.RawOb+fmt.Sprintf(" [%.0f° %.0f° %.0f%%]", wc, hf, rh))
		taf[m.IcaoID] = m.RawTaf
	}

	// Print METARs and station TAF
	for id, metar := range metars {
		// Airport info header
		if !*rawFlag {
			fmt.Printf("++ %s ++\n", airportInfos[id])
		}
		// METAR reports
		if !*tafOnlyFlag {
			for i, report := range metar {
				// loop if maximum number of METAR si reached for one station
				if i > *numberMetarFlag {
					break
				}
				fmt.Println(report)
			}
		}
		// Latest TAF for that station
		if *metarOnlyFlag {
			continue
		}
		fmt.Println(taf[id] + "\n")
	}
	fmt.Printf("\nDownload: %.1f ms. Total: %.1f ms. (Ver. %s)\n", time.Since(startDownload).Seconds()*1000, time.Since(startTotal).Seconds()*1000, ver)
}

// Functions

// wget HTTP fetches URL content
func wget(urlString string, wgetTimeout int) ([]byte, error) {

	timeout := time.Duration(wgetTimeout) * time.Second
	client := http.Client{Timeout: timeout}

	// Get page and check for error (timeout, http ...)
	res, err := client.Get(urlString)

	// unwrap url.error and check error and its type
	if e, ok := err.(*url.Error); ok && e.Timeout() {
		return nil, fmt.Errorf("Timeout")
	} else if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Check HTTP status code
	httpErr := statusMessage(res.StatusCode)
	if httpErr != nil {
		return nil, httpErr
	}

	// Read body
	wgetAnswer, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return wgetAnswer, nil
}

// statusMessage returns a message based on HTTP status code
func statusMessage(statusCode int) error {
	switch statusCode {
	case http.StatusOK: // 200
		return nil
	case http.StatusCreated: // 201
		return fmt.Errorf("resource created successfully (status %d)", statusCode)
	case http.StatusAccepted: // 202
		return fmt.Errorf("request accepted for processing (status %d)", statusCode)
	case http.StatusNoContent: // 204
		return fmt.Errorf("no content returned for given station(s) (status %d)", statusCode)
	case http.StatusPartialContent: // 206
		return fmt.Errorf("partial content returned (status %d)", statusCode)
	default:
		if statusCode >= 200 && statusCode < 300 {
			return fmt.Errorf("successful response but unhandled status code: %d", statusCode)
		}
		return fmt.Errorf("unexpected status code: %d", statusCode)
	}
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

	// wind from kts to kmh
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

// searchAirport searches for airports in the airport data map
func searchAirport(icao2airportInfos map[string]string, query string) string {
	var sb strings.Builder
	query = strings.ToUpper(query)

	for _, info := range icao2airportInfos {
		if strings.Contains(strings.ToUpper(info), query) {
			sb.WriteString(fmt.Sprintf("  %s\n", info))
		}
	}

	if sb.Len() == 0 {
		return fmt.Sprintf("  No METAR station found for: '%s'\n", query)
	}

	return sb.String()
}

// listCountries lists all countries and ISO codes that match the search string
// Using strings.Builder for better performance
func listCountries(search string) string {
	var sb strings.Builder
	search = strings.ToUpper(search)

	for _, line := range data.CountryList {
		if strings.Contains(strings.ToUpper(line), search) {
			splitLine := strings.Split(line, ";")
			if len(splitLine) < 3 {
				continue // skip invalid lines
			}
			sb.WriteString(fmt.Sprintf(
				"  %s %s (%s)\n",
				splitLine[1], // country name
				splitLine[0], // country code
				splitLine[2], // ISO code or additional info
			))
		}
	}

	if sb.Len() == 0 {
		return fmt.Sprintf("  %s: Nothing found in country list.\n", search)
	}

	return sb.String()
}

// listAirports lists all airports for one or more countries (ISO country code)
// Using strings.Builder for better performance
func listAirports(countryCodes []string, code2country map[string][]string) string {
	var sb strings.Builder

	for _, code := range countryCodes {
		code = strings.ToUpper(code)

		// Check if country code exists
		countryInfo, ok := code2country[code]
		if !ok {
			return fmt.Sprintf(
				"\n  Unknown country code: %s.\n  Use option -lc [search text] to list country codes.\n",
				code,
			)
		}

		// Country title line
		header := fmt.Sprintf("\n  %s (%s)\n", countryInfo[0], countryInfo[1])
		var countryAdList strings.Builder

		// Loop through all airports for that country code
		airportCount := 0
		for _, line := range data.AdList {
			splitLine := strings.Split(line, ";")
			if len(splitLine) < 4 {
				continue // skip invalid lines
			}
			if code == splitLine[3] {
				countryAdList.WriteString(fmt.Sprintf(
					"  %s %-3s %s %s\n",
					splitLine[0],
					splitLine[1],
					splitLine[3],
					splitLine[2],
				))
				airportCount++
			}
		}
		// If at least one airport found, write to main string builder
		if airportCount > 0 {
			sb.WriteString(header)
			sb.WriteString(countryAdList.String())
			sb.WriteString(fmt.Sprintf("\n  %d airports found", airportCount))
		}
		if airportCount == 0 {
			return "  No airport found.\n"
		}
	}

	return sb.String()
}
