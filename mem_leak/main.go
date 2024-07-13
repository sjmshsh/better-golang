package main

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

func demo() {
	s := createStringWithLengthOnHeap(1 << 20) // 1M bytes
	f(s)
}

func createStringWithLengthOnHeap(cap uint64) string {
	data := make([]byte, int(cap))

	return string(data)
}
