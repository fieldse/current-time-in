// Miscellaneous small helpers
package citylookup

// Truncate truncates a string to given maximum length
func Truncate(s string, chars int) string {
	if chars >= len(s) {
		return s
	}
	return s[:chars] + "..."
}
