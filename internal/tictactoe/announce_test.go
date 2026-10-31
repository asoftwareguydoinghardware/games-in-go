package tictactoe_test

import (
	ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tictactoe"
	"testing"
)

func TestAssignPlayerNotifiesPlayerOfTheirNumber(t *testing.T) {
	testAssignPlayerNotifiesPlayerOfTheirNumber(t, 0)
	testAssignPlayerNotifiesPlayerOfTheirNumber(t, 1)
}

func testAssignPlayerNotifiesPlayerOfTheirNumber(t *testing.T, playerNum int) {
	t.Logf("testAssignPlayerNotifiesPlayerOfTheirNumber(%d)", playerNum)
	g := ttt.New()
	player := newMockPlayer(0)

	g.AssignPlayer(playerNum, player)

	if !player.assignedPlayerNumber {
		t.Errorf("g.AssignPlayer(%d, player) did not call player.AssignPlayerNumber", playerNum)
	}

	if have, want := player.playerNum, playerNum; have != want {
		t.Errorf("g.AssignPlayer(%d, player) did not call player.AssignPlayer() with %d, called with %d", playerNum, want, have)
	}

}

func TestAnnounceMoveCalled(t *testing.T) {
	const initialPlayer = 0

	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	g.InitialPlayer(initialPlayer)
	g.Play()

	t.Logf("Initial player = %d, player0 moves = %v, player1 moves = %v", initialPlayer, p0.moves, p1.moves)
	if !p0.announceMoveCalled {
		t.Errorf("Play method did not call AnnounceMove method.")
	}
}
