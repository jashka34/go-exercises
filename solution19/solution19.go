package solution19

import "fmt"

func UniqueUserIDs(userIDs []int64) []int64 {
	if len(userIDs) <= 1 {
		return userIDs
	}
	// sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })

	retm := make(map[int64]int, len(userIDs))
	ret := make([]int64, len(userIDs))

	for i, v := range userIDs {
		// fmt.Println(i, ":", id, "uniq:", uniq)
		retm[v] = i - 1
	}

	for i, v := range retm {
		fmt.Println("i:", i, "v:", v)
		ret[v] = i
	}

	fmt.Println("--------")
	fmt.Println(userIDs)
	fmt.Println(retm, "len:", len(retm))
	fmt.Println(ret)
	fmt.Println("--------")
	return ret[:len(retm)]
}
