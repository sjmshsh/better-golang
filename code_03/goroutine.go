package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Println("recover panic")
			}
			wg.Done()
		}()
		fmt.Println("Hello goroutine...")
		panic("err")
	}()
	wg.Wait()
	fmt.Println("Hello main ... ")
}
