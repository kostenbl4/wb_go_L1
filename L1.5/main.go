package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func writeChannel(wg *sync.WaitGroup, ctx context.Context, c chan int) {
	defer wg.Done()
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			c <- i
		}
	}
}

func readChannel(wg *sync.WaitGroup, ctx context.Context, c chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(<-c)
		}
	}
}

func main() {
	var n int

	fmt.Scan(&n)

	c := make(chan int, 1)
	defer close(c)

	ctx := context.Background()
	dur := time.Duration(n) * time.Second
	ctx, cancel := context.WithTimeout(ctx, dur)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go writeChannel(&wg, ctx, c)

	go readChannel(&wg, ctx, c)
	
	wg.Wait()
}
