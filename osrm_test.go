package osrm

import "testing"
import "net/http"

func TestBuildRouteUrl(t *testing.T) {
	c := NewClient("http://example.com")
	url, err := c.buildRouteUrl(RouteOptions{
		Profile: "driving",
		Locations: []Location{
			Location{
				Lat: 42.878473,
				Lon: 74.595532,
			},
			Location{
				Lat: 42.873764,
				Lon: 74.587990,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if url != "http://example.com/route/v1/driving/74.595532,42.878473;74.587990,42.873764" {
		t.Error(url)
	}
}

func TestBuildMatchUrl(t *testing.T) {
	c := NewClient("http://example.com")
	url, err := c.buildMatchUrl(RouteOptions{
		Profile: "driving",
		Locations: []Location{
			Location{
				Lat: 42.878473,
				Lon: 74.595532,
			},
			Location{
				Lat: 42.873764,
				Lon: 74.587990,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if url != "http://example.com/match/v1/driving/74.595532,42.878473;74.587990,42.873764" {
		t.Error(url)
	}
}

func TestQueryIncludesOnlyRouteOptionsThatAreExplicitlySet(t *testing.T) {
	c := NewClient("http://example.com")
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	c.processOptions(&q, RouteOptions{
		Steps: "true",
	})
	if q.Encode() != "steps=true" {
		t.Error(q.Encode())
	}
}
