package main

import (
	"fmt"
	"math"
)

// Получение ключа для группировки
func getKey(n float64) float64 {
	return n - math.Mod(n, 10)
}

func main() {
	// Исходные данные
	vals := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// Создание словаря для группировки
	groups := make(map[float64][]float64)

	// Группировка значений по ключу
	for _, v := range vals {
		groups[getKey(v)] = append(groups[getKey(v)], v)
	}
	// Вывод результата
	fmt.Println(groups)
}
