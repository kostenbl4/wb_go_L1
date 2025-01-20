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

func giveSquares(in <-chan int) <-chan int {
	out := make(chan int)

	go func(){
		defer close(out)
		for num := range in {
			out <- num * num
		}
	}()
	
	return out
}

func sumSquares(in <-chan int) int{
	sum := 0
	for num := range in {
		sum += num
	}
	return sum
}

func main(){
	nums := []int{2, 4, 6, 8, 10}
	out := takeNums(nums)
	squares := giveSquares(out)
	sum := sumSquares(squares)
	println(sum)
}