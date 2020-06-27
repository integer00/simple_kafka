package main

import (
	"context"
	"fmt"
	"time"
	"os"

	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
)

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}

func main() {
	kafkaURL := os.Getenv("KAFKA_URL")
	topic := "my_topic"
	writer := newKafkaWriter(kafkaURL, topic)
	defer writer.Close()
	fmt.Println("start producing ... !!")

	for i := 0;; i++ {
	    msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("Key-%d", i)),
			Value: []byte(fmt.Sprint(uuid.New())),
	    }
	    fmt.Println("sending: ", string(msg.Value))
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)

	}
}