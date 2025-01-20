package main

import "fmt"

func removeElement(slice []int, i int) []int {
	// Проверка на выход за границы массива
	if i < 0 || i >= len(slice) {
		return slice
	}

	// Удаление элемента из массива
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	// Создание массива
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// Удаление элемента из массива
	nums = removeElement(nums, 2)
	fmt.Println(nums)
}
