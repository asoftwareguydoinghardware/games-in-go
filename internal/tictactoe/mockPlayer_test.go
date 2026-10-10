package tictactoe_test

type mockPlayer struct {
	announcedInitialPlayer bool
	initialPlayer          int
	playerNum              int
	moveCalled             bool
}

var mockPlayerForFirstMove int
var mockPlayersAlternated = false
var mockLastPlayer int

func newMockPlayer(playerNum int) (newPlayer *mockPlayer) {
	var m mockPlayer

	mockPlayerForFirstMove = -1
	mockPlayersAlternated = false
	mockLastPlayer = -1

	m.playerNum = playerNum
	return &m
}

func (m *mockPlayer) AnnounceInitialPlayer(player int) {
	m.announcedInitialPlayer = true
	m.initialPlayer = player
}

func (m *mockPlayer) Move() {
	m.moveCalled = true
	if mockPlayerForFirstMove == -1 {
		mockPlayerForFirstMove = m.playerNum
		mockPlayersAlternated = true
		mockLastPlayer = m.playerNum
	} else if mockLastPlayer == m.playerNum {
		mockPlayersAlternated = false
	} else {
		mockLastPlayer = m.playerNum
	}
}
