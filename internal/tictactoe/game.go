package tictactoe

type Game struct {
	player [2]PlayerIO
}

type PlayerIO interface {
	NotifyGameStart()
	RequestMove() string
	ShareStateChange(stateChange string)
	ReportBadMoveSelection(code int, msg string)
}

func New() *Game {
	var g Game

	return &g
}

func (g *Game) InitializeGame(initialPlayer int) {
	g.player[0].NotifyGameStart()
	g.player[1].NotifyGameStart()
}

func (g *Game) Done() bool {
	return false
}

func (g *Game) HandleValidMoveFromPlayer(player int) {
	var otherPlayer int

	if player == 0 {
		otherPlayer = 1
	} else {
		otherPlayer = 0
	}

	g.player[player].RequestMove()
	g.player[0].ReportBadMoveSelection(0, "")
	g.player[otherPlayer].ShareStateChange("")
}

func (g *Game) SetPlayerIO(player int, io PlayerIO) {
	g.player[player] = io
}
