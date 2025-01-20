package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1 // Элемент не найден
}

func main() {
	// Массив должен быть отсортирован
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Массив: ", nums)
	target := 5
	var result = binarySearch(nums, target)
	fmt.Printf("Индекс %v = %v\n", target, result)

	target = 8
	result = binarySearch(nums, target)
	fmt.Printf("Индекс %v = %v\n", target, result)

	// Встроенный поиск (снизу комментарий из исходников)
	// Search uses binary search to find and return the smallest index i
	// in [0, n) at which f(i) is true, assuming that on the range [0, n),
	target = 5
	index := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	fmt.Printf("Встроенный поиск. Индекс %v = %v\n", target, index)

	target = 8
	index = sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	fmt.Printf("Встроенный поиск. Индекс %v = %v\n", target, index)
}
