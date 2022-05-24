package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Game represents the state of the game
type Game struct {
	Red    []string
	Blue   []string
	Yellow []string
	Black  []string
	Board  []string
}

func NewGame() (*Game, error) {
	words, err := read()
	if err != nil {
		return nil, fmt.Errorf("could not read words, %w", err)
	}

	words = shuffle(words)

	var picked []string
	for i := 0; i < 25; i++ {
		picked = append(picked, words[i])
	}

	return &Game{
		Red:    picked[:9],
		Blue:   picked[9:17],
		Yellow: picked[17:24],
		Black:  picked[24:],
		Board:  shuffle(picked),
	}, nil
}

func read() ([]string, error) {
	f, err := os.Open("words.txt")
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

func shuffle(vals []string) []string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	shuffled := make([]string, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		shuffled[i] = vals[randIndex]
	}

	return shuffled
}
