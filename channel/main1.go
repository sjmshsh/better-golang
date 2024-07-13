package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest1(r chan<- int32) {
	// Simulate a workload.
	time.Sleep(time.Second * 3)
	r <- rand.Int31n(100)
}

func sumSquares1(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20

	ra, rb := make(chan int32), make(chan int32)
	go longTimeRequest1(ra)
	go longTimeRequest1(rb)

	fmt.Println(sumSquares1(<-ra, <-rb))
}
