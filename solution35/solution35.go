package solution35

func sum1(nums []int) int {
	var sum int
	for _, i := range nums {
		// fmt.Println(i, sum)
		sum += i
	}
	return sum
}

func SumWorker(numsCh chan []int, sumCh chan int) {
	// fmt.Println("start")
	for nums := range numsCh {
		// fmt.Println("nums:", nums)
		sumCh <- sum1(nums)
		// fmt.Println("sum:", sum)
	}
}
