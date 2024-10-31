/*
Copyright © 2024 Alexandr Bezverkhniy <alexandr.bezverkhniy@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Perform rests to Kafka (MQ)",
	Long: `Use this command to test Kafka broker. 
	You can use it to produce and consume messages for/from Kafka topics.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kafka called")
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)

	kafkaCmd.PersistentFlags().StringP("host", "", "localhost", "Hostname of the Kafka broker")
	kafkaCmd.PersistentFlags().IntP("port", "", 9092, "Port of the Kafka broker")
	kafkaCmd.PersistentFlags().StringP("topic", "", "test.topic", "Kafka topic name")
}
