package main

import (
	"fmt"
	"log"

	sarama "github.com/Shopify/sarama"
)

func main() {
	fmt.Println("Consumer started. Press Ctrl+C to exit")
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("example", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()

	for {
		msg := <-partitionConsumer.Messages()
		log.Printf("Consumed message: \"%s\" at offset: %d\n", msg.Value, msg.Offset)
	}
}
