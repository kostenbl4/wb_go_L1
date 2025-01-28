package main

import (
	"unicode"
)

func isAllUnique(s string) bool {
	// Создание множества для хранения уникальных символов
	set := make(map[rune]struct{})

	// Перебор всех символов строки
	for _, r := range s {
		lowerR := unicode.ToLower(r) // Приведение символа к нижнему регистру
		if _, ok := set[lowerR]; ok {
			return false
		}
		set[lowerR] = struct{}{}
	}

	return true
}

func main() {
	println(isAllUnique("abcd"))      // true
	println(isAllUnique("abCdefAaf")) // false
	println(isAllUnique("aabcd"))     // false
	println(isAllUnique("Aabcd"))     // false
	println(isAllUnique("абвгД"))     // true
}
