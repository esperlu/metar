## UTIL: update METAR stations
`updateStations.go` updates the METAR stations list.

It combines lists from [aviationweather.gov](https://www.aviationweather.gov/docs/metar/stations.txt) and [ourairports.com](https://ourairports.com/data/airports.csv) web sites.

Change the `dataFile` variable to the actual path where your `ad_list.go` file lives, typically `path/to/the/sources/data/ad_list.go`. The program will insert the updated data into `ad_list.go`

```
const (
	// Change dataFile path to where this ad_list.go lives
	dataFile       string = "/home/jeanluc/golang/src/jeanluc/metar/data/ad_list.go"
	noaaURL        string = "https://www.aviationweather.gov/docs/metar/stations.txt"
	ourairportsURL string = "https://ourairports.com/data/airports.csv"
)
```

Once updated, you will need to recompile or run `metar.go` to hardcode the updated stations into the metar binary.

__Warning__: do not change the var declaration `var AdList` in `ad_list.go` as it works as a marker for this program.
