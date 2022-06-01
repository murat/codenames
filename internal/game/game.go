package game

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/murat/go-utils/slices"
)

// Game represents the state of the game
type Game struct {
	Red   []string
	Blue  []string
	White []string
	Black []string
	Cards []string
}

// NewGame ...
func NewGame(file string) (*Game, error) {
	words, err := read(file)
	if err != nil {
		return nil, fmt.Errorf("could not read words, %w", err)
	}

	words = slices.Shuffle(words)

	var picked []string
	for i := 0; i < 25; i++ {
		picked = append(picked, words[i])
	}

	return &Game{
		Red:   picked[:9],
		Blue:  picked[9:17],
		White: picked[17:24],
		Black: picked[24:],
		Cards: slices.Shuffle(picked),
	}, nil
}

func read(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("could not open file, %w", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("could not close file, %v", err)
		}
	}(f)

	scanner := bufio.NewScanner(f)

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("could not read file, %w", err)
	}

	return words, nil
}
