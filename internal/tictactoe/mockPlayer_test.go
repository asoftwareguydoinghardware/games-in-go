package tictactoe_test

import "fmt"

type appData struct {
	assignedPlayerNum bool
	playerNum         int
}

type mockPlayer struct {
	dataFromApp appData

	announcedInitialPlayer bool
	initialPlayer          int

	announceMoveCalled bool

	mockPlayerNum     int
	moveCalled        bool
	tracker           *tracker
	iAmPlayer1        bool
	moves             moveList
	currentPlayerMove int
}

type trackerState int

const (
	wantInitialMove trackerState = iota
	wantPlayer1Move
	wantPlayer2Move
	otherState
)

type tracker struct {
	state                  trackerState
	mockPlayerForFirstMove int
	mockPlayersAlternated  bool
	currentTrackerMove     int
	firstBadMove           int
	lastMove               int
}

func newMockPlayer(playerNum int) (newPlayer *mockPlayer) {
	var m mockPlayer

	m.mockPlayerNum = playerNum
	if playerNum == 0 {
		m.moves = defaultPlayer0Moves
	} else {
		m.moves = defaultPlayer1Moves
	}
	return &m
}

func (m *mockPlayer) AssignPlayerNumber(pn int) {
	m.dataFromApp.assignedPlayerNum = true
	m.dataFromApp.playerNum = pn
}

func (m *mockPlayer) AnnounceInitialPlayer(player int) {
	m.announcedInitialPlayer = true
	m.initialPlayer = player
}

func (m *mockPlayer) AnnounceMove(player int, move string) {
	m.announceMoveCalled = true
}

func (m *mockPlayer) Move() string {
	tracker := m.tracker
	m.moveCalled = true

	switch tracker.state {
	case wantInitialMove:
		if m.iAmPlayer1 {
			tracker.state = wantPlayer2Move
		} else {
			tracker.state = wantPlayer1Move
		}
		tracker.mockPlayerForFirstMove = m.mockPlayerNum
		tracker.mockPlayersAlternated = true
	case wantPlayer1Move:
		tracker.state = wantPlayer2Move
		if !m.iAmPlayer1 {
			tracker.mockPlayersAlternated = false
			if tracker.firstBadMove == -1 {
				tracker.firstBadMove = tracker.currentTrackerMove
			}
		}
	case wantPlayer2Move:
		tracker.state = wantPlayer1Move
		if m.iAmPlayer1 {
			tracker.mockPlayersAlternated = false
			if tracker.firstBadMove == -1 {
				tracker.firstBadMove = tracker.currentTrackerMove
			}
		}
	}

	if m.currentPlayerMove >= len(m.moves) {
		msg := fmt.Sprintf("internal error, want to return move %d, but last move index is %d, player %d, moves %v", m.currentPlayerMove, len(m.moves)-1, m.mockPlayerNum, m.moves)
		panic(msg)
	}

	move := m.moves[m.currentPlayerMove]
	m.currentPlayerMove++
	if !move.good {
		if m.iAmPlayer1 {
			tracker.state = wantPlayer1Move
		} else {
			tracker.state = wantPlayer2Move
		}
	}

	tracker.currentTrackerMove++
	tracker.lastMove = tracker.currentTrackerMove

	return move.move
}

func assignTracker(m1 *mockPlayer, m2 *mockPlayer) {
	var t tracker

	t.state = wantInitialMove
	t.mockPlayerForFirstMove = -1
	t.mockPlayersAlternated = false
	t.firstBadMove = -1
	t.lastMove = 0

	m1.tracker = &t
	m2.tracker = &t

	m1.iAmPlayer1 = true
	m2.iAmPlayer1 = false
}
