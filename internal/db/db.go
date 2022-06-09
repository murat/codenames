package db

import (
	"github.com/murat/go-boilerplate/internal/game"
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm"
)

type Game struct {
	*gorm.Model
	Game game.Game `gorm:"embedded"`
}

type Player struct {
	*gorm.Model
	GameID uint
	Name   string
	Team   string
}

func New(path string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(path), &gorm.Config{})
}
