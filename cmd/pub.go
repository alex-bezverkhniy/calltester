/*
Copyright © 2024 Alexandr Bezverkhniy <alexandr.bezverkhniy@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/alex-bezverkhniy/calltester/pkg/services"
	"github.com/spf13/cobra"
)

// pubCmd represents the pub command
var pubCmd = &cobra.Command{
	Use:   "pub",
	Short: "Publish message to Kafka",
	Long: `Publish/Produce message to Kafka. 
	You can use it to publish messages to Kafka topics.`,
	Run: func(cmd *cobra.Command, args []string) {
		kafkaService, err := services.NewKafkaServiceByCommand(cmd)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			return
		}
		var data string
		if data, err = cmd.Flags().GetString("data"); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			return
		}
		err = kafkaService.Publish([]byte(data))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			return
		}
	},
}

func init() {
	kafkaCmd.AddCommand(pubCmd)
	pubCmd.Flags().StringP("data", "", `{"test": "test"}`, "Data to publish to the topic")
}
