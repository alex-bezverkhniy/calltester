package main

import (
	"os"

	"calltester/internal/commands/help"
	"calltester/internal/commands/kafka"
	"calltester/internal/models"

	flag "github.com/spf13/pflag"
)

const DefaultCmd = "help"

var (
	pubCmd = flag.NewFlagSet("pub", flag.ExitOnError)
	subCmd = flag.NewFlagSet("sub", flag.ExitOnError)
)

func main() {
	var host string
	flag.StringVar(&host, "host", "localhost", "kafka broker host")
	port := flag.Int("port", 9092, "kafka broker port")
	topic := flag.String("topic", "test", "kafka topic")
	data := flag.String("data", `{"test_key": "test_value"}`, "kafka data")

	flag.Parse()

	regCommands := map[string]models.Command{
		"help": {
			Description: "print this help",
			Handler:     help.Usage,
		},
		"pub": {
			Description: "publish message to kafka",
			Handler:     kafka.Pub,
			Command:     pubCmd,
		},
		"sub": {
			Description: "subscribe message from kafka",
			Handler:     kafka.Sub,
			Command:     subCmd,
		},
	}

	cfg := &models.Config{
		Host:               host,
		Port:               *port,
		Topic:              *topic,
		Data:               *data,
		RegisteredCommands: regCommands,
	}

	cmdName := DefaultCmd
	if c := flag.Arg(0); c != "" {
		cmdName = c
	}

	cmdToRun, ok := regCommands[cmdName]
	if !ok {
		help.Usage(cfg)
		os.Exit(1)
	}

	cmdToRun.Handler(cfg)
}
