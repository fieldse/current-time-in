// Get the current time in a given city
package citylookup

import (
	"github.com/fieldse/current-time-in/pkg/logger"
	"github.com/fieldse/current-time-in/shared"

	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

var Logger = logger.Logger

// CityData represents a single entry from the cities data table
type CityData struct {
	City      string  `json:"city"`
	CityAscii string  `json:"city_ascii"`
	Lat       float32 `json:"lat"`
	Lng       float32 `json:"lng"`
	Pop       float32 `json:"pop"`
	Country   string  `json:"country"`
	Iso2      any     `json:"iso2"` // these sometimes will be "-99" for no obvious reason
	Iso3      string  `json:"iso3"`
	Province  string  `json:"province"`
	Timezone  string  `json:"timezone"`
}

// tzCodeFor returns a full timezone name for cities matching the given city/country name.
// This function fuzzy-matches based on the city name, and attempts to return intelligently
// from cities matching the name.
// eg: "London", "London UK", "London United Kingdom" should all return the same result.
// This will return a slice of results for cities which exist in multiple countries.
func tzCodeFor(_ string) ([]string, error) {
	return []string{""}, shared.ErrorNotImplementedError{}
}

// readCityData reads the raw data from cityMap.json to bytes
func readCityData() ([]byte, error) {
	b, err := os.ReadFile(path.Join(".", "data", "cityMap.json"))
	if err != nil {
		log.Printf("failed to read cityMap.json: %v", err)
	}
	return b, err
}

// loadCityData loads the city data from cityMap.json
func loadCityData() ([]CityData, error) {
	var cityData []CityData
	b, err := readCityData()
	if err != nil {
		return []CityData{}, fmt.Errorf("failed to read city data: %w", err)
	}
	err = json.Unmarshal(b, &cityData)
	if err != nil {
		return []CityData{}, fmt.Errorf("error parsing city data: %w", err)
	}
	return cityData, nil
}

// findCityExact finds a single city matching exactly by name
func findCityExact(rows []CityData, s string) (CityData, error) {
	for _, r := range rows {
		if r.City == s {
			return r, nil
		}
	}
	return CityData{}, fmt.Errorf("city not found: %s", s)
}