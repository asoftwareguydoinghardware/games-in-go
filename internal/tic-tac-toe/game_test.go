package tic_tac_toe_test

import (
	ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tic-tac-toe"
	"testing"
)

func TestInitialCallToDoneReturnsFalse(t *testing.T) {
	player0 := newMockPlayerIO()
	player1 := newMockPlayerIO()
	g := ttt.New()
	g.SetPlayerIO(0, player0)
	g.SetPlayerIO(1, player1)
	g.InitializeGame(0)

	have, want := g.Done(), false

	if have != want {
		t.Errorf("Initial call to Done = %v\n", have)
	}
}

func TestInitializeGameNotifiesBothPlayers(t *testing.T) {
	player0 := newMockPlayerIO()
	player1 := newMockPlayerIO()
	g := ttt.New()
	g.SetPlayerIO(0, player0)
	g.SetPlayerIO(1, player1)
	g.InitializeGame(0)

	if !player0.notifiedOfGameStart {
		t.Errorf("Player 0 not notified of game start")
	}
	if !player1.notifiedOfGameStart {
		t.Errorf("Player 1 not notified of game start")
	}
}
