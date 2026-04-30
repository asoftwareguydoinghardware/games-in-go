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

func TestFirstCallToInitializeTestHasPlayerParamAsZero(t *testing.T) {
	var m mockGame

	gameplay.MainLoop(&m)

	want := 0
	if m.firstPlayerParameter != want {
		t.Errorf("First call to m.InitializalizeGame() initial call's argument is %d want %d", m.firstPlayerParameter, want)
	}
}
