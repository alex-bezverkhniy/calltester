package models

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type (
	Command struct {
		Description string
		Handler     func(*Config)
		Command     *flag.FlagSet
	}
	Config struct {
		Host               string
		Port               int
		Topic              string
		Data               string
		RegisteredCommands map[string]Command
		DefaultCmd         string
	}
)

func (c Config) GetUrl() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
