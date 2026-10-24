package tictactoe_test

import (
	ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tictactoe"
	"slices"
	"testing"
)

func TestRankingReturnsCorrectValueForADraw(t *testing.T) {
	const initialPlayer = 0

	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	assignTracker(p0, p1)

	/*
		Final board config showing players (x,o) and move number
			x0 x8 o1
			o5 x2 x4
			x6 o7 o3
	*/
	initialMoves := moveList{
		{move: "0,0", good: true},
		{move: "1,1", good: true},
		{move: "1,2", good: true},
		{move: "2,0", good: true},
		{move: "0,1", good: true},
	}
	secondaryMoves := moveList{
		{move: "0,2", good: true},
		{move: "2,2", good: true},
		{move: "1,0", good: true},
		{move: "2,1", good: true},
	}

	p0.moves = initialMoves
	p1.moves = secondaryMoves

	t.Logf("game is a draw initial player %d, player 0 moves %v, player1 moves %v\n", initialPlayer, p0.moves, p1.moves)
	if initialPlayer == 0 {
		showBoardMoves(t, p0, p1)
	} else {
		showBoardMoves(t, p1, p0)
	}

	g.InitialPlayer(initialPlayer)
	g.Play()

	results := g.Ranking()

	if have, want := len(results), 1; have != want {
		t.Errorf("In game with initial move by player %d, moves for player 0 %v, and moves for player 1 = %v have len(g.Ranking()) = %d, want %d", initialPlayer, p0.moves, p1.moves, have, want)
		return
	}

	if have, want := slices.Sorted(slices.Values(results[0])), []int{0, 1}; slices.Compare(have, want) != 0 {
		t.Errorf("In game with initial move by player %d, moves for player 0 %v, and moves for player 1 = %v have g.Ranking()[0] = %v, want %v", initialPlayer, p0.moves, p1.moves, have, want)
		return

	}
}

func TestRankingReturnsCorrectValueWhenPlayer0OwnsAColumn(t *testing.T) {
	const initialPlayer = 0

	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	assignTracker(p0, p1)

	/*
		Final board config showing players (x,o) and move number
			x0 x2 x4
			o1 o3 --
			-- -- --
	*/
	initialMoves := moveList{
		{move: "0,0", good: true},
		{move: "0,1", good: true},
		{move: "0,2", good: true},
	}
	secondaryMoves := moveList{
		{move: "1,0", good: true},
		{move: "1,1", good: true},
	}

	p0.moves = initialMoves
	p1.moves = secondaryMoves

	t.Logf("Player 0 wins with column 0 initial player %d, player 0 moves %v, player1 moves %v\n", initialPlayer, p0.moves, p1.moves)
	if initialPlayer == 0 {
		showBoardMoves(t, p0, p1)
	} else {
		showBoardMoves(t, p1, p0)
	}

	g.InitialPlayer(initialPlayer)
	g.Play()

	results := g.Ranking()

	if have, want := len(results), 2; have != want {
		t.Errorf("In game with initial move by player %d, moves for player 0 %v, and moves for player 1 = %v have len(g.Ranking()) = %d, want %d", initialPlayer, p0.moves, p1.moves, have, want)
		return
	}

	if have, want := results[1][0], 1; have != want {
		t.Errorf("In game with initial move by player %d, moves for player 0 %v, and moves for player 1 = %v have g.Ranking()[1][0]) = %d, want %d", initialPlayer, p0.moves, p1.moves, have, want)
	}
}

func TestRankingReturnsCorrectValueWhenPlayer0OwnsARow(t *testing.T) {
	const initialPlayer = 0

	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	assignTracker(p0, p1)

	/*
		Final board config showing players (x,o) and move number
			x0 o1 --
			x2 o3 --
			x3 -- --
	*/
	initialMoves := moveList{
		{move: "0,0", good: true},
		{move: "1,0", good: true},
		{move: "2,0", good: true},
	}
	secondaryMoves := moveList{
		{move: "0,1", good: true},
		{move: "1,1", good: true},
	}

	p0.moves = initialMoves
	p1.moves = secondaryMoves

	t.Logf("Player 0 wins with column 0 initial player %d, player 0 moves %v, player1 moves %v\n", initialPlayer, p0.moves, p1.moves)
	if initialPlayer == 0 {
		showBoardMoves(t, p0, p1)
	} else {
		showBoardMoves(t, p1, p0)
	}

	g.InitialPlayer(initialPlayer)
	g.Play()

	results := g.Ranking()

	if have, want := len(results), 2; have != want {
		t.Errorf("In game with initial move by player %d, moves for player 0 %v, and moves for player 1 = %v have len(g.Ranking()) = %d, want %d", initialPlayer, p0.moves, p1.moves, have, want)
		return
	}

	if have, want := results[1][0], 1; have != want {
		t.Errorf("In game with initial move by player %d, moves for player 0 %v, and moves for player 1 = %v have g.Ranking()[1][0]) = %d, want %d", initialPlayer, p0.moves, p1.moves, have, want)
	}
}
