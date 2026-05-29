package tic_tac_toe_test

type msgPair struct {
	code int
	msg  string
}
type mockPlayer struct {
	notifiedOfGameStart bool

	moves       []string
	currentMove int

	sharedStateChange bool

	badMoveMsgs []msgPair
}

func newMockPlayerIO() *mockPlayer {
	var m mockPlayer

	return &m
}

func (m *mockPlayer) NotifyGameStart() {
	m.notifiedOfGameStart = true
}

func (m *mockPlayer) RequestMove() (move string) {
	if m.currentMove == len(m.moves) {
		return ""
	}

	move = m.moves[m.currentMove]
	m.currentMove++
	return move
}

func (m *mockPlayer) ShareStateChange(state string) {
	m.sharedStateChange = true
}

func (m *mockPlayer) ReportBadMoveSelection(code int, msg string) {
	pair := msgPair{code, msg}
	m.badMoveMsgs = append(m.badMoveMsgs, pair)
}
