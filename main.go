// Get the current time in a given city
package main

import "errors"

// tzCodeFor returns a full timezone name for the given city/country name.
// This function fuzzy-matches based on the city name, and attempts to return intelligently
// from cities matching the name.
// eg: "London", "London UK", "London United Kingdom" should all return the same result.
// This will return a slice of results for cities which exist in multiple countries.
func tzCodeFor(s string) ([]string, error) {
	return []string{""}, errors.New("not implemented")
}
