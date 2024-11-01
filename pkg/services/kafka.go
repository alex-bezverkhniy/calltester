/*
Copyright © 2024 Alexandr Bezverkhniy <alexandr.bezverkhniy@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package services

import (
	"fmt"
	"os"

	c_kafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/cobra"
)

type (
	KafkaService interface {
		Subscribe() error
		Publish(data []byte) error
	}
	KafkaServiceImpl struct {
		topic   string
		host    string
		port    int
		verbose bool
	}
)

func NewKafkaServiceByCommand(cmd *cobra.Command) (KafkaService, error) {
	var host string
	var port int
	var topic string
	var verbose bool
	var err error

	if host, err = cmd.Flags().GetString("host"); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return nil, err
	}
	if port, err = cmd.Flags().GetInt("port"); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return nil, err
	}
	if topic, err = cmd.Flags().GetString("topic"); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return nil, err
	}
	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return nil, err
	}

	kafkaService := NewKafkaService(topic, host, port, verbose)
	return kafkaService, nil
}

func NewKafkaService(topic, host string, port int, verbose bool) KafkaService {
	return &KafkaServiceImpl{topic, host, port, verbose}
}

// Subscribe - Consumes messages from kafka topic
func (k *KafkaServiceImpl) Subscribe() error {
	if k.verbose {
		fmt.Println("kafka subscribe to:")
		fmt.Printf("\tbroker: %s:%d\n", k.host, k.port)
		fmt.Printf("\ttopic: %s\n", k.topic)
	}

	fmt.Println("Consuming messages...")
	fmt.Println("To exit press CTRL+C")
	consumer, err := c_kafka.NewConsumer(&c_kafka.ConfigMap{
		"bootstrap.servers": k.getUrl(),
		"group.id":          "calltester-sub",
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics([]string{k.topic}, nil)
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
			fmt.Printf("Message on - %s:\n%s\n", k.topic, string(msg.Value))
		case c_kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			// fmt.Printf("Ignored %v\n", e)
		}
	}

	consumer.Close()
	return nil
}

// Publish - Publishes message to kafka topic
func (k *KafkaServiceImpl) Publish(data []byte) error {
	if k.verbose {
		fmt.Println("kafka publish to:")
		fmt.Printf("\tbroker: %s:%d\n", k.host, k.port)
		fmt.Printf("\ttopic: %s\n", k.topic)
		fmt.Printf("\tdata: %s\n", data)
	}
	producer, err := c_kafka.NewProducer(&c_kafka.ConfigMap{
		"bootstrap.servers": k.getUrl(),
		"client.id":         "calltester-pub",
		"acks":              "all",
	})
	defer producer.Close()
	if err != nil {
		panic(err)
	}

	delivery_chan := make(chan c_kafka.Event, 10000)

	// TODO: add key
	err = producer.Produce(&c_kafka.Message{
		TopicPartition: c_kafka.TopicPartition{Topic: &k.topic, Partition: c_kafka.PartitionAny},
		Value:          data,
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
	return nil
}

func (k *KafkaServiceImpl) getUrl() string {
	return fmt.Sprintf("%s:%d", k.host, k.port)
}
