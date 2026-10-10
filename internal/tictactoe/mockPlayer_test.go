package tictactoe_test

type mockPlayer struct {
	announcedInitialPlayer bool
	initialPlayer          int
	playerNum              int
	moveCalled             bool
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
	m.moveCalled = true
}
