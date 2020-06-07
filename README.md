# metar

This go program fetch the aviation METAR's and TAF's for a given list of airports in console (terminal) mode. Special care has been put into execution speed using the goroutines for simultaneous data retrieval of METAR's and TAF's.

As an addition to the METAR messages, Wind Chill factor, Heat Factor and Relative Humidity are computed when applicable.


**Retrieve messages for a list of stations (IATA or ICAO codes):**

```$ metar lhr jfk bru uudd```

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

**Find the IATA/ICAO airport code for an airport**

```
$ metar -s munich
$ metar -s new york
```

**Options**

```
-r      raw output (no airport header and no additional factors)
-n <N>  number of metar to print (min. 1, max. 70)
-t <T>  connection timeout in sec. (min.1 max 10)
```
Example:   ```$ metar -r -n 15 bru jfk``` prints the latest 15 metars in raw format


**Help screen:**

```$ metar -h```

## Installation

You will have to compile the sources using the golang tools. Follow this [howto](https://golang.org/doc/code.html) to get you started. The compilation is lightning fast and the [cross-compilation](http://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5) easy.

You can also use the [binaries](https://github.com/esperlu/metar/tree/master/binaries) that I have cross-compiled for your convenience.

OS | Architecture | Direct link
------------ | ----------- | -------------
Linux 64 bit| amd64 | [download](https://github.com/esperlu/metar/blob/master/binaries/linux/amd64/metar?raw=true)
Linux | arm v5 | [download](https://github.com/esperlu/metar/blob/master/binaries/linux/arm5/metar?raw=true)
Linux RPi-1 (Zero, A, A+, B, B+) 32-bit| arm v6 | [download](https://github.com/esperlu/metar/blob/master/binaries/linux/arm6/metar?raw=true)
Linux RPi-2B 32-bit | arm v7 | [download](https://github.com/esperlu/metar/blob/master/binaries/linux/arm7/metar?raw=true)
Darwin (OS X) | amd64 | [download](https://github.com/esperlu/metar/blob/master/binaries/darwin/amd64/metar?raw=true)
Windows 32-bit| 386 | [download](https://github.com/esperlu/metar/blob/master/binaries/windows/amd64/metar.exe?raw=true)
Windows 64-bit| amd64 | [download](https://github.com/esperlu/metar/blob/master/binaries/windows/amd64/metar.exe?raw=true)


Once downloaded Linux and Mac users will have to make that file executable: `$ chmod +x metar`

To run it (Linux and Mac): `$ ./metar`

## Bug report
Rough edges are not excluded. Please [report](https://github.com/esperlu/metar/issues) any bugs.

## Credits
These aviation weather messages are retrieved from http://aviationweather.gov in real time.

## Author
Jean-Luc Lacroix
