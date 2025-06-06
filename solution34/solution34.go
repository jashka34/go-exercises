package solution34

func sum(nums []int, ch chan<- int) {
	var r int
	for i := range nums {
		// fmt.Println(nums[i])
		r += nums[i]
	}
	ch <- r
}

func MaxSum(nums1, nums2 []int) []int {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go sum(nums1, ch1)
	go sum(nums2, ch2)
	r1 := <-ch1
	r2 := <-ch2
	// fmt.Println("r1:", r1)
	// fmt.Println("r2:", r2)
	if r1 >= r2 {
		return nums1
	} else {
		return nums2
	}
}
