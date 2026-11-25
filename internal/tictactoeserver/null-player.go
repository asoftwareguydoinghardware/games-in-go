package tictactoeserver

type nullPlayer struct {
	playerNumber int
}

func (n nullPlayer) AnnounceInitialPlayer(ip int) {
}

func (n nullPlayer) AnnounceMove(pn int, move string) {
}

func (n nullPlayer) AnnounceRanking(ranking [][]int) {
}

func (n nullPlayer) Move() string {
	return ""
}

func (n nullPlayer) AssignPlayerNumber(pn int) {
}
