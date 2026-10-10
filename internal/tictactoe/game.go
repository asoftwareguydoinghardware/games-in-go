package tictactoe

type Game struct {
}

func New() (game *Game) {
	var g Game

	return &g
}

func (g *Game) Done() bool {
	return true
}
