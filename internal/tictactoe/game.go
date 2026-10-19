package tictactoe

type Game struct {
	players       [2]Player
	initialPlayer int
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

	move := g.players[player].Move()
	if move == "" || move == "x" {
		g.players[player].Move()
	}
	player = (player + 1) % 2
	g.players[player].Move()
}
