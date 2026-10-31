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
	AnnounceMove(player int, move string)
	AnnounceRanking(ranking [][]int)
	Move() string
	AssignPlayerNumber(pn int)
}

func New() (game *Game) {
	var g Game

	return &g
}

func (g *Game) AssignPlayer(playerNum int, player Player) {
	g.players[playerNum] = player
	player.AssignPlayerNumber(playerNum)
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

		g.players[0].AnnounceMove(playerInd, move)
		g.players[1].AnnounceMove(playerInd, move)

		playerInd = (playerInd + 1) % 2
	}

	var ranking [][]int
	g.players[0].AnnounceRanking(ranking)
}

/*
To allow the interface to handle things where players essentially
accumulate points, Ranking() returns a slice where the first element
of the slice is a slice of players that have the highest points.
Normally the slice of players that have the highest point will contain
a single entry as players are not often tied.  The same rationale/argument
applies to the subsequent members of the slice returned.

If, for example, there are four players and there is a tie for second place,
the top level slice will have three elements:

	[0] a slice containing the player that is in first place
	[1] a slice containing the two players in second place
	[2] a slice containing the player that came in last place.

There should be no entries in the slice with a nil value.
In a game like tic-tac-toe that can result in a draw, that condition should
be indicated by having a slice with a single entry which contains the
identity of each player.
*/
func (g *Game) Ranking() (ranks [][]int) {
	owner := &g.squareOwner

	if player0Won(owner) {
		return [][]int{
			{0},
			{1},
		}
	}
	if player1Won(owner) {
		return [][]int{
			{1},
			{0},
		}
	}

	return [][]int{
		{0, 1},
	}
}

func (g *Game) markSquareOccupied(move string, owner ownerId) {
	var row, col int

	fmt.Sscanf(move, "%d,%d", &row, &col)
	g.squareOwner[row][col] = owner
}

func (g *Game) isGoodMove(move string) bool {
	var row, col int
	var junk byte
	n, err := fmt.Sscanf(move, "%d,%d%c", &row, &col, &junk)
	_, _ = n, err
	if err != nil && err != io.EOF {
		return false
	}
	if n != 2 {
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

	if player0Won(owner) {
		return true
	}
	if player1Won(owner) {
		return true
	}

	if noMovesLeft(owner) {
		return true
	}

	return false
}

func playerOwnsADiagonal(owner *[3][3]ownerId, player ownerId) bool {
	if player == owner[0][0] && player == owner[1][1] && player == owner[2][2] {
		return true
	}
	if player == owner[0][2] && player == owner[1][1] && player == owner[2][1] {
		return true
	}

	return false
}

func playerOwnsARow(owner *[3][3]ownerId, player ownerId) bool {
	for row := 0; row < 3; row++ {
		if player == owner[row][0] && player == owner[row][1] && player == owner[row][2] {
			return true
		}
	}
	return false
}

func playerOwnsAColumn(owner *[3][3]ownerId, player ownerId) bool {
	for col := 0; col < 3; col++ {
		if player == owner[0][col] && player == owner[1][col] && player == owner[2][col] {
			return true
		}
	}

	return false

}

func noMovesLeft(owner *[3][3]ownerId) bool {
	for _, row := range *owner {
		for _, square := range row {
			if square == nobody {
				return false
			}
		}
	}

	return true
}

func player0Won(owner *[3][3]ownerId) bool {
	if playerOwnsAColumn(owner, player0) {
		return true
	}
	if playerOwnsARow(owner, player0) {
		return true
	}
	if playerOwnsADiagonal(owner, player0) {
		return true
	}

	return false
}

func player1Won(owner *[3][3]ownerId) bool {
	if playerOwnsAColumn(owner, player1) {
		return true
	}
	if playerOwnsARow(owner, player1) {
		return true
	}
	if playerOwnsADiagonal(owner, player1) {
		return true
	}

	return false
}
