package tictactoe_test

type mockPlayer struct {
	announcedInitialPlayer bool
	initialPlayer          int
	playerNum              int
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
