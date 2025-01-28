package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Счетчик с использованием мьютекса
type MutexCounter struct {
	mu    sync.Mutex
	value int
}

func (c *MutexCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *MutexCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

type AtomicCounter struct {
	value atomic.Int64
}

func (ac *AtomicCounter) Increment() {
	ac.value.Add(1)
}

func (ac *AtomicCounter) Value() int {
	return int(ac.value.Load())
}

func main() {
	var wg sync.WaitGroup
	counter := &MutexCounter{}

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
	fmt.Println("Финальное значение счетчика с мьютексом:", counter.Value())

	// Атомарный счетчик
	atomicCounter := &AtomicCounter{}

	// Запуск 1000 горутин, каждая из которых инкрементирует счетчик
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomicCounter.Increment()
		}()
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод итогового значения счетчика
	fmt.Println("Финальное значение атомарного счетчика:", atomicCounter.Value())
}
