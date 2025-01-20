package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	// Запуск 1000 горутин, каждая из которых инкрементирует счетчик
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод итогового значения счетчика
	fmt.Println("Финальное значение счетчика:", counter.Value())
}
