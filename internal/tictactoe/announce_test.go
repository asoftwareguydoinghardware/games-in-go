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
