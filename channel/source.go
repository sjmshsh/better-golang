package main

import (
	"fmt"
	"math/rand"
	"time"
)

func source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3)+1
	// Sleep 1s/2s/3s.
	time.Sleep(time.Duration(rb) * time.Second)
	c <- ra
}

func main() {
	rand.Seed(time.Now().UnixNano())

	startTime := time.Now()
	// 注意，如果有N 个源，则通信通道的容量必须至少为N-1，以避免与丢弃的响应相对应的 goroutine 被永远阻塞。
	c := make(chan int32, 5)
	for i := 0; i < cap(c); i++ {
		go source(c)
	}
	// Only the first response will be used.
	rnd := <-c
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)
}
