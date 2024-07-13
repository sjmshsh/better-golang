package main

import (
	"context"
	"sync"
	"time"
)

func main() {
	c := make(chan int, 10)
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.TODO())

	// 专门关闭的协程
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
		close(c)
	}()

	// senders
	for i := 0; i < 10; i++ {
		go func(ctx context.Context, id int) {
			select {
			case <-ctx.Done():
				return
			case c <- id: // 入队s
			}
		}(ctx, i)
	}

	// receivers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range c {
				_ = v
			}
		}()
	}

	wg.Wait()
}
