package tictactoe

import "fmt"
import "io"

type ownerId int

const (
	nobody ownerId = iota
	player0
	player1
)

type Game struct {
	players       [2]Player
	initialPlayer int
	squareOwner   [3][3]ownerId
}

type Player interface {
	AnnounceInitialPlayer(ip int)
	Move() string
}

func New() (game *Game) {
	var g Game

	return &g
}

func (g *Game) Done() bool {
	return false
}

func (g *Game) AssignPlayer(playerNum int, player Player) {
	g.players[playerNum] = player
}

func (g *Game) InitialPlayer(player int) {
	g.initialPlayer = player
}

func (g *Game) Play() {
	g.players[0].AnnounceInitialPlayer(g.initialPlayer)
	g.players[1].AnnounceInitialPlayer(g.initialPlayer)

	playerInd := g.initialPlayer

	for !g.done() {
		var owner ownerId

		move := g.players[playerInd].Move()
		if playerInd == 0 {
			owner = player0
		} else {
			owner = player1
		}
		for !g.isGoodMove(move) {
			move = g.players[playerInd].Move()
		}
		g.markSquareOccupied(move, owner)
		playerInd = (playerInd + 1) % 2
	}
}

func (g *Game) markSquareOccupied(move string, owner ownerId) {
	var row, col int

	fmt.Sscanf(move, "%d,%d", &row, &col)
	g.squareOwner[row][col] = owner
}

func (g *Game) isGoodMove(move string) bool {
	if move == "" || move == "x" {
		return false
	}

	var row, col int
	var junk byte
	n, err := fmt.Sscanf(move, "%d,%d%c", &row, &col, &junk)
	_, _ = n, err
	if err != nil && err != io.EOF {
		return false
	}
	if n == 3 {
		return false
	}

	if row >= 3 || row <= -1 {
		return false
	}
	if col >= 3 || col <= -1 {
		return false
	}

	if g.squareOwner[row][col] != nobody {
		return false
	}

	return true
}

func (g *Game) done() bool {
	owner := &g.squareOwner

	if player0 == owner[0][0] && player0 == owner[0][1] && player0 == owner[0][2] {
		return true
	}
	if player0 == owner[0][0] && player0 == owner[1][0] && player0 == owner[2][0] {
		return true
	}
	if player1 == owner[0][0] && player1 == owner[0][1] && player1 == owner[0][2] {
		return true
	}
	if player0 == owner[2][0] && player0 == owner[2][1] && player0 == owner[2][2] {
		return true
	}

	if player0 == owner[0][1] && player0 == owner[1][1] && player0 == owner[2][1] {
		return true
	}
	if player0 == owner[0][2] && player0 == owner[1][2] && player0 == owner[2][2] {
		return true
	}
	return false
}
