package tictactoeserver

import ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tictactoe"
import "context"

type assignPlayerArg struct {
	playerNum int
	player    ttt.Player
}
type gameObserver struct {
	assignPlayerArgs []assignPlayerArg

	initialPlayer []int

	playCalled                         bool
	playCalledAfterInitialPlayerCalled bool
}

func (obs *gameObserver) AssignPlayer(ctx context.Context, pn int, player ttt.Player) {
	args := assignPlayerArg{pn, player}
	obs.assignPlayerArgs = append(obs.assignPlayerArgs, args)
}

func (obs *gameObserver) InitialPlayer(player int) {
	obs.initialPlayer = append(obs.initialPlayer, player)
}

func (obs *gameObserver) Play(ctx context.Context) {
	if obs.playCalled {
		// Currently not expecting Play to be called twice; however,
		// the logic of detecting order of calls could break if
		// the usage model needed to call Play() more than once.
		// Panic for now.
		panic("Play() called twice?")
	}
	obs.playCalled = true
	obs.playCalledAfterInitialPlayerCalled = len(obs.initialPlayer) > 0
}
