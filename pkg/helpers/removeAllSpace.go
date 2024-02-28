package helpers

import "strings"

// RemoveAllSpace - remove all spaces and return
// the result as string
func RemoveAllSpace(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
