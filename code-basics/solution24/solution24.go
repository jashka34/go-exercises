package solution24

import (
	"strings"
	"unicode"
)

func latinLetters(s string) string {
	ret := &strings.Builder{}
	for _, w := range s {
		if unicode.Is(unicode.Latin, w) {
			ret.WriteRune(w)
		}
	}
	return ret.String()
}
