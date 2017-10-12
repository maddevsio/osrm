# osrm
OSRM Api wrapper for Go


# Examples
## Basic
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

## Strongly typed api response for match request
```go
package main

import (
	"github.com/maddevsio/osrm"
	"encoding/json"
	"log"
)

func main() {
	osrmapi := osrm.NewClient("http://127.0.0.1:5000")
	options := osrm.RouteOptions{
		Locations : []osrm.Location{
			{Lon: 23.746366, Lat: 37.957386},
			{Lon: 23.748366, Lat: 37.953386},
		},
		Profile : "driving",
		Steps : "true",
	}
	b, err := osrmapi.Match(options)
	log.Printf("%s", b)
	log.Printf("%v", err)

	var jsonData osrm.MatchResponse
	if err := json.Unmarshal(b, &jsonData); err != nil {
		panic(err)
	}
	log.Printf("%T", jsonData)
	log.Printf("%s", jsonData.Code)
	log.Printf("%v", err)
}
```