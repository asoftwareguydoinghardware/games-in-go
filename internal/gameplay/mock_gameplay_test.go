package gameplay_test

type mockGame struct {
	initializeGameCount  int
	firstPlayerParameter int

	calledDone       bool
	doneReturnValues []bool
	doneCount        int
}

func (m *mockGame) InitializeGame(player int) {
	m.initializeGameCount++
	if m.initializeGameCount == 1 {
		m.firstPlayerParameter = player
	}
}

func (m *mockGame) Done() bool {
	m.calledDone = true
	i := m.doneCount
	m.doneCount++

	if i >= len(m.doneReturnValues) {
		return true
	}
	return m.doneReturnValues[i]
}
