package tictactoe

import "fmt"
import "io"

type Game struct {
	players        [2]Player
	initialPlayer  int
	squareOccupied [3][3]bool
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

	player := g.initialPlayer

	occupied := &g.squareOccupied
	for !(occupied[0][0] && occupied[0][1] && occupied[0][2]) {
		move := g.players[player].Move()
		for !isGoodMove(move) {
			move = g.players[player].Move()
		}
		g.markSquareOccupied(move)
		player = (player + 1) % 2
	}
}

func (g *Game) markSquareOccupied(move string) {
	var row, col int

	fmt.Sscanf(move, "%d,%d", &row, &col)
	g.squareOccupied[row][col] = true
}

func isGoodMove(move string) bool {
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

	return true
}
