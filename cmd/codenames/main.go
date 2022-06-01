package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/murat/go-boilerplate/internal/api"
)

var (
	port string
)

func main() {
	if err := run(os.Args); err != nil {
		log.Printf("could not start app, %v", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.StringVar(&port, "port", "8080", "-port 8080")
	if err := flags.Parse(args[1:]); err != nil {
		return fmt.Errorf("could not parse flags, err: %w", err)
	}

	return api.Start(port)
}
