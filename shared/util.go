// Miscellaneous small helpers
package shared

import (
	"fmt"
	"os"
)

// Exists checks if a file or directory exists
func Exists(fp string) bool {
	_, err := os.Stat(fp)
	return !os.IsNotExist(err)
}

// ExistsOrPanic checks if a file or directory exists; panics on error
func ExistsOrPanic(fp string) {
	if !Exists(fp) {
		panic(fmt.Sprintf("path does not exist: %s", fp))
	}
}

// Truncate truncates a string to given maximum length
func Truncate(s string, chars int) string {
	if chars >= len(s) {
		return s
	}
	return s[:chars] + "..."
}
