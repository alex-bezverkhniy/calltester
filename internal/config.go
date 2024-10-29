package internal

import "fmt"

type Config struct {
	Host  string
	Port  int
	Topic string
	Data  string
}

func (c Config) GetUrl() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
