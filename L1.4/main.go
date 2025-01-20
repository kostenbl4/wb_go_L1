package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func work(wg *sync.WaitGroup, in chan any, ctx context.Context, id int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Printf("worker %d got value %v\n", id, <-in)
		}
	}
}

func startWorkers(wg *sync.WaitGroup, in chan any, ctx context.Context, n int) {
	for i := 0; i < n; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			go work(wg, in, ctx, i)
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	mainChannel := make(chan any)
	defer close(mainChannel)

	var wg sync.WaitGroup
	wg.Add(n)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	
	go startWorkers(&wg, mainChannel, ctx, n)
	
	go func(){
		for i := 0; ; i++ {
			mainChannel <- i
		}	
	}()
	
	wg.Wait()
}
