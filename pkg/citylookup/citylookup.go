// Get the current time in a given city
package citylookup

import (
	"strings"

	"github.com/fieldse/current-time-in/pkg/logger"

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

// filterFuzzyCityCountryName filters the cities data by fuzzing matching against the given
// city/country name combination
// If an exact match by city name is found, it returns immediately.
// If no match is found by exact city name, it will try to match based on fuzzy search
// against country name and city/country combination
// eg: "London", "London UK", "London United Kingdom" should all return the same result.
func filterFuzzyCityCountryName(rows []CityData, s string) ([]CityData, error) {
	var filtered []CityData

	// Exact match by name
	exactMatch, err := findCityExact(rows, s)
	if err == nil { // if we have exact match, return immediately
		filtered = append(filtered, exactMatch)
		return filtered, nil
	}

	// If no name match, see if we can split out a country name
	logger.Logger.Info().Msgf("no exact match found for city name. Attempting fuzzy match")

	words := strings.Split(s, "")
	numWords := len(words)
	if numWords > 1 { // we have more than one word
		for i := 0; i < numWords; i++ {
			index := numWords - i
			lastItem := words[index]
			logger.Logger.Info().Msgf("attempting match for country name on %s", lastItem)

			byCountry := filterByCountry(rows, lastItem)
			if len(byCountry) > 0 {
				logger.Logger.Info().Msgf("got %d matches by country name", len(byCountry))
				return byCountry, nil

			}
		}
	}

	return []CityData{}, fmt.Errorf("no match found for %s", s)
}
