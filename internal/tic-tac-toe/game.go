package tic_tac_toe

type Game struct {
}

type PlayerIO interface {
}

func New() *Game {
	return nil
}

func (game *Game) InitializeGame(initialPlayer int) {
}

func (game *Game) Done() bool {
	return false
}

func (game *Game) HandleValidMoveFromPlayer(player int) {
}

func (game *Game) SetPlayerIO(player int, io PlayerIO) {
}
