package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "topic_test", 0)
	conn.SetReadDeadline(time.Now().Add(time.Second * 10))

	//mes, _ := conn.ReadMessage(1e6)
	//fmt.Println(string(mes.Value))
	batch := conn.ReadBatch(1e3, 1e6)
	bytes := make([]byte, 1e6)

	for {
		_, err := batch.Read(bytes)
		if err != nil {
			break
		}
		fmt.Println(string(bytes))
	}
}
