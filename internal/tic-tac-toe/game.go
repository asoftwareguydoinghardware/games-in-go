package tic_tac_toe

type Game struct {
	player [2]PlayerIO
}

type PlayerIO interface {
	NotifyGameStart()
	RequestMove() string
}

func New() *Game {
	var g Game

	return &g
}

func (game *Game) InitializeGame(initialPlayer int) {
	game.player[0].NotifyGameStart()
	game.player[1].NotifyGameStart()
}

func (game *Game) Done() bool {
	return false
}

func (game *Game) HandleValidMoveFromPlayer(player int) {
	game.player[player].RequestMove()
}

func (game *Game) SetPlayerIO(player int, io PlayerIO) {
	game.player[player] = io
}
