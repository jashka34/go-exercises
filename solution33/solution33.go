package solution33

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxSum(nums1, nums2 []int) []int {
	// ret := []int{}
	l := max(len(nums1), len(nums2))
	// fmt.Println("l:", l)
	var s1, s2, ls1, ls2 int
	ls1 = len(nums1)
	ls2 = len(nums2)
	// fmt.Println("ls1:", ls1, "ls2:", ls2)
	for i := 0; i < l; i++ {
		// fmt.Println("i:", i)
		if i < ls1 {
			s1 += nums1[i]
			// fmt.Println("nums1:", nums1[i])
		}
		if i < ls2 {
			s2 += nums2[i]
			// fmt.Println("nums2:", nums2[i])
		}
	}
	// fmt.Println("sum:", s1, s2)
	if s1 >= s2 {
		return nums1
	} else {
		return nums2
	}
}
