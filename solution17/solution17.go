package solution17

func IntsCopy(src []int, maxLen int) []int {
	if maxLen <= 0 {
		return []int{}
	}
	l := 0
	if maxLen > len(src) {
		l = len(src)
	} else {
		l = maxLen
	}
	res := make([]int, l)
	copy(res, src)
	return res
}
