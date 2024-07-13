package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// 一个接收者，N 个发送者，唯一的接收者通过关闭一个额外的信号通道说“请停止发送更多内容
func main() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	dataCh := make(chan int)
	stopCh := make(chan struct{})

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				// The try-receive operation is to try
				// to exit the goroutine as early as
				// possible. For this specified example,
				// it is not essential.
				select {
				case <-stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first
				// branch in the second select may be
				// still not selected for some loops if
				// the send to dataCh is also unblocked.
				// But this is acceptable for this
				// example, so the first select block
				// above can be omitted.
				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}

	// receiver
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == Max-1 {
				close(stopCh)
				return
			}
			log.Println(value)
		}
	}()

	wgReceivers.Wait()
}
