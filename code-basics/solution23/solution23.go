package solution23

import (
	"unicode"
)

func isASCII(s string) bool {
	// fmt.Println(s, unicode.MaxASCII)
	for _, ch := range s {
		// fmt.Println(s[i], ch)
		if ch > unicode.MaxASCII {
			// fmt.Println("false")
			return false
		}
	}
	return true
}
