package main

import (
	"encoding/json"
	"fmt"
)

// 基本类型的默认值会被缺省, 除了具有默认长度的数组
// 指针类型为nil的时候会被缺省
type Student struct {
	Id      int       `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Age     int       `json:"age,omitempty"`
	Address *string   `json:"address,omitempty"`
	Like    []string  `json:"list,omitempty"`
	Love    [1]string `json:"love,omitempty"`
}

func main() {
	// 只赋值Id、Name，不赋值Age属性
	stu := Student{
		Id:   1,
		Name: "zs",
	}
	bytes, _ := json.Marshal(stu)
	fmt.Println(string(bytes))
}
