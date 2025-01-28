package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func writeChannel(wg *sync.WaitGroup, ctx context.Context, c chan<- int) {
	defer wg.Done()
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			close(c) // Закрываем канал здесь
			return
		case c <- i:
		}
	}
}

func readChannel(wg *sync.WaitGroup, c <-chan int) {
	defer wg.Done()
	for elem := range c {
		fmt.Println(elem)
	}
}

func main() {
	var n int
	fmt.Print("Enter duration in seconds: ")
	fmt.Scanln(&n)

	c := make(chan int)

	ctx := context.Background()
	dur := time.Duration(n) * time.Second
	ctx, cancel := context.WithTimeout(ctx, dur)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go writeChannel(&wg, ctx, c)
	go readChannel(&wg, c)

	wg.Wait()
	fmt.Println("Program finished")
}
