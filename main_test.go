package main

import (
	"fmt"

	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the city returns a correct city/country code by name
func TestCityCode(t *testing.T) {

	testCases := []struct {
		name        string
		expected    string
		numExpected int
	}{
		{"New York", "America/New_York", 1},         //  split word city
		{"New York city", "America/New_York", 1},    //  split word city plus "City" suffix
		{"philadephia", "America/New_York", 1},      //  lower case
		{"Washington", "America/New_York", 1},       //  city with ambiguous state equivalent
		{"Washington, D.C.", "America/New_York", 1}, //  city with abbreviation suffix
		{"San Francisco", "America/Los_Angeles", 1}, //  split word city
		{"Dallas TX", "America/Chicago", 1},         //  city with state abbreviation -- note this one may be multiple TZ codes
		{"Dallas, Texas", "America/Chicago", 1},     //  city with state name -- note this one may be multiple TZ codes
		{"Chicago", "America/Chicago", 1},
		{"Brisbane", "Australia/Brisbane", 1},    // international city with no country name
		{"Paris", "Europe/Paris", 1},             // international city with no country name
		{"Paris, France", "Europe/Paris", 1},     // internationa city/country with comma
		{"Frankfurt", "Europe/Berlin", 1},        // international city with no country name
		{"Tokyo", "Asia/Tokyo", 1},               // international city with no country name
		{"Birmingham", "", 2},                    // international city with ambiguous country -- this could be US or UK
		{"Birmingham UK", "Europe/London", 1},    // international city with ambiguous country, with country code -- should be UK
		{"Birmingham USA", "America/Chicago", 1}, // international city with ambiguous country , with country code -- should be US CST
		{"London", "Europe/London", 1},           // international city with no country name
		{"London England", "Europe/London", 1},   // international city/country with no comma
		{"London UK", "Europe/London", 1},        // international city/country with alternative country name
		{"Cairo", "Africa/Cairo", 1},             // international city with multiple countries, but an obvious default
		{"Cairo, Egypt", "Africa/Cairo", 1},
		{"Beijing", "Asia/Shanghai", 1},
	}

	for _, tc := range testCases {
		res, err := tzCodeFor(tc.name)
		assert.Nilf(t, err, "should not return error")
		assert.Equalf(t, tc.expected, res, "expected %s to be %s", tc.expected, res)
		assert.Equalf(t, tc.numExpected, len(res), "expected %d results, got %d", tc.numExpected, len(res))
	}
}

// Test loading the city data from cityMap.json
func Test_loadCityData(t *testing.T) {
	data, err := loadCityData()
	assert.Nilf(t, err, "should load city data without error")
	assert.NotEmpty(t, data, "data should not be empty")
	AppLogger.Debug("city data: ", data)
}

// Test reading the raw city data from file
func Test_readCityData(t *testing.T) {
	data, err := readCityData()
	assert.Nilf(t, err, "read data file failed: %v", err)
	assert.NotEmptyf(t, data, "data should not be empty")
}

func Test_findCityExact(t *testing.T) {
	data, err := loadCityData()
	if err != nil {
		panic(err)
	}
	res, err := findCityExact(data, "New York")
	assert.Nil(t, err)
	assert.Equalf(t, "New York", res.City, "name should match")
	AppLogger.Debug(fmt.Sprintf("city result: %v", res.City))
}

func Test_testLog(t *testing.T) {
	AppLogger.Info("testing logging")
}
