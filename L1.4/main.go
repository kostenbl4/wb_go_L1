package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(wg *sync.WaitGroup, in <-chan any, ctx context.Context, id int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d exiting\n", id)
			return
		case val, ok := <-in:
			if !ok { // Канал закрыт, завершаем воркер
				fmt.Printf("worker %d exiting (channel closed)\n", id)
				return
			}
			fmt.Printf("worker %d got value %v\n", id, val)
		}
	}
}

func main() {
	var n int
	fmt.Print("Enter the number of workers: ")
	fmt.Scanln(&n)

	mainChannel := make(chan any)

	var wg sync.WaitGroup
	wg.Add(n)

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < n; i++ {
		go worker(&wg, mainChannel, ctx, i)
	}

	// Горутина для записи данных в канал
	go func() {
		i := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				mainChannel <- i
				i++
				time.Sleep(500 * time.Millisecond) // Задержка для наглядности
			}
		}
	}()

	// Обработка Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c // Блокируем main до получения сигнала

	fmt.Println("Shutting down...")
	cancel() // Отменяем контекст, чтобы воркеры завершились

	close(mainChannel) // Закрываем канал после отмены контекста
	wg.Wait()          // Ждем завершения всех воркеров

	fmt.Println("All workers exited.")
}
