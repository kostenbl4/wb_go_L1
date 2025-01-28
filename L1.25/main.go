package main

import (
	"fmt"
	"time"
)

func sleep1(dur time.Duration) {
	// time.After возвращает канал, который закрывается через указанное время
	<-time.After(dur)
}

func sleep2(duration time.Duration) {
	// time.Now возвращает текущее время
	// time.Now().Add добавляет указанное время к текущему времени
	end := time.Now().Add(duration)

	// Цикл, который выполняется, пока текущее время меньше времени окончания - неэффективный способ, тк так как постоянно занимает процессорное время
	for time.Now().Before(end) {
		// Пустой цикл для ожидания
	}
}

func sleep3(duration time.Duration) {
	// time.NewTimer возвращает таймер, который отправляет текущее время в канал C через указанное время
	timer := time.NewTimer(duration)
	<-timer.C
}

func sleep4(duration time.Duration) {
	end := time.Now().Add(duration)

	// Цикл, который выполняется, пока текущее время меньше времени окончания - неэффективный способ, тк так как постоянно занимает процессорное время
	for {
		if time.Now().After(end) {
			return
		}
	}
}

func main() {
	sleep1(1 * time.Second)
	fmt.Println("sleep1: Спим 1 секунду")
	sleep2(1 * time.Second)
	fmt.Println("sleep2: Спим 1 секунду")
	sleep3(1 * time.Second)
	fmt.Println("sleep3: Спим 1 секунду")
	sleep4(1 * time.Second)
	fmt.Println("sleep4: Спим 1 секунду")
}
