package services

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type (
	HttpService interface {
		MakeRequest() error
	}
	HttpServiceImpl struct {
		url     string
		method  string
		data    string
		proxy   string
		verbose bool
	}
)

func NewHttpServiceByCommand(cmd *cobra.Command) (HttpService, error) {
	var url string
	var method string
	var data string
	var proxy string
	var verbose bool
	var err error

	if url, err = cmd.Flags().GetString("url"); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return nil, err
	}
	if method, err = cmd.Flags().GetString("method"); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return nil, err
	}
	if data, err = cmd.Flags().GetString("data"); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return nil, err
	}
	if proxy, err = cmd.Flags().GetString("proxy"); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return nil, err
	}
	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return nil, err
	}
	s := NewHttpService(url, method, data, proxy, verbose)

	return s, nil
}

func NewHttpService(url, method, data, proxy string, verbose bool) HttpService {
	return &HttpServiceImpl{url, method, data, proxy, verbose}
}

func (s *HttpServiceImpl) MakeRequest() error {
	if s.verbose {
		fmt.Println("http request:")
		fmt.Printf("\turl: %s\n", s.url)
		fmt.Printf("\tmethod: %s\n", s.method)
		fmt.Printf("\tdata: %s\n", s.data)
		fmt.Printf("\tproxy: %s\n", s.proxy)
	}
	return nil
}
