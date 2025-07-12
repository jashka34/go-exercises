package solution08

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Greetings(name string) string {
	retName := name
	retName = strings.Trim(retName, " ")
	retName = strings.ToLower(retName)
	retName = cases.Title(language.Russian).String(retName)
	return fmt.Sprintf("Привет, %s!", retName)
}
