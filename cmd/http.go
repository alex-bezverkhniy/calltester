/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Send HTTP requests",
	Long:  `You can use it to test HTTP server.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("http called")
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
	httpCmd.PersistentFlags().StringP("method", "m", "GET", "Method of the request (GET, POST, PUT, DELETE, PATCH)")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
