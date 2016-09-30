# metar

This Golang program gets the aviation METAR's and TAF's for a givel list of airports.

##Installation

You will have to compile the sources using the golang tools.

##Usage

**Retrieve messages for a list of stations (IATA or ICAO codes):**
$ metar lhr jfk bru uudd


**Find the ICAO/IATA airport code for an airport**
$ metar -s boston

**Help screen:**
$ metar -h
