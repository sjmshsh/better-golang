package main

import "strings"

var s0 string // a package-level variable

func f(s1 string) {
	s0 = s1[:50]
	// Now, s0 shares the same underlying memory block
	// with s1. Although s1 is not alive now, but s0
	// is still alive, so the memory block they share
	// couldn't be collected, though there are only 50
	// bytes used in the block and all other bytes in
	// the block become unavailable.
}

// 为了避免这种内存泄漏，我们可以将子字符串转换为一个 []byte值，然后将该[]byte值转换回string。
func f1(s1 string) {
	s0 = string([]byte(s1[:50]))
}

// 编译器优化之后可能会变得无效
func f2(s1 string) {
	s0 = (" " + s1[:50])[1:]
}

// Go1.18 开始提供了Clone来进行克隆字符串
func f3(s1 string) {
	s0 = strings.Clone(s1)
}

func demo() {
	s := createStringWithLengthOnHeap(1 << 20) // 1M bytes
	f(s)
}

func createStringWithLengthOnHeap(cap uint64) string {
	data := make([]byte, int(cap))

	return string(data)
}
