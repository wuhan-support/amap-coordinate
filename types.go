package amap_coordinate

import "net/http"

type SearchClient struct {
	client *http.Client

	apiKey string
}

type SearchResponsePOI struct {
	Location string `json:"location"`
}

type SearchResponse struct {
	POIs []SearchResponsePOI `json:"pois"`
}

type SearchPOIResult struct {
	Latitude string
	Longitude string
}
