package tictactoe

type Game struct {
}

type Player interface {
}

func New() (game *Game) {
	var g Game

	return &g
}

func (g *Game) Done() bool {
	return false
}

func (g *Game) AssignPlayer(playerNum int, player Player) {
}

func (g *Game) InitialPlayer(player int) {
}

func (g *Game) Play() {
}
