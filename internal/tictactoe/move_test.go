package tictactoe_test

import (
	ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tictactoe"
	"testing"
)

type badMoveTestCase struct {
	player        int
	index         int
	move          string
	initialPlayer int
}

func TestBadMove(t *testing.T) {
	testcases := []badMoveTestCase{
		badMoveTestCase{player: 0, index: 0, move: "", initialPlayer: 0},
		badMoveTestCase{player: 0, index: 0, move: "x", initialPlayer: 0},
		badMoveTestCase{player: 0, index: 0, move: "4,1", initialPlayer: 0},
		badMoveTestCase{player: 0, index: 0, move: "5,1", initialPlayer: 0},
		badMoveTestCase{player: 0, index: 0, move: "3,1", initialPlayer: 0},
		badMoveTestCase{player: 0, index: 0, move: "-1,1", initialPlayer: 0},
		badMoveTestCase{player: 0, index: 0, move: "1 1", initialPlayer: 0},
		badMoveTestCase{player: 0, index: 0, move: "1,x", initialPlayer: 0},
		badMoveTestCase{player: 0, index: 0, move: "1,4", initialPlayer: 0},
		badMoveTestCase{player: 0, index: 0, move: "1,5", initialPlayer: 0},
		badMoveTestCase{player: 0, index: 0, move: "1,3", initialPlayer: 0},
		badMoveTestCase{player: 0, index: 0, move: "1,-2", initialPlayer: 0},
		badMoveTestCase{player: 0, index: 0, move: "1,-1", initialPlayer: 0},
		badMoveTestCase{player: 0, index: 0, move: "1,1x", initialPlayer: 0},
	}

	for _, tc := range testcases {
		testBadMove(t, tc)
	}
}

var defaultPlayer0Moves = moveList{
	{move: "0,0", good: true},
	{move: "0,1", good: true},
	{move: "0,2", good: true},
}

var defaultPlayer1Moves = moveList{
	{move: "1,0", good: true},
	{move: "2,0", good: true},
	{move: "2,1", good: true},
}

type move struct {
	move string
	good bool
}
type moveList []move

func testBadMove(t *testing.T, tc badMoveTestCase) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	p0List := defaultPlayer0Moves
	p1List := defaultPlayer1Moves
	if tc.player == 0 {
		p0List = addBadMoveToList(p0List, tc)
	} else {
		p1List = addBadMoveToList(p1List, tc)
	}

	p0.moves = p0List
	p1.moves = p1List

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	g.InitialPlayer(tc.initialPlayer)
	g.Play()

	if p0.tracker.firstBadMove != -1 {
		var firstList, secondList moveList

		if tc.player == 0 {
			firstList, secondList = p0List, p1List
		} else {
			firstList, secondList = p1List, p0List
		}

		t.Errorf("for move sequences %v and %v play did not properly alternate on move %d", firstList, secondList, p0.tracker.firstBadMove)
	}
}

func listFromArray(array []string) moveList {
	list := make(moveList, len(array))
	for i, v := range array {
		list[i].move = v
		list[i].good = true
	}

	return list
}

func addBadMoveToList(list moveList, badData badMoveTestCase) moveList {
	badEntry := move{move: badData.move, good: false}
	initial := list[:badData.index]
	ending := list[badData.index:]

	result := append(moveList{}, initial...)
	result = append(result, badEntry)
	result = append(result, ending...)

	return result
}

/*
tests to force correction to code based on tests that used one
based index values instead of zero based values
*/
func TestGoodMoves(t *testing.T) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)
	p0List := defaultPlayer0Moves
	p1List := defaultPlayer1Moves

	p0.moves = p0List
	p1.moves = p1List

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	g.InitialPlayer(0)
	g.Play()

	if p0.tracker.firstBadMove != -1 {
		t.Errorf("for move sequences %v and %v play did not properly alternate on move %d", p0List, p1List, p0.tracker.firstBadMove)
	}
}

type multipleBadMovesTestCase struct {
	player        int
	index         int
	move          string
	numBadMoves   int
	initialPlayer int
}

func TestMultipleBadMovesRepeatRequestForMove(t *testing.T) {
	testcases := []multipleBadMovesTestCase{
		{player: 0, index: 0, move: "x,y", numBadMoves: 2, initialPlayer: 0},
		{player: 0, index: 0, move: "x,y", numBadMoves: 3, initialPlayer: 0},
	}

	for _, tc := range testcases {
		testMultipleBadMovesRepeatRequestForMove(t, tc)
	}
}

func testMultipleBadMovesRepeatRequestForMove(t *testing.T, tc multipleBadMovesTestCase) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	p0List := defaultPlayer0Moves
	p1List := defaultPlayer1Moves
	if tc.player == 0 {
		p0List = addRepeatedBadMoveToList(p0List, tc)
	} else {
		p1List = addRepeatedBadMoveToList(p1List, tc)
	}

	p0.moves = p0List
	p1.moves = p1List

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	g.InitialPlayer(tc.initialPlayer)
	g.Play()

	if p0.tracker.firstBadMove != -1 {
		var firstList, secondList moveList

		if tc.player == 0 {
			firstList, secondList = p0List, p1List
		} else {
			firstList, secondList = p1List, p0List
		}

		t.Errorf("for move sequences %v and %v play did not properly alternate on move %d", firstList, secondList, p0.tracker.firstBadMove)
	}
}

func addRepeatedBadMoveToList(list moveList, badData multipleBadMovesTestCase) moveList {
	badEntry := move{move: badData.move, good: false}
	initial := list[:badData.index]
	ending := list[badData.index:]

	result := append(moveList{}, initial...)
	for i := 0; i < badData.numBadMoves; i++ {
		result = append(result, badEntry)
	}

	result = append(result, ending...)

	return result
}

func TestGameRunsUntilSimpleWin(t *testing.T) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	g.InitialPlayer(0)
	g.Play()

	if want := 5; p0.tracker.lastMove != want {
		t.Errorf("for move sequences %v and %v last move = %d, want %d", p0.moves, p1.moves, p0.tracker.lastMove, want)
	}
}

func TestGameRunsUntilSimpleWin2(t *testing.T) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	// If player 1 goes first, there will be six moves until
	// the game is over
	g.InitialPlayer(1)
	g.Play()

	if want := 6; p0.tracker.lastMove != want {
		t.Errorf("for move sequences %v and %v last move = %d, want %d", p0.moves, p1.moves, p0.tracker.lastMove, want)
	}
}

func TestGameRunsUntilSimpleWin3(t *testing.T) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	// it still takes 6 moves if player 0's last two moves are played
	// in reverse order
	p0.moves[2], p0.moves[1] = p0.moves[1], p0.moves[2]

	g.InitialPlayer(1)
	g.Play()

	if want := 6; p0.tracker.lastMove != want {
		t.Errorf("for move sequences %v and %v last move = %d, want %d", p0.moves, p1.moves, p0.tracker.lastMove, want)
	}
}

func TestGameRunsUntilThirdRowOwned(t *testing.T) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	p0.moves = moveList{
		{move: "2,0", good: true},
		{move: "2,1", good: true},
		{move: "2,2", good: true},
	}
	p1.moves = moveList{
		{move: "1,0", good: true},
		{move: "1,1", good: true},
	}

	g.InitialPlayer(0)
	g.Play()

	if want := 5; p0.tracker.lastMove != want {
		t.Errorf("for move sequences %v and %v last move = %d, want %d", p0.moves, p1.moves, p0.tracker.lastMove, want)
	}
}

func TestCannotClaimOpponentsSquare(t *testing.T) {
	const initialPlayer = 0

	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	p0.moves = moveList{
		{move: "0,0", good: true},
		{move: "0,1", good: true},
		{move: "0,2", good: true},
	}
	p1.moves = moveList{
		{move: "0,0", good: false}, // player 0 already claimed 0,0
		{move: "2,0", good: true},
		{move: "2,1", good: true},
		{move: "2,2", good: true},
	}

	p0List := p0.moves
	p1List := p1.moves

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	g.InitialPlayer(initialPlayer)
	g.Play()

	if p1.tracker.firstBadMove != -1 {
		t.Errorf("for move sequences %v and %v play did not properly alternate on move %d", p0List, p1List, p0.tracker.firstBadMove)
	}
}

func TestGameContinuesIfThirdRowNotOwnedBySamePlayer(t *testing.T) {
	testGameContinuesIfThirdRowNotOwnedBySamePlayer(t, 0)
	testGameContinuesIfThirdRowNotOwnedBySamePlayer(t, 1)
}

func testGameContinuesIfThirdRowNotOwnedBySamePlayer(t *testing.T, winningPlayer int) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	winningMoves := moveList{
		{move: "0,0", good: true},
		{move: "0,1", good: true},
		{move: "2,2", good: true},
		{move: "0,2", good: true},
	}
	losingMoves := moveList{
		{move: "2,0", good: true},
		{move: "2,1", good: true},
		{move: "1,2", good: true},
		{move: "1,0", good: true},
	}

	if winningPlayer == 0 {
		p0.moves = winningMoves
		p1.moves = losingMoves
	} else {
		p0.moves = losingMoves
		p1.moves = winningMoves
	}

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	g.InitialPlayer(winningPlayer)
	g.Play()

	if want := 7; p0.tracker.lastMove != want {
		t.Errorf("For move sequences %v and %v play ended after %d moves, want %d", p0.moves, p1.moves, p0.tracker.lastMove, want)
	}
}

func TestGameEndOnADraw(t *testing.T) {
	testGameEndOnADraw(t, 0)
	testGameEndOnADraw(t, 1)
}

func testGameEndOnADraw(t *testing.T, initialPlayer int) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	assignTracker(p0, p1)

	/* should correspond to this board config using a notation of
	Zn where Z is either x or o, and n is the move (from a game, not
	player perspective).
		x0 o1 x2
		o3 x4 o5
		o7 x6 o8
	*/
	moves := []string{
		"0,0", "0,1", "0,2", "1,0", "1,1", "1,2", "2,1", "2,0", "2,2",
	}

	if initialPlayer == 0 {
		assignGoodMoves(p0, p1, moves)
	} else {
		assignGoodMoves(p1, p0, moves)
	}

	g.InitialPlayer(initialPlayer)
	g.Play()

	if want := 9; p0.tracker.lastMove != want {
		t.Errorf("For move sequences %v and %v play ended after %d moves, want %d", p0.moves, p1.moves, p0.tracker.lastMove, want)
	}
}

func assignGoodMoves(firstPlayer *mockPlayer, secondPlayer *mockPlayer, moves []string) {
	list := [2]moveList{}

	i := 0

	for _, m := range moves {
		move := move{move: m, good: true}
		list[i] = append(list[i], move)
		i = (i + 1) % 2
	}

	firstPlayer.moves = list[0]
	secondPlayer.moves = list[1]
}
