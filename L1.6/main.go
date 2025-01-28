package main

import (
	"context"
	"sync/atomic"
	"time"
)

func workWithCancel(cancel <-chan struct{}) {
	for {
		select {
		case <-cancel:
			println("stopping worker...")
			return
		default:
			time.Sleep(time.Second)
			println("working...")
		}
	}
}

func workWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			println("stopping worker...")
			return
		default:
			time.Sleep(time.Second)
			println("working...")
		}
	}
}

func workWithCancelFunc() (work func(), cancelFunc func()) {
	cancel := make(chan struct{})
	cancelled := false

	work = func() {
		for {
			select {
			case <-cancel:
				println("stopping worker...")
				return
			default:
				time.Sleep(time.Second)
				println("working...")
			}

		}
	}

	cancelFunc = func() {
		if cancelled {
			return
		}
		cancelled = true
		close(cancel)
	}

	return
}

func flagWorker(stopFlag *int32) {
	for {
		if atomic.LoadInt32(stopFlag) == 1 {
			println("stopping worker...")
			return
		}
		println("working...")
		time.Sleep(time.Second)
	}
}

func main() {
	// 1 способ
	// Использование канала завершения
	println("1")
	cancelChan := make(chan struct{})
	go workWithCancel(cancelChan)

	time.Sleep(3 * time.Second)
	close(cancelChan)

	time.Sleep(time.Second)

	// 2 способ
	// Использование контекста
	// Контекст с завершением
	println("2.1")
	ctx, cancel := context.WithCancel(context.Background())
	go workWithContext(ctx)
	time.Sleep(3 * time.Second)
	cancel()

	time.Sleep(time.Second)

	// Контекст с завершением по таймауту - через заданное время (есть еще по дедлайну, то есть к заданному времени)
	// Можно вызвать cancel до истечения таймера
	println("2.2")
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	go workWithContext(ctx)
	_ = cancel
	time.Sleep(4 * time.Second)

	// 3 способ
	// Рукописный cancel() с возвращаемой функцией work
	println("3")
	work, cancel := workWithCancelFunc()
	go work()
	time.Sleep(3 * time.Second)
	cancel()

	time.Sleep(1 * time.Second)

	// 4 способ
	// Горутина сама остановится, если в ней нет бесконечного цикла или она не заблокирована
	println("4")
	go func() {
		println("working...")
		time.Sleep(time.Second)
		println("working...")
		time.Sleep(time.Second)
		println("working...")
		time.Sleep(time.Second)
	}()
	time.Sleep(3 * time.Second)

	// 5 способ
	// Использование флагов (через атомики или мьютексы)
	println("5")
	var stopFlag int32

	go flagWorker(&stopFlag)

	time.Sleep(3 * time.Second)
	atomic.StoreInt32(&stopFlag, 1)
	time.Sleep(time.Second)

	//+ Остановка по закрытию канала данных: Если горутина читает данные из канала, ее можно остановить, закрыв этот канал. Полезно, если горутина предназначена для обработки данных из канала.
}
