package main

import "fmt"

// Target — это интерфейс, с которым ожидает работать клиент
type Target interface {
	Request() string
}

// Adaptee — это существующий класс с другим интерфейсом
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Adaptee: специфичный запрос."
}

// Adapter делает Adaptee совместимым с интерфейсом Target
type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee: adaptee}

	fmt.Println("Client: Я могу работать с адаптером:")
	fmt.Println(adapter.Request())
}
