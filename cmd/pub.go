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
