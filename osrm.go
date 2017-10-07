package osrm

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const API_VERSION = "v1"

const (
	RouteService = "route"
	MatchService = "match"
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
	return oc.buildUrl(RouteService, options)
}

func (oc *Client) buildMatchUrl(options RouteOptions) (string, error) {
	return oc.buildUrl(MatchService, options)
}

func (oc *Client) buildUrl(service string, options RouteOptions) (string, error) {
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
	url := fmt.Sprintf("%s/%s/%s/%s/%s", oc.RootURL, service, API_VERSION, options.Profile, path)
	return url, nil
}

func (oc *Client) processOptions(q *url.Values, options RouteOptions) {
	if options.Alternatives != "" {
		q.Add("alternatives", options.Alternatives)
	}
	if options.Steps != "" {
		q.Add("steps", options.Steps)
	}

	if options.Annotations != "" {
		q.Add("annotations", options.Annotations)
	}
	if options.Geometries != "" {
		q.Add("geometries", options.Geometries)
	}
	if options.Overview != "" {
		q.Add("overview", options.Overview)
	}
	if options.ContinueStraight != "" {
		q.Add("continue_straight", options.ContinueStraight)
	}
}

func (oc *Client) RouteTo(options RouteOptions) ([]byte, error) {
	return oc.Query(MatchService, options)
}

func (oc *Client) Match(options RouteOptions) ([]byte, error) {
	return oc.Query(MatchService, options)
}

func (oc *Client) Query(service string, options RouteOptions) ([]byte, error) {
	url, err := oc.buildUrl(service, options)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	oc.processOptions(&q, options)

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