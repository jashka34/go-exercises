package solution35

func SumWorker(numsCh chan []int, sumCh chan int) {
	// fmt.Println("start")
	var sum int
	for nums := range numsCh {
		for _, i := range nums {
			// fmt.Println(i, sum)
			sum += i
		}
	}
	sumCh <- sum
}
