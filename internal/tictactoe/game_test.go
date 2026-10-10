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

func TestDoneReturnsFalseByDefault(t *testing.T) {
	g := ttt.New()
	done := g.Done()
	if done {
		t.Errorf("Done() returns true on initial call")
	}
}

func TestPlayAnnouncesInitialPlayer(t *testing.T) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)
	g.InitialPlayer(0)

	failIfAnnounced(t, p0)
	failIfAnnounced(t, p1)

	g.Play()

	failIfNotAnnounced(t, p0)
	failIfNotAnnounced(t, p1)
}

func failIfAnnounced(t *testing.T, m *mockPlayer) {
	if m.announcedInitialPlayer {
		t.Errorf("Announced initial player to player %d before Play() called", m.playerNum)
	}
}

func failIfNotAnnounced(t *testing.T, m *mockPlayer) {
	if !m.announcedInitialPlayer {
		t.Errorf("Did not Announce initial player to player %d after Play() called", m.playerNum)
	}
}
