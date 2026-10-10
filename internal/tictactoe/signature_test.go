package tictactoe_test

import (
	ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tictactoe"
	"testing"
)

func TestNewExists(t *testing.T) {
	var game *ttt.Game

	game = ttt.New()
	_ = game
}

func TestDoneExists(t *testing.T) {
	var done bool

	g := ttt.New()
	done = g.Done()
	_ = done
}
