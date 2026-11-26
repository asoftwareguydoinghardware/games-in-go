package tictactoeserver

import ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tictactoe"
import "context"

type nullGame struct{}

func (*nullGame) AssignPlayer(int, ttt.Player) {
}

func (*nullGame) InitialPlayer(int) {
}

func (*nullGame) AssignPlayerContext(context.Context, int, ttt.Player) {
}
