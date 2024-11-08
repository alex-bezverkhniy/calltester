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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type (
	HttpService interface {
		MakeRequest() error
	}
	HttpServiceImpl struct {
		url       string
		method    string
		data      string
		proxy     string
		headers   []string
		verbose   bool
		client    *http.Client
		out       *os.File
		formatter *colorjson.Formatter
	}
)

func NewHttpServiceByCommandAndMethod(cmd *cobra.Command, method string, args []string) (HttpService, error) {
	var url string
	var data string
	var proxy string
	var verbose bool
	var headers []string
	var err error

	if url, err = cmd.Flags().GetString("url"); err != nil {
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
	if headers, err = cmd.Flags().GetStringArray("header"); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return nil, err
	}
	if len(url) == 0 {
		if url, err = tryGetURL(args); err != nil {
			return nil, err
		}
	}
	url = fixURL(url)

	s := NewHttpService(url, method, data, proxy, headers, verbose)

	return s, nil
}

func NewHttpServiceByCommand(cmd *cobra.Command, args []string) (HttpService, error) {
	var method string
	var err error

	if method, err = cmd.Flags().GetString("method"); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return nil, err
	}
	return NewHttpServiceByCommandAndMethod(cmd, method, args)
}

func NewHttpService(url, method, data, proxy string, headers []string, verbose bool) HttpService {
	client := http.Client{}
	f := colorjson.NewFormatter()
	f.KeyColor = color.New(color.FgBlue)
	return &HttpServiceImpl{url, strings.ToUpper(method), data, proxy, headers, verbose, &client, os.Stderr, f}
}

func (s *HttpServiceImpl) MakeRequest() error {
	// Request body
	var dataReader io.Reader
	if len(s.data) > 0 {
		dataReader = bytes.NewReader([]byte(s.data))
	}

	// Request
	req, err := http.NewRequest(s.method, s.url, dataReader)
	if err != nil {
		return err
	}

	// Headers
	for _, h := range s.headers {
		parts := strings.SplitN(h, ":", 2)
		if len(parts) == 2 {
			req.Header.Add(parts[0], parts[1])
		}
	}

	s.printRequest(req)

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}

	return s.printResponse(resp)
}

func (s *HttpServiceImpl) printRequest(req *http.Request) {
	if s.verbose {
		path := req.URL.Path
		if len(path) == 0 {
			path = "/"
		}
		fmt.Fprintf(s.out, "> %s %s %s\n", req.Method, path, req.Proto)
		fmt.Fprintf(s.out, "> Host: %s\n", req.Host)
		for h, v := range req.Header {
			val := ""
			if len(v) > 0 {
				val = v[0]
			}
			fmt.Fprintf(s.out, "> %s: %s\n", h, val)
		}

		for h, v := range req.Trailer {
			val := ""
			if len(v) > 0 {
				val = v[0]
			}
			fmt.Fprintf(s.out, "> %s: %s\n", h, val)
		}
		fmt.Fprintln(s.out, ">")
	}
}

func (s *HttpServiceImpl) printResponse(resp *http.Response) error {
	if s.verbose {
		path := resp.Request.URL.Path
		if len(path) == 0 {
			path = "/"
		}
		fmt.Fprintf(s.out, "< %s %s %s\n", resp.Proto, path, resp.Status)
		for h, v := range resp.Header {
			val := ""
			if len(v) > 0 {
				val = v[0]
			}
			fmt.Fprintf(s.out, "< %s: %s\n", h, val)
		}
		fmt.Fprintln(s.out, "<")
	}
	respBody, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if len(respBody) > 0 {
		// Use jsoncolor for json output
		contentType := resp.Header.Get("Content-Type")
		if strings.HasPrefix(contentType, "application/json") {
			if err = s.printJsonColored(respBody); err != nil {
				return err
			}
		} else {
			fmt.Fprintf(s.out, "%s\n", respBody)
		}
	}
	return nil
}

func (s *HttpServiceImpl) printJsonColored(data []byte) error {
	var obj interface{}
	var d []byte
	var err error
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	s.formatter.Indent = 2
	d, err = s.formatter.Marshal(obj)
	if err != nil {
		return err
	}
	fmt.Fprintln(s.out, string(d))
	return nil
}

func tryGetURL(args []string) (string, error) {
	if len(args) > 0 {
		url := args[0]
		return url, nil
	}
	return "", fmt.Errorf("url is required")
}

func fixURL(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	return url
}
