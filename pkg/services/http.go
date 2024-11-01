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
