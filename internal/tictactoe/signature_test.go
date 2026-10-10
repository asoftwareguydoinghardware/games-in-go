package tictactoe_test

import (
	ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tictactoe"
	"testing"
)

func TestNewExists(t *testing.T) {
	var game *ttt.Game

	game = ttt.New()
	_ = game
}

func TestDoneExists(t *testing.T) {
	var done bool

	g := ttt.New()
	done = g.Done()
	_ = done
}

func TestAssignPlayerExists(t *testing.T) {
	playerNum := 0
	player := newMockPlayer()
	g := ttt.New()

	g.AssignPlayer(playerNum, player)
}

func TestInitialPlayerExists(t *testing.T) {
	g := ttt.New()
	g.InitialPlayer(0)
}

func TestPlayExists(t *testing.T) {
	g := ttt.New()
	g.AssignPlayer(0, newMockPlayer)
	g.AssignPlayer(1, newMockPlayer)
	g.InitialPlayer(0)
	g.Play()
}
