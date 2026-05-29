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

func testHandleValidMoveFromPlayerCallsRequestMoveForCorrectPlayer(t *testing.T, player int) {
	players := [2]*mockPlayer{newMockPlayerIO(), newMockPlayerIO()}
	g := ttt.New()
	g.SetPlayerIO(0, players[0])
	g.SetPlayerIO(1, players[1])
	players[player].moves = []string{"1"}
	g.InitializeGame(0)

	g.HandleValidMoveFromPlayer(player)

	if players[player].currentMove != 1 {
		t.Errorf("HandleValidMoveFromPlayer() did not call RequestMove() for player %d", player)
	}
}

func TestHandleValidMoveFromPlayerCallsRequestMoveForCorrectPlayer(t *testing.T) {
	testHandleValidMoveFromPlayerCallsRequestMoveForCorrectPlayer(t, 0)
	testHandleValidMoveFromPlayerCallsRequestMoveForCorrectPlayer(t, 1)
}

func testHandleValidMoveFromPlayerCallsShareStateChange(t *testing.T, player int) {
	players := [2]*mockPlayer{newMockPlayerIO(), newMockPlayerIO()}
	g := ttt.New()
	g.SetPlayerIO(0, players[0])
	g.SetPlayerIO(1, players[1])
	players[player].moves = []string{"1"}
	g.InitializeGame(0)

	otherPlayer := 0
	if player == 0 {
		otherPlayer = 1
	}

	g.HandleValidMoveFromPlayer(player)

	if !players[otherPlayer].sharedStateChange {
		t.Errorf("HandleValidMoveFromPlayer(%d) did not call ShareStateChange() for player %d", player, otherPlayer)
	}
}

func TestHandleValidMoveFromPlayerCallsShareStateChange(t *testing.T) {
	testHandleValidMoveFromPlayerCallsShareStateChange(t, 0)
	testHandleValidMoveFromPlayerCallsShareStateChange(t, 1)
}
