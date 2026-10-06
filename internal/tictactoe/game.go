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
	for row := 0; row < 3; row++ {
		if g.rowOwned(row) {
			return true
		}
	}

	if g.columnOwned(0) {
		return true
	}
	if g.columnOwned(1) {
		return true
	}

	return false
}

func (g *Game) rowOwned(row int) bool {
	owners := g.squareOwners

	first := row * 3
	if owners[first] == neither {
		return false
	}

	if owners[first] == owners[first+1] && owners[first+1] == owners[first+2] {
		return true
	}

	return false
}

func (g *Game) columnOwned(column int) bool {
	owners := g.squareOwners

	switch column {
	case 0:
		if owners[0] == neither {
			return false
		}
		if owners[0] == owners[3] && owners[3] == owners[6] {
			return true
		}
	case 1:
		if owners[2] == neither {
			return false
		}
		if owners[1] == owners[4] && owners[4] == owners[7] {
			return true
		}
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
