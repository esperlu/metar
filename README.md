# metar


This Go program is a console (terminal) application that retrieves aviation METARs and TAFs for a given list of airports and other weather stations. Special care has been taken to optimize execution speed by using goroutines for concurrent data retrieval of METARs and TAFs.

In addition to the METAR messages, the program computes the wind chill factor, heat index, and relative humidity when applicable.

## Options


```
  -n  <n>         Set the number of METARs to print per station. N: 1 to 70
  -s  <string>    Search for an airport by IATA/ICAO code
  -lc <string>    List all countries with their ISO code (<string> may be empty)
  -la <string>    List all airports in one or more countries (ISO country codes)
  -t  <t>         Set connection timeout. T: 1 to 10
  -r              Print raw data without additional factors
  -m              METARs only (mutually exclusive with -f)
  -f              TAFs only (mutually exclusive with -m)
  -h              Show this help screen
```


## Retrieve messages for a list of stations (IATA or ICAO codes)

```sh
$ metar cdg FACT
```
(Case insensitive)

Example output:

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


At the end of each METAR, the three values in brackets are the computed:  
```[ wind chill factor | heat index | relative humidity % ]```


## Find the IATA/ICAO airport code for an airport

```sh
$ metar -s munich
$ metar -s "new york"
```


## List ISO country codes

```sh
$ metar -lc
$ metar -lc africa
```


## List all airports in specified countries using ISO country codes

```sh
$ metar -la fr
$ metar -la it pt es uk
```

## Help screen

```sh
$ metar -h
```

## Installation


You will need to compile the sources using the Go tools. Follow this [how-to](https://go.dev/doc/tutorial/compile-install) to get started. Compilation is lightning fast, and [cross-compilation](http://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5) is easy.

### Install the latest Go for your platform

* Easy way: install the [latest version binaries](https://golang.org/dl/) or use your distro's package manager (may not always be the latest version)
* Advanced: [compile Go from source](https://golang.org/doc/install/source)

### Get this metar repository

1. Run the following command to install the metar repo in the directory defined by your `GOPATH` environment variable:
  ```sh
  go get github.com/esperlu/metar
  ```
2. Navigate to the local sources: `<GOPATH>/src/github.com/esperlu/metar`
3. Try it out: run the following command to get the METAR weather reports for Brussels (BRU, BE) and New York (JFK, US):
  ```sh
  go run metar.go bru jfk
  ```
4. If successful, compile the metar sources and data:
   * To compile the binary and save it in the current directory:
    ```sh
    go build metar.go
    ```
   * To compile the binary and install it in the binary folder defined by the `GOBIN` environment variable:
    ```sh
    go install metar.go
    ```
    This will make the binary accessible and executable system-wide.


## Utilities

The airport list and METAR stations list are hardcoded for speed. However, these lists are subject to change. To update the lists, run the `updateStations.go` program in the `util` directory. Then recompile the main program `metar.go` to hardcode the updated lists.



## Bug report

Rough edges are not excluded. Please [report any bugs](https://github.com/esperlu/metar/issues).


## Credits

METAR weather messages are retrieved from NOAA's aviationweather.gov in real time.
METAR stations list and names are compiled from:
* [aviationweather.gov](https://aviationweather.gov/data/cache/metars.cache.csv.gz)
* [ourairports.com](https://ourairports.com/data/airports.csv)

----
#### (c) Jean-Luc Lacroix
