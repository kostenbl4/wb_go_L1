package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "snow dog sun"

	// Разбиваем строку на слова
	words := strings.Fields(s)

	// Создаем массив для перевернутых слов
	reversed := make([]string, len(words))

	// Переворачиваем слова
	for i := 0; i < len(words); i++ {
		reversed[i] = words[len(words)-1-i]
	}

	// Выводим исходную строку и перевернутую
	fmt.Print(s, " — ")
	fmt.Println(strings.Join(reversed, " "))

	// Вариант без использования дополнительной памяти
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	fmt.Print(s, " — ")
	fmt.Println(strings.Join(words, " "))
}
