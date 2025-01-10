package main

import (
	"encoding/json"
	"fmt"
)

type Inner struct {
	Name string `json:"name,omitempty"`
}

type Outer struct {
	InnerField Inner `json:"inner_field,omitempty"`
}

func main() {
	outer := Outer{}
	// 使用omitempty，InnerField是零值，所以输出时会被忽略
	data, _ := json.Marshal(outer)
	fmt.Println(string(data)) // 输出：{}
}
