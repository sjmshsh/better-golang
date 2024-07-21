package main

import "sync"

func main() {
	var l sync.RWMutex
	l.Lock()
	foo(l)
	l.Lock()
	l.Unlock()
}

func foo(l sync.RWMutex) {
	l.Unlock()
}
