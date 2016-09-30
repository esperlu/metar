# metar

This Golang program fetch the aviation METAR's and TAF's for a given list of airports.

##Installation

You will have to compile the sources using the golang tools. Follow this [howto](https://golang.org/doc/code.html) to get you started. The compilation is lightning fast and the [cross-compilation](http://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5) easy.

##Usage

**Retrieve messages for a list of stations (IATA or ICAO codes):**

```$ metar lhr jfk bru uudd```


**Find the IATA/ICAO airport code for an airport**

```$ metar -s boston```

**Help screen:**

```$ metar -h```

##Credits
These aviation weather messages are retrieved from http://aviationweather.org
