package tictactoe_test

import (
	"fmt"
	ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tictactoe"
	"strings"
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

func testHandleValidMoveFromPlayerCallsReportBadMoveSelectionForBadMove(t *testing.T, player int) {
	players := [2]*mockPlayer{newMockPlayerIO(), newMockPlayerIO()}
	g := ttt.New()
	g.SetPlayerIO(0, players[0])
	g.SetPlayerIO(1, players[1])
	players[player].moves = []string{"-1", "0"}
	g.InitializeGame(0)

	g.HandleValidMoveFromPlayer(player)

	if len(players[player].badMoveMsgs) == 0 {
		t.Errorf("Bad move sequence: %v did not call ReportBadMoveSelection()", players[player].moves)
	}
}

func TestHandleValidMoveFromPlayerCallsReportBadMoveSelectionForBadMove(t *testing.T) {
	testHandleValidMoveFromPlayerCallsReportBadMoveSelectionForBadMove(t, 0)
	testHandleValidMoveFromPlayerCallsReportBadMoveSelectionForBadMove(t, 1)
}

func TestHandleValidMoveFromPlayerDoesNotCallReportBadMoveSelectionForGoodMove(t *testing.T) {
	const player = 0

	players := [2]*mockPlayer{newMockPlayerIO(), newMockPlayerIO()}
	g := ttt.New()
	g.SetPlayerIO(0, players[0])
	g.SetPlayerIO(1, players[1])
	players[player].moves = []string{"0"}
	g.InitializeGame(0)

	g.HandleValidMoveFromPlayer(player)

	if len(players[player].badMoveMsgs) != 0 {
		t.Errorf("Good move sequence: %v called ReportBadMoveSelection()", players[player].moves)
	}
}

func TestHandleValidMoveFromPlayerCallsReportBadMoveSelectionForGoodMoveApropriately(t *testing.T) {
	type testCase struct {
		move   string
		isGood bool
	}
	testCases := []testCase{
		{"fred", false},
		{"1", true},
		{"5", true},
		{"9", false},
		{"", false},
		{" 0", true},
		{"-32", false},
		{" 10", false},
		{" 1d", false},
		{" 2 ", true},
		{" 0x1", true},
	}

	for _, tc := range testCases {
		const player = 0

		players := [2]*mockPlayer{newMockPlayerIO(), newMockPlayerIO()}
		g := ttt.New()
		g.SetPlayerIO(0, players[0])
		g.SetPlayerIO(1, players[1])
		players[player].moves = []string{tc.move, "0"}
		g.InitializeGame(0)

		g.HandleValidMoveFromPlayer(player)

		if tc.isGood {
			if len(players[player].badMoveMsgs) != 0 {
				t.Errorf("Good move sequence: %v called ReportBadMoveSelection()", players[player].moves)
			}
		} else {
			have, want := len(players[player].badMoveMsgs), 1
			if have != want {
				t.Errorf("Bad move sequence: %v called ReportBadMoveSelection() %d times, want %d calls", players[player].moves, have, want)
			}
		}
	}
}

func testBadMoveRerequestsMove(t *testing.T, badMoves int) {
	const playerNum = 1

	players := [2]*mockPlayer{newMockPlayerIO(), newMockPlayerIO()}
	g := ttt.New()
	g.SetPlayerIO(0, players[0])
	g.SetPlayerIO(1, players[1])
	player := players[playerNum]
	moves := make([]string, badMoves+1)
	for i := 0; i < badMoves; i++ {
		moves[i] = "zed"
	}
	moves[badMoves] = "0"
	player.moves = moves

	g.InitializeGame(playerNum)

	g.HandleValidMoveFromPlayer(playerNum)

	have, want := len(player.badMoveMsgs), badMoves
	if have != want {
		t.Errorf("HandleValidMoveFromPlayer requeried %d times, want %d", have, want)
	}
}

func TestBadMoveRerequestsMove(t *testing.T) {
	testBadMoveRerequestsMove(t, 2)
	testBadMoveRerequestsMove(t, 20)
}

func testBadMoveResponse(t *testing.T, initialMove string, wantCode int, wantSubstring string) {
	const playerNum = 0

	players := [2]*mockPlayer{newMockPlayerIO(), newMockPlayerIO()}
	g := ttt.New()
	g.SetPlayerIO(0, players[0])
	g.SetPlayerIO(1, players[1])
	player := players[playerNum]
	moves := make([]string, 2)
	moves[0] = initialMove
	moves[1] = "0"
	player.moves = moves

	g.InitializeGame(playerNum)
	g.HandleValidMoveFromPlayer(playerNum)

	badMove := player.badMoveMsgs[0]
	haveCode := badMove.code
	haveString := badMove.msg

	if haveCode != wantCode {
		t.Errorf("for move sequence %v initial move have code %03d, want %03d", moves, haveCode, wantCode)
	}
	if !strings.Contains(haveString, wantSubstring) {
		t.Errorf("for move sequence %v initial move have error string %q want to contain %q", moves, haveString, wantSubstring)
	}
}

func TestBadMoveResponse(t *testing.T) {
	testBadMoveResponse(t, "zed", 500, "nvalid number")
	testBadMoveResponse(t, "zed", 500, "zed")
	testBadMoveResponse(t, "aaron", 500, "aaron")
	testBadMoveResponse(t, "-1", 501, "nvalid move, must be in range 0-8")
	testBadMoveResponse(t, "9", 501, "nvalid move, must be in range 0-8")
	testBadMoveResponse(t, "1a", 500, "nvalid number")

}

func TestDuplicateMoveReportsError(t *testing.T) {
	testDuplicateMoveReportsError(t, 7)
}

func testDuplicateMoveReportsError(t *testing.T, move int) {
	const (
		playerNum      = 0 // baked into diagnostic messages, do not change
		otherPlayerNum = iota
	)
	var usedMoves [9]bool

	findMove := func() int {
		for i := 0; i < len(usedMoves); i++ {
			if !usedMoves[i] {
				return i
			}
		}

		return -1
	}

	players := [2]*mockPlayer{newMockPlayerIO(), newMockPlayerIO()}
	g := ttt.New()
	g.SetPlayerIO(0, players[0])
	g.SetPlayerIO(1, players[1])
	player := players[playerNum]
	otherPlayer := players[otherPlayerNum]

	moves := make([]string, 3)
	moves[0] = fmt.Sprintf("%d", move)
	usedMoves[move] = true
	moves[1] = moves[0]
	freeMove := findMove()
	usedMoves[freeMove] = true
	moves[2] = fmt.Sprintf("%d", freeMove)
	player.moves = moves

	otherMoves := make([]string, 1)
	freeMove = findMove()
	otherMoves[0] = fmt.Sprintf("%d", freeMove)
	otherPlayer.moves = otherMoves

	g.InitializeGame(playerNum)
	g.HandleValidMoveFromPlayer(playerNum)
	g.HandleValidMoveFromPlayer(otherPlayerNum)
	g.HandleValidMoveFromPlayer(playerNum)

	have, want := len(player.badMoveMsgs), 1
	if have != want {
		t.Errorf("For player %d moves %v have %d bad moves want %d", playerNum, moves, have, want)
		if have < want {
			return
		}
	}

	badMoveMsg := player.badMoveMsgs[0]
	haveCode, wantCode := badMoveMsg.code, 502
	if haveCode != wantCode {
		t.Errorf("For player %d moves %v have code %d want %d", playerNum, moves, haveCode, wantCode)
	}

	haveMove, wantMove := badMoveMsg.moveNum, 1
	if haveMove != wantMove {
		t.Errorf("For player %d moves %v; have bad move on query %d want %d", playerNum, moves, haveMove, wantMove)
	}

	haveMsg, wantMsg := badMoveMsg.msg, "quare occupied"
	if !strings.Contains(haveMsg, wantMsg) {
		t.Errorf("For player %d moves %v; have msg %q want to contain %q", playerNum, moves, haveMsg, wantMsg)
	}
}
