package tictactoe

import (
	"fmt"
	"io"
)

type Game struct {
	player    [2]PlayerIO
	lastError int
	lastMsg   string
}

type PlayerIO interface {
	NotifyGameStart()
	RequestMove() string
	ShareStateChange(stateChange string)
	ReportBadMoveSelection(code int, msg string)
}

func New() *Game {
	var g Game

	return &g
}

func (g *Game) InitializeGame(initialPlayer int) {
	g.player[0].NotifyGameStart()
	g.player[1].NotifyGameStart()
}

func (g *Game) Done() bool {
	return false
}

func (g *Game) HandleValidMoveFromPlayer(player int) {
	var otherPlayer int

	if player == 0 {
		otherPlayer = 1
	} else {
		otherPlayer = 0
	}

	move := g.player[player].RequestMove()
	for !g.isValidMove(move) {
		code := g.lastError
		msg := g.lastMsg
		g.player[player].ReportBadMoveSelection(code, msg)
		move = g.player[player].RequestMove()
	}
	g.player[otherPlayer].ShareStateChange("")
}

func (g *Game) isValidMove(move string) (valid bool) {
	var num int
	var junk byte

	const badIntMsg = "Invalid number %q"
	const badInt = 500
	const rangeErrorMsg = "Invalid move, must be in range 0-8"
	const rangeError = 501

	matched, err := fmt.Sscanf(move, "%v %c", &num, &junk)
	if matched != 1 {
		g.lastError = badInt
		g.lastMsg = fmt.Sprintf(badIntMsg, move)
		return false
	}
	if err != nil && err != io.EOF {
		return false
	}
	if num < 0 {
		g.lastError = rangeError
		g.lastMsg = rangeErrorMsg
		return false
	}
	if num >= 9 {
		return false
	}

	return true
}

func (g *Game) SetPlayerIO(player int, io PlayerIO) {
	g.player[player] = io
}
