# metar

This go program fetch the aviation METAR's and TAF's for a given list of airports in console (terminal) mode.

##Installation

You will have to compile the sources using the golang tools. Follow this [howto](https://golang.org/doc/code.html) to get you started. The compilation is lightning fast and the [cross-compilation](http://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5) easy.

You can also use the [binaries](https://github.com/esperlu/metar/tree/master/binaries) that I have cross compiled.

##Usage

**Retrieve messages for a list of stations (IATA or ICAO codes):**

```$ metar lhr jfk bru uudd```
At the end of the METAR's, the three values between brackets, are computed to output the ```[ wind chill factor | heat factor | relative humidity % ]```

**Find the IATA/ICAO airport code for an airport**

```$ metar -s boston```

**Help screen:**

```$ metar -h```

##Credits
These aviation weather messages are retrieved from http://aviationweather.gov in real time.
