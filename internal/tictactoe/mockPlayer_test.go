package tictactoe_test

type mockPlayer struct {
	announcedInitialPlayer bool
	initialPlayer          int
	playerNum              int
	moveCalled             bool
	tracker                *tracker
	iAmPlayer1             bool
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
}

func newMockPlayer(playerNum int) (newPlayer *mockPlayer) {
	var m mockPlayer

	m.playerNum = playerNum
	return &m
}

func (m *mockPlayer) AnnounceInitialPlayer(player int) {
	m.announcedInitialPlayer = true
	m.initialPlayer = player
}

func (m *mockPlayer) Move() {
	tracker := m.tracker
	m.moveCalled = true
	if tracker.state == wantInitialMove {
		if m.iAmPlayer1 {
			tracker.state = wantPlayer2Move
		} else {
			tracker.state = wantPlayer1Move
		}
		tracker.mockPlayerForFirstMove = m.playerNum
		tracker.mockPlayersAlternated = true
	} else if tracker.state == wantPlayer1Move {
		tracker.state = wantPlayer2Move
		if !m.iAmPlayer1 {
			tracker.mockPlayersAlternated = false
		}
	} else if tracker.state == wantPlayer2Move {
		tracker.state = wantPlayer1Move
		if m.iAmPlayer1 {
			tracker.mockPlayersAlternated = false
		}
	}
}

func assignTracker(m1 *mockPlayer, m2 *mockPlayer) {
	var t tracker

	t.state = wantInitialMove
	t.mockPlayerForFirstMove = -1
	t.mockPlayersAlternated = false

	m1.tracker = &t
	m2.tracker = &t

	m1.iAmPlayer1 = true
	m2.iAmPlayer1 = false
}
