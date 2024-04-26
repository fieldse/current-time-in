package citylookup

import (
	"fmt"
	"log"
	"testing"

	"github.com/fieldse/current-time-in/shared"
	"github.com/stretchr/testify/suite"
)

type CityLookupTests struct {
	suite.Suite
	cityData []CityRow
}

// Run all tests
func TestCityLookup(t *testing.T) {
	log.Println("TestCityLookup()")
	suite.Run(t, new(CityLookupTests))
}

// Test the city returns a correct city/country code by name
func (t *CityLookupTests) TestCityCode() {

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
		res, err := filterFuzzyCityCountryName(t.cityData, tc.name)
		t.Assert().Nilf(err, "should not return error")
		t.Assert().Equalf(tc.numExpected, len(res), "expected %d results, got %d", tc.numExpected, len(res))
	}
}

// Test loading the city data from cityMap.json
func (t *CityLookupTests) Test_loadCityData() {
	data, err := loadCityData()
	t.Assert().Nilf(err, "should load city data without error")
	t.Assert().NotEmpty(data, "data should not be empty")
	Logger.Debug().Msgf("city data: %v", shared.Truncate(fmt.Sprintf("%v", data), 1000))
}

// Test reading the raw city data from file
func (t *CityLookupTests) Test_readCityData() {
	data, err := readCityData()
	t.Assert().Nilf(err, "read data file failed: %v", err)
	t.Assert().NotEmptyf(data, "data should not be empty")
}

func (t *CityLookupTests) Test_findCityExact() {
	res, err := findCityExact(t.cityData, "New York")
	t.Assert().Nil(err)
	t.Assert().Equalf("New York", res.City, "name should match")
	Logger.Debug().Msgf("city result: %v", res.City)
}

func (t *CityLookupTests) Test_filterByCountry() {
	var cityNames []string
	var expectedCities = []string{"Denver", "Houston", "Miami", "Atlanta", "Chicago",
		"Los Angeles", "Washington, D.C.", "New York"}

	// Filter by united states
	res := filterByCountry(t.cityData, "United States")
	t.Assert().NotEmptyf(res, "results should not be empty")

	// Aggregate results
	for _, city := range res {
		cityNames = append(cityNames, city.City)
	}

	// Confirm expected cities in results
	for _, c := range expectedCities {
		t.Assert().Containsf(cityNames, c, "filtered list should contain %s", c)
	}

	// Filter by nonexistent
	res = filterByCountry(t.cityData, "foobar")
	t.Assert().Emptyf(res, "nonexistent should be empty")

	// Filter by lowercase and partial match
	res = filterByCountry(t.cityData, "united")
	t.Assert().NotEmptyf(res, "results should not be empty")

	cityNames = []string{}
	expectedCities = []string{"New York", "London"}

	// Aggregate results
	for _, city := range res {
		cityNames = append(cityNames, city.City)
	}

	// Confirm expected cities in results
	for _, c := range expectedCities {
		t.Assert().Containsf(cityNames, c, "filtered list should contain %s", c)
	}
}

// Load the city data once for test suite
func (t *CityLookupTests) SetupSuite() {
	data, err := loadCityData()
	if err != nil {
		panic("load testing data failed: %s" + err.Error())
	}
	t.cityData = data
}
