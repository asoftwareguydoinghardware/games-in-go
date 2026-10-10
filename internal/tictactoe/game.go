package tictactoe

type Game struct {
	players [2]Player
}

type Player interface {
	AnnounceInitialPlayer(ip int)
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
}

func (g *Game) Play() {
	g.players[0].AnnounceInitialPlayer(1)
	g.players[1].AnnounceInitialPlayer(1)
}
