package solution20

func MostPopularWord(words []string) string {
	var ret string
	wm := make(map[string]int)
	for _, w := range words {
		wm[w] = wm[w] + 1
	}
	max := 0
	for w, val := range wm {
		if val > max {
			max = val
			ret = w
		}
	}
	return ret
}
