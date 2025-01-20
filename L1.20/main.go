package main

import (
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
	print(s, " — ")
	println(strings.Join(reversed, " "))
}
