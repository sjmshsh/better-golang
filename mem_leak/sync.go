package main

import "sync"

type Counter struct {
	sync.Mutex
	n int64
}

func (c *Counter) Increase(d int64) (r int64) {
	c.Lock()
	c.n += d
	r = c.n
	c.Unlock()
	return
}

func (c Counter) Value() (r int64) {
	c.Lock()
	r = c.n
	c.Unlock()
	return
}

// 并发问题
// 在 Value() 方法中,由于使用了值接收者,每次调用方法时都会创建 Counter 结构体的副本。
// 这意味着对副本上的锁操作不会影响原始的 Counter 实例,从而无法正确地应用 sync.Mutex 带来的同步机制。
// 这可能会导致竞争条件和其他并发相关的问题。

// 不必要的内存分配
