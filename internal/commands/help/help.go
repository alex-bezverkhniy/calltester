package help

import (
	"fmt"
	"os"
	"sort"

	"calltester/internal/models"

	flag "github.com/spf13/pflag"
)

func Usage(cfg *models.Config) {
	fmt.Fprintf(os.Stderr, "Usage: calltester [COMMAND=%s]\n", cfg.DefaultCmd)
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Commands:")
	names := []string{}
	for name := range cfg.RegisteredCommands {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Fprintf(os.Stderr, "  %-16s %s\n", name, cfg.RegisteredCommands[name].Description)
	}
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Global Options:")
	flag.PrintDefaults()
}
