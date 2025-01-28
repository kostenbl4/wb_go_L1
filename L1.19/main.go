package main

import "fmt"

func reverse(s string) string {
	// Преобразуем строку в массив рун, чтобы работать с символами Unicode
	runes := []rune(s)
	res := make([]rune, len(runes))

	// Записываем символы в обратном порядке
	for i := 0; i < len(runes); i++ {
		res[i] = runes[len(runes)-1-i]
	}
	return string(res)
}

// вариант без использования дополнительной памяти
func reverse2(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	// Примеры использования функции reverse
	word := "Hello World!"
	reversed := reverse(word)
	fmt.Println(reversed)

	word = "Привет, Мир!"
	reversed = reverse(word)
	fmt.Println(reversed)

	// Примеры использования функции reverse2
	word = "Hello World!"
	reversed = reverse2(word)
	fmt.Println(reversed)

}
