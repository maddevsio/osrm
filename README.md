# osrm
OSRM Api wrapper for Go

Example usage:

```go
package main

import (
	"github.com/maddevsio/osrm"
	"log"
)

func main() {
	osrmapi := osrm.NewClient("http://127.0.0.1:5000")
	options := osrm.RouteOptions{}
	options.Profile = "driving"
	options.Locations = []osrm.Location{
		{Lon: 13.388860, Lat: 52.517037},
		{Lat: 13.397634, Lon: 52.529407},
	}
	b, err := osrmapi.RouteTo(options)
	log.Printf("%s", b)
	log.Printf("%v", err)
}
```
