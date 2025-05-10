package solution20

import "fmt"

func MostPopularWord(words []string) string {
	// var ret string
	wm := make(map[string]int)
	max := 0
	var mword string
	for _, w := range words {
		wm[w] = wm[w] + 1
		if wm[w] > max {
			max = wm[w]
			mword = w
			fmt.Println("w:", w, "max:", max, "mword:", mword)
		}
	}
	// for w, val := range wm {
	// 	if val > max {
	// 		max = val
	// 		ret = w
	// 	}
	// }
	return mword
}
