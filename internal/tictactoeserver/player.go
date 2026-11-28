package tictactoeserver

import "context"

type altPlayer struct {
	playerNumber int
}

func (n altPlayer) AnnounceInitialPlayer(ip int) {
}

func (n altPlayer) AnnounceMove(pn int, move string) {
}

func (n altPlayer) AnnounceRanking(ranking [][]int) {
}

func (n altPlayer) Move() string {
	return ""
}

func (n altPlayer) AssignPlayerNumber(ctx context.Context, pn int) {
}
