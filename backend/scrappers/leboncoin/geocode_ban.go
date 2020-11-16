package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Geometry ...
type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// Property ...
type Property struct {
	Name     string `json:"name"`
	City     string `json:"city"`
	Citycode string `json:"citycode"`
	LocType  string `json:"type"`
	Label    string `json:"label"`
}

// Feature ...
type Feature struct {
	Type       string `json:"type"`
	Geometry   `json:"geometry"`
	Properties Property `json:"properties"`
}

// ToString returns meta to string in JSON format
func (f *Feature) ToString() (*string, error) {
	var err error
	var encode []byte
	if encode, err = json.Marshal(&f); err != nil {
		return nil, err
	}
	str := string(encode)
	return &str, nil
}

// FeatureCollection ...
type FeatureCollection struct {
	Type     string    `json:"type"`
	Version  string    `json:"version"`
	Features []Feature `json:"features"`
}

const (
	// BanAPIURL ...
	BanAPIURL = "https://api-adresse.data.gouv.fr"
	// BanSearch ...
	BanSearch = "search"
)

// DoSearch ...
func DoSearch(query string) (*FeatureCollection, error) {
	var err error
	var apiResp *http.Response
	var resp FeatureCollection
	var url = BanAPIURL + "/" + BanSearch

	if apiResp, err = http.DefaultClient.Get(url + "/?q=" + strings.ReplaceAll(query, " ", "+")); err != nil {
		return nil, err
	}
	if apiResp.StatusCode == http.StatusOK {
		if err = json.NewDecoder(apiResp.Body).Decode(&resp); err != nil {
			return nil, err
		}
	} else {
		fmt.Printf("Failed to do search for %s got status %v\n", query, apiResp.StatusCode)
	}
	return &resp, nil
}
