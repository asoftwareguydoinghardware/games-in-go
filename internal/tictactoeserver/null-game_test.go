package tictactoeserver

import ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tictactoe"
import "context"

type nullGame struct{}

func (*nullGame) InitialPlayer(int) {
}

func (*nullGame) AssignPlayer(context.Context, int, ttt.Player) {
}
