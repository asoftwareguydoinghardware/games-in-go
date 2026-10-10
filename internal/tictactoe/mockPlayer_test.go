package tictactoe_test

type mockPlayer struct {
	announcedInitialPlayer bool
}

func newMockPlayer() (newPlayer *mockPlayer) {
	var m mockPlayer

	return &m
}

func (m *mockPlayer) AnnounceInitialPlayer(player int) {
	m.announcedInitialPlayer = true
}
