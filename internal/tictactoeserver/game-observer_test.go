package tictactoeserver

import ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tictactoe"

type assignPlayerArg struct {
	playerNum int
	player    ttt.Player
}
type gameObserver struct {
	assignPlayerArgs []assignPlayerArg

	initialPlayer []int
}

func (obs *gameObserver) AssignPlayer(pn int, player ttt.Player) {
	args := assignPlayerArg{pn, player}
	obs.assignPlayerArgs = append(obs.assignPlayerArgs, args)
}

func (obs *gameObserver) InitialPlayer(player int) {
	obs.initialPlayer = append(obs.initialPlayer, player)
}
