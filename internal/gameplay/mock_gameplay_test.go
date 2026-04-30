package gameplay_test

type mockGame struct {
	initializeGameCount  int
	firstPlayerParameter int

	calledDone       bool
	doneReturnValues []bool
	doneCount        int

	calledHandleValidMoveFromPlayer bool
	havePlayerForMove               []int
	playerForMoveCount              int
	overflowedHavePlayerForMove     bool
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

func (m *mockGame) HandleValidMoveFromPlayer(player int) {
	m.calledHandleValidMoveFromPlayer = true

	i := m.playerForMoveCount
	m.playerForMoveCount++

	if i < len(m.havePlayerForMove) {
		m.havePlayerForMove[i] = player
	} else {
		m.overflowedHavePlayerForMove = true
	}
}
