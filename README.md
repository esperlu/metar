# metar

This go program is a console (terminal) mode program to fetch the aviation METAR's and TAF's for a given list of airports and other weather stations. Special care has been put into execution speed using the goroutines for concurrent data retrieval of METAR's and TAF's.

As an addition to the METAR messages, Wind Chill factor, Heat Factor and Relative Humidity are computed when applicable.

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

```$ metar lhr jfk BRU uudd``` (Case insensitive)

The output looks like this:

```
UUDD (DME) Moscow (Domododevo), Russia
UUDD 301700Z 22004MPS 9000 -SHRA BKN014CB 15/12 Q1007 R14R/290095 R14L/290094 NOSIG [15 15 82%]
UUDD 301630Z 21005MPS 9999 SCT033 15/12 Q1007 WS ALL RWY R14R/010095 R14L/290094 NOSIG [15 15 82%]
UUDD 301600Z 23005MPS 200V260 9999 FEW020 15/12 Q1007 R14R/010095 R14L/290094 NOSIG [15 15 82%]
UUDD 301530Z 22005MPS 190V250 9999 FEW020 15/11 Q1007 WS R14L R14R/010095 R14L/290094 NOSIG [15 15 77%]
TAF UUDD 301355Z 3015/0121 24007MPS 9999 BKN007 TX18/0111Z TN10/0102Z TEMPO 3015/3021 24013MPS 2100 -SHRA SCT004 BKN015CB

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

You will have to compile the sources using the golang tools. Follow this [howto](https://golang.org/doc/code.html) to get you started. The compilation is lightning fast and the [cross-compilation](http://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5) easy.

If you don't want to compile it yourself, skip the points below and go the [binaries](#binaries)

### Install the latest Go for your plateform

* Easy way: install the [latest version binaries](https://golang.org/dl/) or use your distro package (not always the latest version)
* Less easy way: [compile Go from source](https://golang.org/doc/install/source)

### Get this metar repo

1. `go get github.com/esperlu/metar` This will install this git repo in the directory defined in your `GOPATH` environment variable.
2. navigate to the now local sources `<GOPATH>/src/github.com/esperlu/metar`
3. give it a try: `go run metar.go bru jfk`
4. if successfull, compile the metar sources and data: `go build metar.go` or `go install metar.go` to install the binary in the binary folder defined in the `GOBIN` environment variable to make it accessible and executable system wide.


## Binaries

You can also use the [binaries](https://github.com/esperlu/metarBin) that I have cross-compiled for your convenience.

OS | Architecture | Direct link
------------ | ----------- | -------------
Linux 64 bit| amd64 | [download](https://github.com/esperlu/metarBin/blob/master/linux/amd64/metar?raw=true)
Linux | arm v5 | [download](https://github.com/esperlu/metarBin/blob/master/linux/arm5/metar?raw=true)
Linux RPi-1 (Zero, A, A+, B, B+) 32-bit| arm v6 | [download](https://github.com/esperlu/metarBin/blob/master/linux/arm6/metar?raw=true)
Linux RPi-2B 32-bit | arm v7 | [download](https://github.com/esperlu/metarBin/blob/master/linux/arm7/metar?raw=true)
Darwin (OS X) | amd64 | [download](https://github.com/esperlu/metarBin/blob/master/darwin/amd64/metar?raw=true)
Windows 32-bit| 386 | [download](https://github.com/esperlu/metarBin/blob/master/windows/386/metar.exe?raw=true)
Windows 64-bit| amd64 | [download](https://github.com/esperlu/metarBin/blob/master/windows/amd64/metar.exe?raw=true)

Once downloaded **Linux and Mac** users will have to make that file executable: `$ chmod +x metar`

To run it (Linux and Mac): `$ ./metar`

## Bug report
Rough edges are not excluded. Please [report](https://github.com/esperlu/metar/issues) any bugs.

## Credits
METAR weather messages are retrieved from http://aviationweather.gov in real time.
METAR station list and names are compiled from and aviationweather.gov and ourairports.com

---
#### (c) Jean-Luc Lacroix
