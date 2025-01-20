package main

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

func main() {
	// Примеры использования функции reverse
	word := "Hello World!"
	reversed := reverse(word)
	println(reversed)

	word = "Привет, Мир!"
	reversed = reverse(word)
	println(reversed)
}
