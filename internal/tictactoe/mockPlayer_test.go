package tictactoe_test

type mockPlayer struct {
	announcedInitialPlayer bool
	initialPlayer          int
	playerNum              int
	moveCalled             bool
	tracker                *tracker
}

type trackerState int

const (
	wantInitialMove trackerState = iota
	otherState
)

type tracker struct {
	state                  trackerState
	mockPlayerForFirstMove int
	mockPlayersAlternated  bool
	mockLastPlayer         int
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
		tracker.state = otherState
		tracker.mockPlayerForFirstMove = m.playerNum
		tracker.mockPlayersAlternated = true
		tracker.mockLastPlayer = m.playerNum
	} else if tracker.mockLastPlayer == m.playerNum {
		tracker.mockPlayersAlternated = false
	} else {
		tracker.mockLastPlayer = m.playerNum
	}
}

func assignTracker(m1 *mockPlayer, m2 *mockPlayer) {
	var t tracker

	t.state = wantInitialMove
	t.mockPlayerForFirstMove = -1
	t.mockPlayersAlternated = false
	t.mockLastPlayer = -1

	m1.tracker = &t
	m2.tracker = &t
}
