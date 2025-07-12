package solution14

func SafeWrite(nmbs [5]int, i, val int) [5]int {
	if i >= 0 && i <= 4 {
		nmbs[i] = val
	}
	return nmbs
}
