package solution20

func MostPopularWord(words []string) string {
	wm := make(map[string]int, len(words))
	max := 0
	for _, w := range words {
		wm[w]++
		if wm[w] > max {
			max = wm[w]
			// fmt.Println("w:", w, "max:", max)
		}
	}
	var mword string
	for _, w := range words {
		if wm[w] == max {
			mword = w
			break
		}
	}
	return mword
}
