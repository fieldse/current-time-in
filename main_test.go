package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the city returns a correct city/country code by name
func TestCityCode(t *testing.T) {

	testCases := []struct {
		name     string
		expected string
	}{
		{"New York", "America/New_York"},         //  split word city
		{"New York city", "America/New_York"},    //  split word city plus "City" suffix
		{"philadephia", "America/New_York"},      //  lower case
		{"Washington", "America/New_York"},       //  city with ambiguous state equivalent
		{"Washington, D.C.", "America/New_York"}, //  city with abbreviation suffix
		{"San Francisco", "America/Los_Angeles"}, //  split word city
		{"Dallas TX", "America/Chicago"},         //  city with state abbreviation -- note this one may be multiple TZ codes
		{"Dallas, Texas", "America/Chicago"},     //  city with state name -- note this one may be multiple TZ codes
		{"Chicago", "America/Chicago"},
		{"Brisbane", "Australia/Brisbane"},    // international city with no country name
		{"Paris", "Europe/Paris"},             // international city with no country name
		{"Paris, France", "Europe/Paris"},     // internationa city/country with comma
		{"Frankfurt", "Europe/Berlin"},        // international city with no country name
		{"Tokyo", "Asia/Tokyo"},               // international city with no country name
		{"Birmingham", ""},                    // international city with ambiguous country -- this could be US or UK
		{"Birmingham UK", "Europe/London"},    // international city with ambiguous country -- should be UK
		{"Birmingham USA", "America/Chicago"}, // international city with ambiguous country -- should be US CST
		{"London", "Europe/London"},           // international city with no country name
		{"London England", "Europe/London"},   // international city/country with no comma
		{"London UK", "Europe/London"},        // international city/country with alternative country name
		{"Cairo", "Africa/Cairo"},             // international city with multiple countries, but an obvious default
		{"Cairo, Egypt", "Africa/Cairo"},
		{"Beijing", "Asia/Shanghai"},
	}

	for _, tc := range testCases {
		res, err := tzCodeFor(tc.name)
		assert.Nilf(t, err, "should not return error")
		assert.Equalf(t, tc.expected, res, "expected %s to be %s", tc.expected, res)
	}
}
