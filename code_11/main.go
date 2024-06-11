package main

import (
	"fmt"
)

// 在Go语言中，指针类型不能作为map的键（key）的主要原因是因为指针的值是动态的，
// 并且可能会发生变化。当使用指针作为map的键时，如果两个指针指向同一个内存地址，
// 它们被认为是相等的，但是如果指针所指向的值发生变化，那么这两个指针就不再相等了。

type Student struct {
	Id   string
	Name string
}

func TestMapPointKey() {
	m := make(map[*Student]struct{})

	m[&Student{Id: "1", Name: "zs"}] = struct{}{}

	_, ok := m[&Student{Id: "1", Name: "zs"}]
	fmt.Println(ok) // false

	stu := &Student{
		Id:   "2",
		Name: "lxu",
	}

	m[stu] = struct{}{}
	_, ok = m[stu]
	fmt.Println(ok) // true
}

func main() {
	TestMapPointKey()
}
