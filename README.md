# metar

This Go program is a console (terminal) mode program that retrieves aviation METARs and TAFs for a given list of airports and other weather stations. Special care has been taken to optimize execution speed by using goroutines for concurrent data retrieval of METARs and TAFs.

In addition to the METAR messages, wind chill factor, heat index, and relative humidity are computed when applicable.

## Options

```
  -n  <n>         Set number of Metars to print per station. N: 1 to 70
  -s  <string>    Search IATA/ICAO code for an airport
  -lc <string>    List all countries with their ISO code (<string> may be empty)
  -la <string>    List all airports in one or more countries (ISO country codes)
  -t  <t>         Connection timeout T: 1 to 10
  -r              Print raw data w/o the additional factors
  -m              METARS only (mutually exclusive with -f)
  -f              TAFS only (mutually exclusive with -m)
  -h              This help screen
```

## Retrieve messages for a list of stations (IATA or ICAO codes)

```$ metar cdg FACT``` (Case insensitive)

The output looks like this:

```
LFPG (CDG) Charles de Gaulle International Airport (Paris), France FR (EU)
LFPG 060600Z 03003KT CAVOK 19/10 Q1017 NOSIG [19 19 56%]
LFPG 060530Z 02004KT CAVOK 17/09 Q1017 NOSIG [17 17 59%]
LFPG 060500Z 02006KT CAVOK 17/09 Q1017 NOSIG [17 17 59%]
LFPG 060430Z 03005KT CAVOK 16/09 Q1017 NOSIG [16 16 63%]
TAF 060500Z 0606/0712 04005KT CAVOK TX36/0712Z TN17/0606Z

FACT (CPT) Cape Town International Airport (Cape Town), South Africa ZA (AF)
FACT 060600Z VRB02KT 9999 FEW030 BKN045 07/06 Q1030 NOSIG [7 7 93%]
FACT 060500Z VRB01KT 9999 -RA FEW035 BKN045 07/06 Q1030 NOSIG [7 7 93%]
FACT 060400Z 05004KT 9999 FEW035 BKN050 06/05 Q1029 NOSIG [5 6 93%]
TAF 060400Z 0606/0712 03005KT 9999 SCT020 BKN030 TX17/0712Z TN08/0706Z FM061300 18013KT CAVOK FM062300 VRB03KT CAVOK


```

At the end of the METAR's, the three values between brackets are the computed  ```[ wind chill factor | heat factor | relative humidity % ]```

## Find the IATA/ICAO airport code for an airport

```
$ metar -s munich
$ metar -s new york
```

## List ISO country codes

```
$ metar -lc
$ metar -lc africa

```

## List all country airports using ISO country codes
```
$ metar -la fr
$ metar -la it pt es uk
```
## Help screen

```$ metar -h```

## Installation

You will need to compile the sources using the Golang tools. Follow this [howto](https://go.dev/doc/tutorial/compile-install) to get started. The compilation is lightning fast and [cross-compilation](http://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5) is easy.

If you do not want to compile it yourself, you can skip the steps below and download the [binaries](#binaries)

### Install the latest Go for your plateform

* Easy way: install the [latest version binaries](https://golang.org/dl/) or use your distro package (not always the latest version)
* Less easy way: [compile Go from source](https://golang.org/doc/install/source)

### Get this metar repo

1. Run the following command to install the metar repo in the directory defined in your `GOPATH` environment variable:  
`go get github.com/esperlu/metar` 
2. Navigate to the now local sources: `<GOPATH>/src/github.com/esperlu/metar`
3. Give it a try: run the following command to get the METAR weather reports for Brussels BRU (BE) and New York JFK (US):  
`go run metar.go bru jfk`
4. If successfull, compile the metar sources and data:
    * To compile the binary and save it in the current directory, run the following command:  
    `go build metar.go`
    * To compile the binary and install it in the binary folder defined in the `GOBIN` environment variable, run the following command:  
    `go install metar.go`  
    This will make the binary accessible and executable system wide.

## Utilities

The airport list and METAR stations list are hardcoded for the sake of speed. However, these lists are subject to change. To update the lists, run the `updateStations.go` program in the `util` directory. Then recompile the main program `metar.go` to hardcode the updated lists.


## Bug report
Rough edges are not excluded. Please [report](https://github.com/esperlu/metar/issues) any bugs.

## Credits
METAR weather messages are retrieved from NOAA's aviationweather.gov in real time.
METAR stations list and names are compiled from :
* [aviationweather.gov](https://www.aviationweather.gov/docs/metar/stations.txt)
* [ourairports.com](https://ourairports.com/data/airports.csv)

---
#### (c) Jean-Luc Lacroix
