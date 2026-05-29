package tictactoe

import "unicode"

type Game struct {
	player [2]PlayerIO
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
	if !isValidMove(move) {
		g.player[player].ReportBadMoveSelection(0, "")
	}
	g.player[otherPlayer].ShareStateChange("")
}

func isValidMove(move string) (valid bool) {
	if len(move) != 1 {
		return false
	}
	if !unicode.IsDigit(rune(move[0])) {
		return false
	}
	if move[0] == '9' {
		return false
	}

	return true
}

func (g *Game) SetPlayerIO(player int, io PlayerIO) {
	g.player[player] = io
}
