package amap_coordinate

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func NewClient (apiKey string) *SearchClient {
	return &SearchClient{
		client: &http.Client{
			Timeout:       time.Minute,
		},
		apiKey: apiKey,
	}
}

func (c *SearchClient) Search(name string, types string, city string) (SearchPOIResult, error) {
	u := url.Values{}
	u.Add("key", c.apiKey)
	u.Add("keywords", name)
	u.Add("types", types)
	u.Add("city", city)

	uu := fmt.Sprintf("https://restapi.amap.com/v3/place/text?%s", u.Encode())
	spew.Dump(uu)

	response, err := http.Get(uu)
	if err != nil {
		return SearchPOIResult{}, err
	}
	if response.StatusCode != 200 {
		spew.Dump(response.Body)
		return SearchPOIResult{}, fmt.Errorf("not 200 status code: %v, %v", response.StatusCode, response)
	}

	var r SearchResponse
	err = json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		return SearchPOIResult{}, err
	}
	spew.Dump(r)

	location := r.POIs[0].Location
	segments := strings.Split(location, ",")
	return SearchPOIResult{
		Latitude:  segments[0],
		Longitude: segments[1],
	}, nil
}
