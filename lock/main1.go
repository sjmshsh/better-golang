package main

import (
	"sync"
	"time"
)

var l sync.RWMutex

// 可读锁内使用可读锁
func readAndRead() {
	l.RLock()
	defer l.RUnlock()

	l.RLock()
	defer l.RUnlock()
}

func lockAndLock() { // 全局锁内使用全局锁
	l.Lock()
	defer l.Unlock()

	l.Lock()
	defer l.Unlock()
}

// 这个也是不可以的
// 全局锁内使用可读锁
func lockAndRead() {
	l.Lock()
	defer l.Unlock()

	l.RLock()
	defer l.RUnlock()
}

// 这个是不可用的
func readAndLock() { // 可读锁内使用全局锁
	l.RLock()
	defer l.RUnlock()

	l.Lock()
	defer l.Unlock()
}

func main() {
	readAndRead()
	// readAndLock()

	// lockAndRead()
	// lockAndLock()

	time.Sleep(5 * time.Second)
}
