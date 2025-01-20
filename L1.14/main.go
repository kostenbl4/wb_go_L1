package main

import (
	"fmt"
	"reflect"
)

func getIfaceValue(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Println("Type: integer\tValue: ", v)
	case string:
		fmt.Println("Type: string\tValue: ", v)
	case bool:
		fmt.Println("Type: boolean\tValue: ", v)
	default:
		// Используем reflect для определения типов каналов
		rType := reflect.TypeOf(val)
		if rType.Kind() == reflect.Chan {
			// Определяем тип элементов канала
			fmt.Printf("Type: channel of %s\tValue: %v\n", rType.Elem(), val)
		} else {
			fmt.Println("Type: unknown\tValue: ", v)
		}
	}
}

func main() {
	// Тестирование функции getIfaceValue
	// Передаем разные типы данных
	// и выводим их тип и значение
	getIfaceValue(10) // integer

	getIfaceValue("Hello") // string

	getIfaceValue(true) // boolean

	chInt := make(chan int)
	defer close(chInt)
	getIfaceValue(chInt) // channel of int

	chIface := make(chan interface{})
	defer close(chIface)
	getIfaceValue(chIface) // channel of interface{}
}