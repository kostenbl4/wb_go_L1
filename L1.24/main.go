package main

import (
	"fmt"
	"math"
)

// Интерфейс для доступа к методам Point
type Point interface {
	X() float64
	Y() float64
	DistanceTo(other Point) float64
}

// Структура point с инкапсулированными параметрами x и y
type point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) *point {
	return &point{x: x, y: y}
}

// Методы для получения значений x и y
func (p *point) X() float64 {
	return p.x
}

func (p *point) Y() float64 {
	return p.y
}

// Метод для вычисления расстояния до другой точки
func (p *point) DistanceTo(other Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.X(), 2) + math.Pow(p.y-other.Y(), 2))
}

func main() {
	// Пример использования
	var p1, p2 Point = NewPoint(1, 2), NewPoint(4, 6)

	fmt.Printf("Точка 1: x=%.2f, y=%.2f\n", p1.X(), p1.Y())
	fmt.Printf("Точка 2: x=%.2f, y=%.2f\n", p2.X(), p2.Y())

	fmt.Printf("Расстояние между точками: %.2f\n", p1.DistanceTo(p2))
}