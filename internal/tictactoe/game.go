package tictactoe

type Game struct {
	players       [2]Player
	initialPlayer int
}

type Player interface {
	AnnounceInitialPlayer(ip int)
}

func New() (game *Game) {
	var g Game

	g.initialPlayer = 1
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
}
