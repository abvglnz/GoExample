package stringutils

import "unicode/utf8"

// Length returns the number of characters in s.
func Length(s string) int {
	return utf8.RuneCountInString(s)
}
