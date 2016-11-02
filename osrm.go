package osrm

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	RoutePath = "route/v1"
	MatchPath = "match/v1"
)

type (
	Client struct {
		Client  *http.Client
		RootURL string
	}
	RouteOptions struct {
		Profile          string
		Alternatives     string // accepts true or false
		Steps            string // accepts true or false
		Annotations      string // accepts true or false
		Geometries       string // accepts polyline or geojson
		Overview         string // accepts simplified, full or false
		ContinueStraight string // accepts default, true or false
		Locations        []Location
	}
	Location struct {
		Lon float64
		Lat float64
	}
)

// NewClient initializes a not client for OSRM backend with gived rootURL
func NewClient(rootURL string) *Client {
	return &Client{
		Client:  &http.Client{},
		RootURL: rootURL,
	}
}

func (oc *Client) buildRouteUrl(options RouteOptions) (string, error) {
	if options.Profile == "" {
		return "", errors.New("Profile can't be blank")
	}
	if len(options.Locations) < 2 {
		return "", errors.New("Should be more than 2 locations to build route")
	}
	var locations []string
	for _, location := range options.Locations {
		locations = append(locations, fmt.Sprintf("%f,%f", location.Lon, location.Lat))
	}
	path := strings.Join(locations, ";")
	url := fmt.Sprintf("%s/%s/%s/%s", oc.RootURL, RoutePath, options.Profile, path)
	return url, nil
}

func (oc *Client) RouteTo(options RouteOptions) ([]byte, error) {
	url, err := oc.buildRouteUrl(options)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("alternatives", options.Alternatives)
	q.Add("steps", options.Steps)
	q.Add("annotations", options.Annotations)
	q.Add("geometries", options.Geometries)
	q.Add("overview", options.Overview)
	q.Add("continue_straight", options.ContinueStraight)
	req.URL.RawQuery = q.Encode()
	resp, err := oc.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
