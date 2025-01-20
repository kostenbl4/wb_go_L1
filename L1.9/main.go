package main

func readNumbers(numbers []int) <-chan int {
	out := make(chan int) // Канал, куда кладутся числа из списка

	// Горутина, которая кладет числа в канал
	go func() {
		defer close(out)
		for _, v := range numbers {
			out <- v
		}
	}()

	return out
}

func writeSquares(in <-chan int) <-chan int {
	out := make(chan int) // Канал, куда кладутся квадраты чисел

	// Горутина, которая кладет квадраты чисел в канал
	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()

	return out
}

func printSquares(in <-chan int) {
	// Горутина, которая печатает числа из канала
	for v := range in {
		println(v)
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	printSquares(writeSquares(readNumbers(numbers)))
}
