package tictactoe

import (
	"fmt"
	"io"
)

type owner int

const (
	player1 = owner(0)
	player2 = iota
	neither = iota
)

type Game struct {
	player       [2]PlayerIO
	lastError    int
	lastMsg      string
	moveNum      int
	squareOwners [9]owner
}

type PlayerIO interface {
	NotifyGameStart()
	RequestMove() string
	ShareStateChange(stateChange string)
	ReportBadMoveSelection(code int, msg string)
}

func New() *Game {
	var g Game

	for i := 0; i < len(g.squareOwners); i++ {
		g.squareOwners[i] = neither
	}

	return &g
}

func (g *Game) Initialize(initialPlayer int) {
	g.player[0].NotifyGameStart()
	g.player[1].NotifyGameStart()
}

func (g *Game) Done() bool {
	owners := g.squareOwners
	if owners[0] == player1 && owners[1] == player1 && owners[2] == player1 {
		return true
	}
	if owners[0] == player2 && owners[1] == player2 && owners[2] == player2 {
		return true
	}

	return false
}

func (g *Game) HandleValidMoveFromPlayer(player int) {
	var otherPlayer int

	if player == 0 {
		otherPlayer = 1
	} else {
		otherPlayer = 0
	}

	g.moveNum++
	move := g.player[player].RequestMove()
	for !g.isValidMove(move) {
		code := g.lastError
		msg := g.lastMsg
		g.player[player].ReportBadMoveSelection(code, msg)
		move = g.player[player].RequestMove()
	}
	square, _ := g.moveAsNumber(move)
	g.squareOwners[square] = owner(player)
	g.player[otherPlayer].ShareStateChange("")
}

func (g *Game) isValidMove(move string) (valid bool) {
	if square, err := g.moveAsNumber(move); err != nil {
		return false
	} else if g.squareOwners[square] != neither {
		g.lastError = 502
		g.lastMsg = "Bad move: square occupied"
		return false
	}

	return true
}

func (g *Game) moveAsNumber(move string) (num int, err error) {
	var junk byte

	const badIntMsg = "Invalid number %q"
	const badInt = 500
	const rangeErrorMsg = "Invalid move, must be in range 0-8"
	const rangeError = 501

	matched, err := fmt.Sscanf(move, "%v %c", &num, &junk)
	if matched != 1 {
		g.lastError = badInt
		g.lastMsg = fmt.Sprintf(badIntMsg, move)
		err = fmt.Errorf("%d", g.lastError)
		return -1, err
	}
	if err != nil && err != io.EOF {
		g.lastError = badInt
		g.lastMsg = fmt.Sprintf(badIntMsg, move)
		err = fmt.Errorf("%d", g.lastError)
		return -1, err
	}
	if num < 0 || num >= 9 {
		g.lastError = rangeError
		g.lastMsg = rangeErrorMsg
		err = fmt.Errorf("%d", g.lastError)
		return -1, err
	}

	return num, nil
}

func (g *Game) SetPlayerIO(player int, io PlayerIO) {
	g.player[player] = io
}
