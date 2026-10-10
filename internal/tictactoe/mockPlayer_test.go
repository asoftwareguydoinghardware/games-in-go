package tictactoe_test

type mockPlayer struct {
}

func newMockPlayer() (newPlayer *mockPlayer) {
	var m mockPlayer

	return &m
}
