package tictactoe_test

import (
	ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tictactoe"
	"testing"
)

func TestNewReturnsNonNilValue(t *testing.T) {
	var game *ttt.Game

	game = ttt.New()
	if nil == game {
		t.Errorf("New() returns nil")
	}
}
