package main

import (
	"strings"
)

func isAllUnique(s string) bool {
	// Создание множества для хранения уникальных символов
	set := make(map[rune]struct{})
	s = strings.ToLower(s) // Приведение строки к нижнему регистру

	// Перебор всех символов строки
	for _, r := range s {
		if _, ok := set[r]; ok {
			return false
		}
		set[r] = struct{}{}
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
