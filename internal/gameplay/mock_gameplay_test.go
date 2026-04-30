package gameplay_test

type mockGame struct {
	initializeGameCount  int
	firstPlayerParameter int
}

func (m *mockGame) InitializeGame(player int) {
	m.initializeGameCount++
	if m.initializeGameCount == 1 {
		m.firstPlayerParameter = player
	}
}
