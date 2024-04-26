// Tests for the search string parser functions
package citylookup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	cityRow       CityRow
	shouldContain []string
}

// Generate some test case cities with different string types, including punctuation
func testCities() []testCase {
	rows := []testCase{

		{
			cityRow: CityRow{
				City:      "Los Angeles",
				CityAscii: "Los Angeles",
				Country:   "United States of America",
				Iso2:      "US",
				Iso3:      "USA",
				Province:  "California",
				StateAnsi: "CA",
				Timezone:  "America/Los_Angeles",
			},
			shouldContain: []string{"los angeles", "united states", "california", "usa"},
		},
		{
			cityRow: CityRow{
				City:      "Washington, D.C.",
				CityAscii: "Washington, D.C.",
				Country:   "United States of America",
				Iso2:      "US",
				Iso3:      "USA",
				Province:  "District of Columbia",
				Timezone:  "America/New_York",
			},
			shouldContain: []string{"washington dc", "united states", "usa", "district of columbia"},
		},
		{
			cityRow: CityRow{
				City:      "Christiansted",
				CityAscii: "Christiansted",
				Country:   "United States Virgin Islands",
				Iso2:      "VI",
				Iso3:      "VIR",
				Province:  "Virgin Islands",
				StateAnsi: "VI",
				Timezone:  "America/St_Thomas",
			},
			shouldContain: []string{"christiansted", "united states virgin islands", "vi", "vir", "virgin islands", "vi"},
		},
		{
			cityRow: CityRow{
				City:      "Bac Lieu",
				CityAscii: "Bac Lieu",
				Country:   "Vietnam",
				Iso2:      "VN",
				Iso3:      "VNM",
				Province:  "Bạc Liêu",
				Timezone:  "Asia/Ho_Chi_Minh",
			},
			shouldContain: []string{"bac lieu", "vietnam", "bạc liêu", "vn", "vnm"},
		},
		{
			cityRow: CityRow{
				City:      "Sao Paulo",
				CityAscii: "Sao Paulo",
				Country:   "Brazil",
				Iso2:      "BR",
				Iso3:      "BRA",
				Province:  "São Paulo",
				Timezone:  "America/Sao_Paulo",
			},
			shouldContain: []string{"sao paulo", "brazil", "br", "bra", "são paulo"},
		},
		{
			cityRow: CityRow{
				City:      "Fray Bentos",
				CityAscii: "Fray Bentos",
				Country:   "Uruguay",
				Iso2:      "UY",
				Iso3:      "URY",
				Province:  "Río Negro",
				Timezone:  "America/Montevideo",
			},
			shouldContain: []string{"fray bentos", "uruguay", "uy", "ury", "río negro"},
		},
		{
			cityRow: CityRow{
				City:      "Qal eh-ye Now",
				CityAscii: "Qal eh-ye",
				Country:   "Afghanistan",
				Iso2:      "AF",
				Iso3:      "AFG",
				Province:  "Badghis",
				Timezone:  "Asia/Kabul",
			},
			shouldContain: []string{"qal ehye now", "qal ehye", "afghanistan", "af", "afg", "badghis"},
		},
	}
	return rows
}

func Test_replaceNonAlpha(t *testing.T) {
	cases := []struct {
		s      string
		expect string
	}{
		{"abcdefg", "abcdefg"},                // no changes
		{"AbcDEFG", "AbcDEFG"},                // respect casing
		{"abc de f g", "abc de f g"},          // respect spacing (no change)
		{"abcdef123", "abcdef"},               // replace numeric
		{"@#$%^*()`,./abc-defg!", "abcdefg"},  // replace punctuation
		{"123456", ""},                        // completely replace nonalpha string
		{"Washington, D.C.", "Washington DC"}, // respect casing
	}
	for _, x := range cases {
		res := replaceNonAlpha(x.s)
		assert.Equal(t, x.expect, res)
	}
}

func Test_parseSearchableString(t *testing.T) {
	cases := testCities()
	for _, c := range cases {
		res, err := parseSearchableString(c.cityRow)
		assert.Nil(t, err)
		for _, v := range c.shouldContain {
			assert.Containsf(t, res, v, "result should contain %s", v)
		}
	}
}
