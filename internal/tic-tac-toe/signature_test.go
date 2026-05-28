package tic_tac_toe_test

import (
	ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tic-tac-toe"
	"testing"
)

func TestNewExists(t *testing.T) {
	var game *ttt.Game

	game = ttt.New()
	_ = game
}

func TestInitializeGameMethodExists(t *testing.T) {
	game := ttt.New()
	game.InitializeGame(0)
}

func TestDoneMethodExists(t *testing.T) {
	var done bool

	game := ttt.New()
	done = game.Done()
	_ = done
}
