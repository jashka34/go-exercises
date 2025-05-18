package solution26

func MergeNumberLists(numberLists ...[]int) []int {
	var ret []int
	ret = []int{}
	for _, nl := range numberLists {
		// fmt.Println(nl)
		ret = append(ret, nl...)
	}
	return ret
}
