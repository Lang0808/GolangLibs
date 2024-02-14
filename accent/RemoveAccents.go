package accent

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Remove Vietnamese Accent in string
func RemoveAccent(s string) string {
	s = strings.ReplaceAll(s, "Đ", "D") // Don't know why Đ|đ cannot be transformed to D|d so we need to use this
	s = strings.ReplaceAll(s, "đ", "d")
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	res, _, _ := transform.String(t, s)
	return res
}
