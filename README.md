# metar

This go program fetch the aviation METAR's and TAF's for a given list of airports in console (terminal) mode.

##Installation

You will have to compile the sources using the golang tools. Follow this [howto](https://golang.org/doc/code.html) to get you started. The compilation is lightning fast and the [cross-compilation](http://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5) easy.

You can also use the [binaries](https://github.com/esperlu/metar/tree/master/binaries) that I have cross compiled.

##Usage

**Retrieve messages for a list of stations (IATA or ICAO codes):**

```$ metar lhr jfk bru uudd```

The output looks like this:

```UUDD (DME) Moscow (Domododevo), Russia
UUDD 301700Z 22004MPS 9000 -SHRA BKN014CB 15/12 Q1007 R14R/290095 R14L/290094 NOSIG [15 15 82%] 
UUDD 301630Z 21005MPS 9999 SCT033 15/12 Q1007 WS ALL RWY R14R/010095 R14L/290094 NOSIG [15 15 82%] 
UUDD 301600Z 23005MPS 200V260 9999 FEW020 15/12 Q1007 R14R/010095 R14L/290094 NOSIG [15 15 82%] 
UUDD 301530Z 22005MPS 190V250 9999 FEW020 15/11 Q1007 WS R14L R14R/010095 R14L/290094 NOSIG [15 15 77%] 
```

At the end of the METAR's, the three values between brackets are the computed  ```[ wind chill factor | heat factor | relative humidity % ]```

**Find the IATA/ICAO airport code for an airport**

```$ metar -s boston```

**Help screen:**

```$ metar -h```

##Credits
These aviation weather messages are retrieved from http://aviationweather.gov in real time.
