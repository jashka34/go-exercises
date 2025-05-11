package solution20

import "fmt"

func MostPopularWord(words []string) string {
	// var ret string
	wm := make(map[int]string, len(words))
	max := 0
	i := 0
	for _, w := range words {
		if w > max {
			max = wm[i]
			fmt.Println("w:", w, "max:", max)
		}
	}
	var mword string
	// for w, val := range wm {
	// 	if val > max {
	// 		max = val
	// 		ret = w
	// 	}
	// }
	return mword
}
