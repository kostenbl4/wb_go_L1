package main

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

func printSquares(in <-chan int) {
	for num := range in {
		println(num * num)
	}
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	out := takeNums(nums)
	printSquares(out)
}
