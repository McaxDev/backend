package main

import "log"

func main() {

	if err := Init(); err != nil {
		log.Fatalf("服务启动失败%v\n", err)
	}

}
