package tic_tac_toe_test

import (
	ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tic-tac-toe"
	"testing"
)

func TestInitialCallToDoneReturnsFalse(t *testing.T) {
	g := ttt.New()
	g.InitializeGame(0)
	have, want := g.Done(), false

	if have != want {
		t.Errorf("Initial call to Done = %v\n", have)
	}
}
