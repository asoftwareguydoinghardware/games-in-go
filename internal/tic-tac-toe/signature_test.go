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
	player0 := newMockPlayerIO()
	player1 := newMockPlayerIO()
	game.SetPlayerIO(0, player0)
	game.SetPlayerIO(1, player1)

	game.InitializeGame(0)
}

func TestDoneMethodExists(t *testing.T) {
	var done bool

	game := ttt.New()
	done = game.Done()
	_ = done
}

func TestHandleValidMoveFromPlayerExists(t *testing.T) {
	game := ttt.New()
	player0 := newMockPlayerIO()
	player1 := newMockPlayerIO()
	game.SetPlayerIO(0, player0)
	game.SetPlayerIO(1, player1)

	game.HandleValidMoveFromPlayer(0)
}

func TestSetPlayerIOMethodExists(t *testing.T) {
	var playerIo ttt.PlayerIO

	game := ttt.New()

	game.SetPlayerIO(0, playerIo)
}
