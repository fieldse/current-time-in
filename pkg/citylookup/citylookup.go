// Get the current time in a given city
package citylookup

import (
	"strings"

	"github.com/fieldse/current-time-in/pkg/logger"

	"fmt"
)

var Logger = logger.Logger

// CityRow represents a single entry from the cities data table
type CityRow struct {
	City      string  `json:"city"`
	CityAscii string  `json:"city_ascii"`
	Lat       float32 `json:"lat"`
	Lng       float32 `json:"lng"`
	Pop       float32 `json:"pop"`
	Country   string  `json:"country"`
	Iso2      any     `json:"iso2"` // these sometimes will be "-99" for no obvious reason
	Iso3      string  `json:"iso3"`
	Province  string  `json:"province"`
	StateAnsi string  `json:"state_ansi"`
	Timezone  string  `json:"timezone"`

	// unclear how these differ from City/Province, but included in some rows
	ExactCity     string `json:"exactCity"`
	ExactProvince string `json:"exactProvince"`

	SearchString string // this wil be indexed as city/country/province/iso2,3 combination for fuzzy search
}

// findCityExact finds a single city matching exactly by name
func findCityExact(rows []CityRow, s string) (CityRow, error) {
	for _, r := range rows {
		if r.City == s {
			return r, nil
		}
	}
	return CityRow{}, fmt.Errorf("city not found: %s", s)
}

// filterByCountry filters cities by country name on a case-insensitive substring match
// example: "united" would return "United States" and "United Kingdom"
func filterByCountry(rows []CityRow, countryName string) []CityRow {
	var filtered []CityRow
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
func filterFuzzyCityCountryName(rows []CityRow, s string) ([]CityRow, error) {
	var filtered []CityRow

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

	return []CityRow{}, fmt.Errorf("no match found for %s", s)
}
