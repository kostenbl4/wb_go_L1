package main

import "sync"

func takeNums(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, num := range nums {
			out <- num
		}
	}()

	return out
}

func giveSquares(in <-chan int) <-chan int {
	out := make(chan int)

	const nWorkers = 3

	var wg sync.WaitGroup
	wg.Add(nWorkers)

	for i := 0; i < nWorkers; i++ {
		go func() {
			defer wg.Done()
			for num := range in {
				out <- num * num
			}
		}()
	}

	go func() {
		// ждем завершения всех горутин
		wg.Wait()
		close(out)
	}()

	return out
}

func sumSquares(in <-chan int) int {
	sum := 0
	for num := range in {
		sum += num
	}
	return sum
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	out := takeNums(nums)
	squares := giveSquares(out)
	sum := sumSquares(squares)
	println(sum)
}
