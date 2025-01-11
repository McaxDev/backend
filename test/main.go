package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type TestStruct struct {
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func main() {
	// 创建一个包含零值 time.Time 的结构体实例
	data := TestStruct{
		Name: "Test",
	}

	// 序列化为 JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON Marshal Error:", err)
		return
	}

	// 打印序列化后的 JSON 字符串
	fmt.Println("Serialized JSON:", string(jsonData))

	// 验证 Timestamp 字段是否出现在 JSON 中
	if string(jsonData) == `{"name":"Test"}` {
		fmt.Println("Test Passed: Zero value time.Time with omitempty is excluded from JSON.")
	} else {
		fmt.Println("Test Failed: Zero value time.Time with omitempty is included in JSON.")
	}
}
