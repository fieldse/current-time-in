// Tests for the search string parser functions
package citylookup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

// func Test_parseSearchString(t *testing.T) {

// }
