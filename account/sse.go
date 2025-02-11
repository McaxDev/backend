package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	Clients   = make(map[uint]chan string)
	ClientsMu sync.Mutex
	nextId    uint
)

func BroadCast(msg string) {
	ClientsMu.Lock()
	for _, client := range Clients {
		select {
		case client <- msg:
		default:
			fmt.Println("跳过一个堵塞通道")
		}
	}
	ClientsMu.Unlock()
}

func SSE(c *gin.Context) {

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	client := make(chan string, 10)
	clientId := nextId

	ClientsMu.Lock()
	Clients[clientId] = client
	nextId++
	ClientsMu.Unlock()

	notify := c.Request.Context().Done()

	for {
		select {
		case message := <-client:
			_, _ = fmt.Fprintf(c.Writer, "data: %s\n\n", message)
			c.Writer.Flush()
		case <-notify:
			fmt.Println("客户端断开连接：", clientId)
			ClientsMu.Lock()
			delete(Clients, clientId)
			ClientsMu.Unlock()
			close(client)
			return
		}
	}
}
