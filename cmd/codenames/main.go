package main

import (
	"fmt"
	"log"
	"os"

	"github.com/murat/codenames/internal/game"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Printf("could not start app, %v", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	g, err := game.NewGame(args[0])
	if err != nil {
		return fmt.Errorf("could not start game, %w", err)
	}

	g.Render()

	return nil
}
