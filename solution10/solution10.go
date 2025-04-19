package solution10

import "strings"

func ModifySpaces(s, mode string) string {
	var r string
	switch mode {
	case "dash":
		r = "-"
	case "underscore":
		r = "_"
	default:
		r = "*"
	}
	return strings.ReplaceAll(s, " ", r)
}
