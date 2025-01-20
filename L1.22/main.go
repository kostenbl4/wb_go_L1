package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаем большие числа
	a := big.NewInt(1 << 21) // 2^21
	b := big.NewInt(1 << 22) // 2^22

	// Сложение
	sum := new(big.Int).Set(a) // Копируем a в sum
	sum.Add(sum, b)
	fmt.Printf("Сумма: %s\n", sum.String())

	// Вычитание
	diff := new(big.Int).Set(a) // Копируем a в diff
	diff.Sub(diff, b)
	fmt.Printf("Разность: %s\n", diff.String())

	// Умножение
	product := new(big.Int).Set(a) // Копируем a в product
	product.Mul(product, b)
	fmt.Printf("Произведение: %s\n", product.String())

	// Деление
	quotient := new(big.Int).Set(b) // Копируем b в quotient
	quotient.Div(quotient, a)
	fmt.Printf("Частное: %s\n", quotient.String())
}
