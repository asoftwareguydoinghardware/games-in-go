package gameplay_test

type mockGame struct {
	initializeGameCount int
}

func (m *mockGame)InitializeGame(player int) {
	m.initializeGameCount++
}
