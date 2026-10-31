package tictactoe_test

import (
	ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tictactoe"
	"slices"
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
	if have, want := announcements[0].move, defaultPlayer0Moves[0].move; have != want {
		t.Errorf("For player 0, Play method first announced move is %q, want %q", have, want)
	}

	if have, want := announcements[0].player, initialPlayer; have != want {
		t.Errorf("For player 0, Play method first announced move by player %d, want %d", have, want)
	}

	if len(p1.dataFromApp.announcedMoves) == 0 {
		t.Errorf("For player 1 Play method did not call AnnounceMove method.")
		return
	}
	announcements = p1.dataFromApp.announcedMoves
	if have, want := announcements[0].move, defaultPlayer0Moves[0].move; have != want {
		t.Errorf("For player 1, Play method first announced move is %q, want %q", have, want)
	}

	if have, want := announcements[0].player, initialPlayer; have != want {
		t.Errorf("For player 1, Play method first announced move by player %d, want %d", have, want)
	}
}

func TestAllMovesAnnouncedCorrectly(t *testing.T) {
	testAllMovesAnnouncedCorrectly(t, 0)
	/*
		Initially thought there would be a test:
			testAllMovesAnnouncedCorrectly(t, 1)
		but that passes.  TDD does not allow (or at least advocates
		against adding tests that pass).
	*/
}

func testAllMovesAnnouncedCorrectly(t *testing.T, initialPlayer int) {
	var secondPlayer int

	/*
		Player 0 will always win with the default moves; however, it takes
		an additonal move to win if player 1 goes first.

		The list of announced moves varies slightly depending on which player
		goes first.  If pXY indicates player X's move Y, then the sequence
		will either start p00, p10, p01, p11 or p10, p00, p11, p10.
	*/

	primaryMoves := defaultPlayer0Moves
	secondaryMoves := defaultPlayer1Moves
	totalMoves := 5

	if initialPlayer == 0 {
		secondPlayer = 1
	} else {
		secondPlayer = 0
		primaryMoves, secondaryMoves = secondaryMoves, primaryMoves
		totalMoves += 1
	}

	wantMoves := make([]announcedMove, totalMoves)
	primary := 0
	secondary := 0
	move := 0
	for {
		moveData := announcedMove{player: initialPlayer, move: primaryMoves[primary].move}
		primary++
		wantMoves[move] = moveData
		move++
		if move == totalMoves {
			break
		}

		moveData = announcedMove{player: secondPlayer, move: secondaryMoves[secondary].move}
		secondary++
		wantMoves[move] = moveData
		move++
		if move == totalMoves {
			break
		}
	}

	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	t.Logf("Initial player = %d, player0 moves = %v, player1 moves = %v", initialPlayer, p0.moves, p1.moves)

	g.InitialPlayer(initialPlayer)
	g.Play()

	if !slices.Equal(p0.dataFromApp.announcedMoves, wantMoves) {
		t.Errorf("*Player 0 have announced moves %v, want %v", p0.dataFromApp.announcedMoves, wantMoves)
	}

	if !slices.Equal(p1.dataFromApp.announcedMoves, wantMoves) {
		t.Errorf("*Player 1 have announced moves %v, want %v", p1.dataFromApp.announcedMoves, wantMoves)
	}
}

func TestAnnounceRankingsCalled(t *testing.T) {
	const initialPlayer = 1 // should not matter if this is 0 or 1

	g := ttt.New()
	p0, p1 := newMockPlayer(0), newMockPlayer(1)

	assignTracker(p0, p1)
	g.AssignPlayer(0, p0)
	g.AssignPlayer(1, p1)

	t.Logf("Initial player = %d, player0 moves = %v, player1 moves = %v", initialPlayer, p0.moves, p1.moves)

	g.InitialPlayer(initialPlayer)
	g.Play()

	if !p0.dataFromApp.announceRankingCalled {
		t.Errorf("AnnounceRanking method not called for player 0")
	}

	if !p1.dataFromApp.announceRankingCalled {
		t.Errorf("AnnounceRanking method not called for player 1")
	}
}
