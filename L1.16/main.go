package main

import (
	"fmt"
	"slices"
)

func quicksort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	// Опорный элемент - первый элемент массива
	pivot := arr[0]
	left := []int{}
	right := []int{}

	// Разбиваем массив на две части
	for _, v := range arr[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	// Рекурсивно вызываем функцию для левой и правой части массива
	left = quicksort(left)
	right = quicksort(right)

	// Склеиваем массивы
	return append(append(left, pivot), right...)
}

func main() {
	fmt.Println("Встроенная сортировка")
	var nums = []int{4, 2, 5, 7, 1, 3, 6}
	fmt.Println("До сортировки: ", nums)
	// Встроенная сортировка основана на quicksort (снизу исходники)
	//func Sort[S ~[]E, E cmp.Ordered](x S) {
	// 	n := len(x)
	// 	pdqsortOrdered(x, 0, n, bits.Len(uint(n)))
	// }
	// pdqsortOrdered sorts data[a:b].
	// The algorithm based on pattern-defeating quicksort(pdqsort), but without the optimizations from BlockQuicksort.
	slices.Sort(nums)
	fmt.Println("После сортировки: ", nums)

	nums = []int{4, 2, 5, 7, 1, 3, 6}
	fmt.Println("Рукописная сортировка")
	fmt.Println("До сортировки: ", nums)
	nums = quicksort(nums)
	fmt.Println("После сортировки: ", nums)
}
