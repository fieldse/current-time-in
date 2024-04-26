// Get the current time in a given city
package citylookup

import (
	"strings"

	"github.com/fieldse/current-time-in/pkg/logger"
	"github.com/fieldse/current-time-in/shared"

	"fmt"
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

// findCityExact finds a single city matching exactly by name
func findCityExact(rows []CityData, s string) (CityData, error) {
	for _, r := range rows {
		if r.City == s {
			return r, nil
		}
	}
	return CityData{}, fmt.Errorf("city not found: %s", s)
}

// filterByCountry filters cities by country name on a case-insensitive substring match
// example: "united" would return "United States" and "United Kingdom"
func filterByCountry(rows []CityData, countryName string) []CityData {
	var filtered []CityData
	countryName = strings.ToLower(countryName)
	for _, r := range rows {
		s := strings.ToLower(r.Country)
		if strings.Contains(s, countryName) {
			filtered = append(filtered, r)
		}
	}
	return filtered
}

// filterByTZCode filters the cities data by fuzzing matching against the given city/country
// name combination
// If an exact match by city name is found, it returns immediately.
// If no match is found by exact city name, it will try to match based on fuzzy search
// against country name and city/country combination
// eg: "London", "London UK", "London United Kingdom" should all return the same result.
func filterByTZCode(rows []CityData, s string) ([]string, error) {
	return []string{""}, shared.ErrorNotImplementedError{}
}
