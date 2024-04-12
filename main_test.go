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
		{"New York", ""},      // test split word city
		{"New York city", ""}, // test split word city plus "City" suffix
		{"philadephia", ""},   // test lower case
		{"Washington", ""},
		{"Washington, D.C.", ""}, // Test city with abbreviation suffix
		{"San Francisco", ""},    // test split word city
		{"Dallas TX", ""},        // Test with state abbreviation
		{"Dallas, Texas", ""},
		{"brisbane", ""},
		{"Chicago", "America/Chicago"},
		{"paris", ""},
		{"Frankfurt", ""},
		{"Tokyo", ""},
		{"Paris, France", ""},  // test city/country with comma
		{"Birmingham", ""},     // test city with ambiguous country -- this could be US or UK
		{"Birmingham UK", ""},  // test city with ambiguous country -- should be UK
		{"Birmingham USA", ""}, // test city with ambiguous country -- should be US
		{"London", ""},
		{"London England", ""}, // test city/country with no comma
		{"London UK", ""},      // test city/country with alternative country name
		{"Cairo", ""},
		{"Cairo, Egypt", ""},
		{"Beijing", ""},
	}

	for _, tc := range testCases {
		assert.Equalf(t, tc.expected, tc.name, "expected %s to be %s", tc.expected, tc.name)
	}
}
