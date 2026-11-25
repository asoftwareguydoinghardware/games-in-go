package tictactoeserver

import ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tictactoe"

type nullGame struct{}

func (*nullGame) AssignPlayer(int, ttt.Player) {
}
