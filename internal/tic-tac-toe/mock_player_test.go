package tic_tac_toe_test

type mockPlayer struct {
	notifiedOfGameStart bool
}

func newMockPlayerIO() *mockPlayer {
	var m mockPlayer

	return &m
}

func (m *mockPlayer) NotifyGameStart() {
	m.notifiedOfGameStart = true
}
