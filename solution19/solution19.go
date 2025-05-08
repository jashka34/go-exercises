package solution19

import "fmt"

func UniqueUserIDs(userIDs []int64) []int64 {
	uniqs := make(map[int64]int, len(userIDs))
	ret := make([]int64, 0)

	for _, v := range userIDs {
		// fmt.Println(i, ":", id, "uniq:", uniq)
		_, exists := uniqs[v]
		if !exists {
			uniqs[v] = 0
			ret = append(ret, v)
		}
	}

	fmt.Println("--------")
	fmt.Println(userIDs)
	fmt.Println(uniqs, "len:", len(uniqs))
	fmt.Println(ret)
	fmt.Println("--------")
	return ret
}
