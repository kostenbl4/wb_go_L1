package main

import (
	"fmt"
	"sync"
)

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

func square(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * num
	}
}

func printSquares(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	in := takeNums(nums)
	out := make(chan int)

	// Запускаем несколько горутин для параллельных вычислений
	const numWorkers = 3
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			square(in, out)
		}()
	}

	// Горутина для закрытия выходного канала после завершения всех вычислений
	go func() {
		wg.Wait()
		close(out)
	}()

	printSquares(out)
}
