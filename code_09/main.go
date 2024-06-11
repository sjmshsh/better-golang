package main

import (
	"fmt"
)

func TestPointSlice() {
	a := "A"
	b := "B"
	c := "C"

	list := make([]string, 0)
	list = append(list, a)
	list = append(list, b)
	list = append(list, c)

	fmt.Printf("list = %+v \n", list)

	cList := make([]*string, 0)
	for _, str := range list {
		cList = append(cList, &str)
	}

	fmt.Printf("cList = [%+v %+v %+v] \n", *cList[0], *cList[1], *cList[2])
}
func main() {
	TestPointSlice()
}
