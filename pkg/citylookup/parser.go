// parsing functions for search strings
package citylookup

import (
	"fmt"
	"regexp"
	"strings"
)

// Regex for nonalphabetical: /[^a-zA-Z\d\s:\u00C0-\u00FF]/g
// Matches all characters outside the Latin alphabet range, A-Z
// plus latin characters with diacritics, like "ñ" and "é"
var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z\d\s:\u00C0-\u00FF]+`)

// replaceNonAlpha removes all non-alphabetic characters from a string
func replaceNonAlpha(s string) string {
	return nonAlphanumericRegex.ReplaceAllString(s, "")
}

// parseSearchIndexRows fills in the searchable index field for each city row
// with combined values from city, country, province, and iso2,3 codes.
// Used for matching against fuzzy search
func parseSearchIndexRows(rows []CityRow) ([]CityRow, error) {
	newRows := make([]CityRow, len(rows))
	for _, r := range rows {
		// generate search string
		searchString, err := parseSearchableString(r)
		if err != nil {
			return []CityRow{}, fmt.Errorf("failed to parse search string from row %+v: %w", r, err)
		}
		// Store the search string and copy the row
		r2 := r
		r2.SearchString = searchString
		newRows = append(newRows, r2)
	}
	return newRows, nil
}

// parseSearchableString combine values from city, country, province, and iso2+3 codes
// to make a searchable string for indexing
func parseSearchableString(row CityRow) (string, error) {
	var words []string
	var cleaned []string
	// The Iso2 field doesn't parse correctly in all rows, so we need to coerce it to string
	words = append(
		words,
		row.City,
		row.CityAscii,
		row.Country,
		row.Province,
		row.ExactCity,
		row.ExactProvince,
		row.Iso3,
		fmt.Sprintf("%v", row.Iso2),
	)
	// clean non-alpha characters
	for _, w := range words {
		cleaned = append(cleaned, strings.ToLower(replaceNonAlpha(w)))
	}
	return strings.Join(cleaned, " "), nil
}
