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

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Send HTTP requests",
	Long:  `You can use it to test HTTP server.`,
	Run: func(cmd *cobra.Command, args []string) {
		httpService, err := services.NewHttpServiceByCommand(cmd)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			return
		}

		if err := httpService.MakeRequest(); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	httpCmd.PersistentFlags().StringP("url", "u", "http://localhost", "URL to send requests to")
	httpCmd.PersistentFlags().StringP("data", "d", `{"test": "test"}`, "Request data")
	httpCmd.PersistentFlags().StringP("proxy", "p", "", "Proxy URL")
	httpCmd.PersistentFlags().StringP("method", "m", "GET", "Method of the request (GET, POST, PUT, DELETE, PATCH, HEAD)")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
