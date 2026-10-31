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

	if !player.dataFromApp.assignedPlayerNum {
		t.Errorf("g.AssignPlayer(%d, player) did not call player.AssignPlayerNumber", playerNum)
	}

	if have, want := player.dataFromApp.playerNum, playerNum; have != want {
		t.Errorf("g.AssignPlayer(%d, player) did not call player.AssignPlayer() with %d, called with %d", playerNum, want, have)
	}

}

func TestAnnounceMoveCalled(t *testing.T) {
	testAnnounceMoveCalled(t, 0)
	testAnnounceMoveCalled(t, 1)
}

func testAnnounceMoveCalled(t *testing.T, initialPlayer int) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	g.InitialPlayer(initialPlayer)
	g.Play()

	t.Logf("Initial player = %d, player0 moves = %v, player1 moves = %v", initialPlayer, p0.moves, p1.moves)
	if len(p0.dataFromApp.announcedMoves) == 0 {
		t.Errorf("Play method did not call AnnounceMove method.")
		return
	}

	announcements := p0.dataFromApp.announcedMoves
	if have, want := announcements[0].move, "0,0"; have != want {
		t.Errorf("For player 0, Play method first announced move is %q, want %q", have, want)
	}

	if have, want := announcements[0].player, initialPlayer; have != want {
		t.Errorf("For player 0, Play method first announced move by player %d, want %d", have, want)
	}
}
