package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/murat/go-utils/slices"
)

// Game represents the state of the game
type Game struct {
	Red    []string
	Blue   []string
	Yellow []string
	Black  []string
	Board  []string
}

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
		Red:    picked[:9],
		Blue:   picked[9:17],
		Yellow: picked[17:24],
		Black:  picked[24:],
		Board:  slices.Shuffle(picked),
	}, nil
}

func (g *Game) Render() {
	fmt.Println(g)
}

func read(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("could not open file, %w", err)
	}
	defer f.Close()

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
