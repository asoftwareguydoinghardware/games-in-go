package tictactoe_test

import (
	"fmt"
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

func TestSimpleColumnWin(t *testing.T) {
	testSimpleColumnWin(t, 0, 0, 0)
	testSimpleColumnWin(t, 0, 0, 1)
	testSimpleColumnWin(t, 0, 0, 2)
	testSimpleColumnWin(t, 0, 1, 0)
	testSimpleColumnWin(t, 0, 1, 1)
	testSimpleColumnWin(t, 0, 1, 2)
}

func testSimpleColumnWin(t *testing.T, initialPlayer, winningPlayer int, column int) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	assignTracker(p0, p1)

	winningMoves := moveList{}
	losingMoves := moveList{}

	for m := 0; m < 3; m++ {
		wMove := fmt.Sprintf("%d,%d", m, column)
		winningMoves = append(winningMoves, move{move: wMove, good: true})
		// lMove will pick the same row as the winning move, but will
		// alternate choosing a column left/right of the winning move
		var offset int
		if m%2 == 0 {
			offset = 1
		} else {
			offset = 2
		}
		lMove := fmt.Sprintf("%d,%d", m, (column+offset)%3)
		losingMoves = append(losingMoves, move{move: lMove, good: true})
	}
	/* If the winner goes first, they will win on their third move and
	their opponent will only be able to play two moves.

	If the loser goes first, they will play three moves, each of which
	is follwed by a move by the winner.

	In each case, having three moves available for the loser is
	enough to ensure we reach an end game.
	*/

	if winningPlayer == 0 {
		p0.moves = winningMoves
		p1.moves = losingMoves
	} else {
		p0.moves = losingMoves
		p1.moves = winningMoves
	}

	t.Logf("winning column %d, winning player %d, initial player %d, player 0 moves %v, player1 moves %v\n", column, winningPlayer, initialPlayer, p0.moves, p1.moves)

	if initialPlayer == 0 {
		showBoardMoves(t, p0, p1)
	} else {
		showBoardMoves(t, p1, p0)
	}

	g.InitialPlayer(initialPlayer)
	g.Play()

	var want int
	if winningPlayer == initialPlayer {
		want = 5
	} else {
		want = 6
	}

	if p0.tracker.lastMove != want {
		t.Errorf("For move sequences %v and %v play ended after %d moves, want %d", p0.moves, p1.moves, p0.tracker.lastMove, want)
	}
}

func showBoardMoves(t *testing.T, firstPlayer, secondPlayer *mockPlayer) {
	type square struct {
		owner   string
		moveNum int
	}
	squares := [3][3]square{}

	for i, m := range firstPlayer.moves {
		row, col := rowColumnFromMove(m.move)
		s := &squares[row][col]
		s.owner = "x"
		s.moveNum = i*2 + 1
	}

	for i, m := range secondPlayer.moves {
		row, col := rowColumnFromMove(m.move)
		s := &squares[row][col]
		s.owner = "o"
		s.moveNum = i*2 + 2
	}

	config := ""
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			s := &squares[r][c]
			if s.owner == "" {
				config += "-- "
			} else {
				config += fmt.Sprintf("%s%d ", s.owner, s.moveNum)
			}
		}
		config += "\n"
	}
	t.Log("Board configuration for X going first:\n" + config)
}

func rowColumnFromMove(m string) (row, col int) {
	n, err := fmt.Sscanf(m, "%d,%d", &row, &col)
	if n != 2 || err != nil {
		msg := fmt.Sprintf("bad move string %q", m)
		panic(msg)
	}

	return row, col
}

func TestSimpleRowWins(t *testing.T) {
	testSimpleRowWin(t, 0, 0, 1)
	testSimpleRowWin(t, 0, 1, 1)
	testSimpleRowWin(t, 0, 1, 2)
}

func testSimpleRowWin(t *testing.T, initialPlayer, winningPlayer int, row int) {
	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	assignTracker(p0, p1)

	winningMoves := moveList{}
	losingMoves := moveList{}

	for m := 0; m < 3; m++ {
		wMove := fmt.Sprintf("%d,%d", row, m)
		winningMoves = append(winningMoves, move{move: wMove, good: true})
		// lMove will pick the same column as the winning move, but will
		// alternate choosing a row up/down of the winning move
		var offset int
		if m%2 == 0 {
			offset = 1
		} else {
			offset = 2
		}
		lMove := fmt.Sprintf("%d,%d", (row+offset)%3, m)
		losingMoves = append(losingMoves, move{move: lMove, good: true})
	}

	if winningPlayer == 0 {
		p0.moves = winningMoves
		p1.moves = losingMoves
	} else {
		p0.moves = losingMoves
		p1.moves = winningMoves
	}

	t.Logf("winning row %d, winning player %d, initial player %d, player 0 moves %v, player1 moves %v\n", row, winningPlayer, initialPlayer, p0.moves, p1.moves)

	if initialPlayer == 0 {
		showBoardMoves(t, p0, p1)
	} else {
		showBoardMoves(t, p1, p0)
	}

	g.InitialPlayer(initialPlayer)
	g.Play()

	var want int
	if winningPlayer == initialPlayer {
		want = 5
	} else {
		want = 6
	}

	if p0.tracker.lastMove != want {
		t.Errorf("For move sequences %v and %v play ended after %d moves, want %d", p0.moves, p1.moves, p0.tracker.lastMove, want)
	}
}
