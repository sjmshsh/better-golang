package main

import (
	"fmt"
	"time"
)

func longRunning(messages <-chan string) {
	for {
		select {
		case <-time.After(time.Minute):
			return
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}

// 接受消息并打印, 同时会在一分钟内一直等待新的消息
// 如果在一分钟内没有收到新的消息, 函数会自动退出
func longRunning1(messages <-chan string) {
	timer := time.NewTimer(time.Minute)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			return
		case msg := <-messages:
			fmt.Println(msg)

			// 这个代码用于丢弃/耗尽在执行第二个分支代码块时在短时间内发送的潜在计时器通
			if !timer.Stop() {
				<-timer.C
			}
		}

		timer.Reset(time.Minute)
	}
}

func main() {

}
