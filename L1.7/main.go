package main

import "fmt"

func changeIthBit(n, val int64, i int) int64 {
	if val == 1 {
		return n | (1 << i)
	} else if val == 0 {
		return n & ^(1 << i)
	}
	return -1
}

func main() {
	var num int64
	fmt.Print("Input decimal n: ")
	fmt.Scan(&num)

	fmt.Printf("Binary n before: %b\n", num)

	var val int64
	fmt.Print("Input val: ")
	fmt.Scan(&val)

	var i int
	fmt.Print("Input i: ")
	fmt.Scan(&i)

	binNum := changeIthBit(num, val, i)
	fmt.Printf("Binary n after: %b\n", binNum)
	fmt.Printf("Decimal n after: %d\n", binNum)
}
