package tictactoe_test

import (
	"context"
	ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tictactoe"
	"testing"
)

func TestNewExists(t *testing.T) {
	var game *ttt.Game

	game = ttt.New()
	_ = game
}

func TestAssignPlayerContextExists(t *testing.T) {
	playerNum := 0
	player := newMockPlayer(0)
	g := ttt.New()
	ctx := context.Background()

	g.AssignPlayer(ctx, playerNum, player)
}

func TestInitialPlayerExists(t *testing.T) {
	g := ttt.New()
	g.InitialPlayer(0)
}

func TestPlayContextExists(t *testing.T) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)
	assignTracker(p0, p1)
	g.AssignPlayer(context.Background(), 0, p0)
	g.AssignPlayer(context.Background(), 1, p1)
	g.InitialPlayer(0)
	g.PlayContext(context.Background())
}

func TestRankingExists(t *testing.T) {
	var results [][]int

	g := ttt.New()
	results = g.Ranking()
	_ = results
}
