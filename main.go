// Get the current time in a given city
package main

import (
	"encoding/json"
	"errors"
	"os"
	"path"
)

// CityData represents a single entry from the cities data table
type CityData struct {
	City      string `json:"city"`
	CityAscii string `json:"city_ascii"`
	Lat       string `json:"lat"`
	Lng       string `json:"lng"`
	Pop       string `json:"pop"`
	Country   string `json:"country"`
	Iso2      string `json:"iso2"`
	Iso3      string `json:"iso3"`
	Province  string `json:"province"`
	Timezone  string `json:"timezone"`
}

// tzCodeFor returns a full timezone name for cities matching the given city/country name.
// This function fuzzy-matches based on the city name, and attempts to return intelligently
// from cities matching the name.
// eg: "London", "London UK", "London United Kingdom" should all return the same result.
// This will return a slice of results for cities which exist in multiple countries.
func tzCodeFor(s string) ([]string, error) {
	return []string{""}, errors.New("not implemented")
}

// loadCityData loads the city data from from cityMap.json
func loadCityData() ([]CityData, error) {
	var cityData []CityData
	b, err := os.ReadFile(path.Join(".", "data", "cityMap.json"))
	if err != nil {
		return []CityData{}, errors.New("error reading city data")
	}
	err = json.Unmarshal(b, &cityData)
	if err != nil {
		return []CityData{}, errors.New("error parsing city data")
	}
	return cityData, nil
}
