package solution34

import "fmt"

func sum(nums []int, ch chan<- int) {
	var r int
	for i := range nums {
		fmt.Println(nums[i])
		r += nums[i]
	}
	ch <- r
}

func MaxSum(nums1, nums2 []int) []int {
	var ret []int
	ch1 := make(chan int)
	go sum(nums1, ch1)
	r1 := <-ch1
	fmt.Println("r1:", r1)
	return ret
}
