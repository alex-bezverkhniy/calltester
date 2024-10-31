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

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Subscribe to Kafka topic",
	Long:  `You can use it to subscribe/consume message from Kafka topic.`,
	Run: func(cmd *cobra.Command, args []string) {
		kafkaService, err := services.NewKafkaServiceByCommand(cmd)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			return
		}
		kafkaService.Subscribe()
	},
}

func init() {
	kafkaCmd.AddCommand(subCmd)
}
