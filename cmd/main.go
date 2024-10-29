package main

import (
	"fmt"
	"os"
	"sort"

	"calltester/internal"
	"calltester/internal/commands/kafka"

	flag "github.com/spf13/pflag"
)

type cmd struct {
	description string
	handler     func(*internal.Config)
	cmd         *flag.FlagSet
}

var (
	pubCmd = flag.NewFlagSet("pub", flag.ExitOnError)
	subCmd = flag.NewFlagSet("sub", flag.ExitOnError)
	cmds   = map[string]cmd{
		"pub": {
			description: "publish message to kafka",
			handler:     kafka.Pub,
			cmd:         pubCmd,
		},
		"sub": {
			description: "subscribe message from kafka",
			handler:     kafka.Sub,
			cmd:         subCmd,
		},
	}
)

const defaultCmd = "pub"

func main() {
	var host string
	flag.StringVar(&host, "host", "localhost", "kafka broker host")
	port := flag.Int("port", 9092, "kafka broker port")
	topic := flag.String("topic", "test", "kafka topic")
	data := flag.String("data", `{"test_key": "test_value"}`, "kafka data")

	flag.Parse()

	cfg := &internal.Config{
		Host:  host,
		Port:  *port,
		Topic: *topic,
		Data:  *data,
	}

	cmdName := defaultCmd
	if c := flag.Arg(0); c != "" {
		cmdName = c
	}

	cmdToRun, ok := cmds[cmdName]
	if !ok {
		usage()
		os.Exit(1)
	}

	cmdToRun.handler(cfg)
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: calltester [COMMAND=%s]\n", defaultCmd)
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Commands:")
	names := []string{}
	for name := range cmds {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Fprintf(os.Stderr, "  %-16s %s\n", name, cmds[name].description)
	}
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Global Options:")
	flag.PrintDefaults()
}
