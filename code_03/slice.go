package main

import (
	"fmt"
	"unsafe"
)

// 切片作为参数是值传递
func main() {
	list := make([]string, 0)
	list = append(list, "a", "b", "c")
	fmt.Println(list)
	fmt.Println(unsafe.Pointer(&list))
	Test(list)
}

func Test(list []string) {
	fmt.Println(list)
	fmt.Println(unsafe.Pointer(&list))
}
