## UTIL: update METAR stations
This program updates the METAR stations list.

It compiles lists from aviationweather.gov and ourairports.com web sites.

Change `dataFile` assigment here below to the actual path where your `ad_list.go` file lives.
Once updated, you will need to recompile or run `metar.go` to hardcode the updated stations into the metar binary.

__Warning__: do not change the var declaration `var AdList` in `ad_list.go` as it works as a marker for this program.
