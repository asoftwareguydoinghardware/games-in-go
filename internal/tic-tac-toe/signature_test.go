package tic_tac_toe_test

import (
        ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tic-tac-toe"    
        "testing"  
)

func TestNewExists(t *testing.T) {
	var game *ttt.Game

	game = ttt.New()
	_ = game
}
