package tic_tac_toe

type Game struct {
	player [2]PlayerIO
}

type PlayerIO interface {
	NotifyGameStart()
	RequestMove() string
	ShareStateChange(stateChange string)
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
	var otherPlayer int

	if player == 0 {
		otherPlayer = 1
	} else {
		otherPlayer = 0
	}

	game.player[player].RequestMove()
	game.player[otherPlayer].ShareStateChange("")
}

func (game *Game) SetPlayerIO(player int, io PlayerIO) {
	game.player[player] = io
}
