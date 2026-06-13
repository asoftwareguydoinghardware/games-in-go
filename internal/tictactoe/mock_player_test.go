package tictactoe_test

import "fmt"

const trackMoves = false

type msgPair struct {
	moveNum int
	code    int
	msg     string
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
	if trackMoves {
		fmt.Printf("requested move number %d ", m.currentMove)
		defer func() { fmt.Printf("returning move %q\n", move) }()
	}

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
	tuple := msgPair{m.currentMove - 1, code, msg}
	m.badMoveMsgs = append(m.badMoveMsgs, tuple)
}
