package main

import "fmt"

const (
	i1 = iota
	i2
	_ //丢弃该值
	i4
)

const (
	n1 = iota * 10
	n2
	n3
)

func main() {
	fmt.Println(i1)
	fmt.Println(i2)
	fmt.Println(i4)

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
}
