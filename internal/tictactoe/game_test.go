package tictactoe_test

import (
	"context"
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

func TestPlayAnnouncesInitialPlayer(t *testing.T) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)
	assignTracker(p0, p1)
	g.AssignPlayer(context.Background(), 0, p0)
	g.AssignPlayer(context.Background(), 1, p1)
	g.InitialPlayer(0)

	failIfAnnounced(t, p0)
	failIfAnnounced(t, p1)

	g.PlayContext(context.TODO())

	failIfNotAnnounced(t, p0)
	failIfNotAnnounced(t, p1)
}

func failIfAnnounced(t *testing.T, m *mockPlayer) {
	if m.dataFromApp.announcedInitialPlayer {
		t.Errorf("Announced initial player to player %d before Play() called", m.playerNum)
	}
}

func failIfNotAnnounced(t *testing.T, m *mockPlayer) {
	if !m.dataFromApp.announcedInitialPlayer {
		t.Errorf("Did not Announce initial player to player %d after Play() called", m.playerNum)
	}
}

func TestPlayAnnouncesCorrectInitialPlayer(t *testing.T) {
	testPlayAnnouncesCorrectInitialPlayer(t, 1)
	testPlayAnnouncesCorrectInitialPlayer(t, 0)
}

func testPlayAnnouncesCorrectInitialPlayer(t *testing.T, initialPlayerNum int) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)
	assignTracker(p0, p1)
	g.AssignPlayer(context.Background(), 0, p0)
	g.AssignPlayer(context.Background(), 1, p1)
	g.InitialPlayer(initialPlayerNum)

	g.PlayContext(context.TODO())

	failIfDidNotAnnounceCorrectPlayer(t, p0, initialPlayerNum)
	failIfDidNotAnnounceCorrectPlayer(t, p1, initialPlayerNum)
}

func failIfDidNotAnnounceCorrectPlayer(t *testing.T, m *mockPlayer, want int) {
	if m.dataFromApp.initialPlayer != want {
		t.Errorf("Play() announced wrong initial player to player %d, have %d, want %d", m.playerNum, m.dataFromApp.initialPlayer, want)
	}
}

func TestDefaultInitialPlayerIsPlayer0(t *testing.T) {
	const defaultInitialPlayer = 0

	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)
	assignTracker(p0, p1)
	g.AssignPlayer(context.Background(), 0, p0)
	g.AssignPlayer(context.Background(), 1, p1)
	// no assignment of initial player
	g.PlayContext(context.TODO())

	failIfDidNotAnnounceCorrectPlayer(t, p0, defaultInitialPlayer)
	failIfDidNotAnnounceCorrectPlayer(t, p1, defaultInitialPlayer)
}

func TestPlayCallsMoveForPlayer0(t *testing.T) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)
	assignTracker(p0, p1)
	g.AssignPlayer(context.Background(), 0, p0)
	g.AssignPlayer(context.Background(), 1, p1)
	g.PlayContext(context.TODO())

	if !p0.dataFromApp.moveCalled {
		t.Errorf("Play() did not call Move() for player %d", p0.playerNum)
	}
}

func TestPlayCallsMoveForPlayer1(t *testing.T) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)
	assignTracker(p0, p1)
	g.AssignPlayer(context.Background(), 0, p0)
	g.AssignPlayer(context.Background(), 1, p1)
	g.PlayContext(context.TODO())

	if !p1.dataFromApp.moveCalled {
		t.Errorf("Play() did not call Move() for player %d", p1.playerNum)
	}
}

func TestPlayCallsMoveForInitialPlayer(t *testing.T) {
	testPlayCallsMoveForInitialPlayer(t, 1)
	testPlayCallsMoveForInitialPlayer(t, 0)
}

func testPlayCallsMoveForInitialPlayer(t *testing.T, initialPlayer int) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)
	assignTracker(p0, p1)
	g.AssignPlayer(context.Background(), 0, p0)
	g.AssignPlayer(context.Background(), 1, p1)
	g.InitialPlayer(initialPlayer)
	g.PlayContext(context.TODO())

	if p0.tracker.mockPlayerForFirstMove != initialPlayer {
		t.Errorf("Have first move by player %d, want %d", p0.tracker.mockPlayerForFirstMove, initialPlayer)
	}
}

func TestPlayAlternatesPlayers(t *testing.T) {
	testPlayAlternatesPlayers(t, 1)
}

func testPlayAlternatesPlayers(t *testing.T, initialPlayer int) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)
	assignTracker(p0, p1)
	g.AssignPlayer(context.Background(), 0, p0)
	g.AssignPlayer(context.Background(), 1, p1)
	g.InitialPlayer(initialPlayer)
	g.PlayContext(context.TODO())

	if !p0.tracker.mockPlayersAlternated {
		t.Errorf("Play() did not alternate between players")
	}
}
