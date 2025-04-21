package solution12

const (
	ok = iota
	cancelled
	unknown
)

func ErrorMessageToCode(msg string) int {
	if msg == "OK" {
		return ok
	} else if msg == "CANCELLED" {
		return cancelled
	} else {
		return unknown
	}
}
