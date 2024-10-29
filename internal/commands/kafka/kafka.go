package kafka

import (
	"fmt"
	"os"

	"calltester/internal"

	c_kafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func Pub(cfg *internal.Config) {
	producer, err := c_kafka.NewProducer(&c_kafka.ConfigMap{
		"bootstrap.servers": cfg.GetUrl(),
		"client.id":         "calltester-pub",
		"acks":              "all",
	})
	defer producer.Close()
	if err != nil {
		panic(err)
	}

	delivery_chan := make(chan c_kafka.Event, 10000)

	err = producer.Produce(&c_kafka.Message{
		TopicPartition: c_kafka.TopicPartition{Topic: &cfg.Topic, Partition: c_kafka.PartitionAny},
		Value:          []byte(cfg.Data),
	},
		delivery_chan,
	)

	e := <-delivery_chan
	m := e.(*c_kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
	close(delivery_chan)
}

func Sub(cfg *internal.Config) {
	fmt.Printf("Starting consumer for topic: %s\n", cfg.Topic)
	fmt.Println("To exit press CTRL+C")
	consumer, err := c_kafka.NewConsumer(&c_kafka.ConfigMap{
		"bootstrap.servers": cfg.GetUrl(),
		"group.id":          "calltester-sub",
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics([]string{cfg.Topic}, nil)
	run := true
	for run == true {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *c_kafka.Message:
			msg := ev.(*c_kafka.Message)
			_, err := consumer.CommitMessage(msg)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Message on - %s:\n%s\n", cfg.Topic, string(msg.Value))
		case c_kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			// fmt.Printf("Ignored %v\n", e)
		}
	}

	consumer.Close()
}
