// Miscellaneous small helpers
package citylookup

import "os"

// Exists checks if a file or directory exists
func Exists(fp string) bool {
	_, err := os.Stat(fp)
	return !os.IsNotExist(err)
}

// Truncate truncates a string to given maximum length
func Truncate(s string, chars int) string {
	if chars >= len(s) {
		return s
	}
	return s[:chars] + "..."
}
