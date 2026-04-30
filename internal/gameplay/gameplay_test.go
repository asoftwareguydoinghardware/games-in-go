package gameplay_test

import (
	"github.com/ASoftwareGuyDoingHardware/games-in-go/internal/gameplay"
	"testing"
)

func TestMainLoopCallsInitializeGame(t *testing.T) {
	var m mockGame

	gameplay.MainLoop(&m)

	if m.initializeGameCount != 1 {
		t.Errorf("MainLoop(m) did not call m.InitializalizeGame")
	}
}
