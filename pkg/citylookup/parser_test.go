// Tests for the search string parser functions
package citylookup

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	cityRow       CityRow
	shouldContain []string
}

// Generate some test case cities with different string types, including punctuation, and
// expected search index string
func testCitiesWithExpect() []testCase {
	rows := []testCase{
		{
			cityRow:       Example_LosAngeles,
			shouldContain: []string{"los angeles", "united states", "california", "usa"},
		},
		{
			cityRow:       Example_Washington,
			shouldContain: []string{"washington dc", "united states", "usa", "district of columbia"},
		},
		{
			cityRow:       Example_Christiansted,
			shouldContain: []string{"christiansted", "united states virgin islands", "vi", "vir", "virgin islands", "vi"},
		},
		{
			cityRow:       Example_BacLieu,
			shouldContain: []string{"bac lieu", "vietnam", "bạc liêu", "vn", "vnm"},
		},
		{
			cityRow:       Example_SaoPaulo,
			shouldContain: []string{"sao paulo", "brazil", "br", "bra", "são paulo"},
		},
		{
			cityRow:       Example_FrayBentos,
			shouldContain: []string{"fray bentos", "uruguay", "uy", "ury", "río negro"},
		},
		{
			cityRow:       Example_QalEhYe,
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
	cases := testCitiesWithExpect()
	for _, c := range cases {
		res, err := parseSearchableString(c.cityRow)
		assert.Nil(t, err)
		for _, v := range c.shouldContain {
			assert.Containsf(t, res, v, "result should contain %s", v)
		}
	}
}

func Test_parseRows(t *testing.T) {
	rows := ExampleCities
	res, err := parseSearchIndexRows(rows)
	assert.Nil(t, err)
	for _, r := range res {
		assert.NotEmptyf(t, r.SearchString, "result should have a search string")
		log.Printf("=== debug: generated searchString: \"%s\" for city: %s", r.SearchString, r.City)
	}

}
