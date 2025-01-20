package main

import "fmt"

func main() {
	var a, b int = 1, 2
	fmt.Printf("Переменные до перестановки a=%v, b=%v\n", a, b)
	a, b = b, a
	fmt.Printf("Переменные после перестановки a=%v, b=%v\n", a, b)
}
