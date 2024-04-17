// Miscellaneous small helpers
package main

// Truncate truncates a string to given maximum length
func Truncate(s string, chars int) string {
	if chars >= len(s) {
		return s
	}
	return s[:chars] + "..."
}
