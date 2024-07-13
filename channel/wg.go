package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var x int32 = 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&x, 1)
			wg.Done()
		}()
	}
	fmt.Println("wait...")
	wg.Wait()
	fmt.Println(atomic.LoadInt32(&x))
}
