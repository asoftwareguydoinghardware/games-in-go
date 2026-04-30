package gameplay_test

import (
	"github.com/ASoftwareGuyDoingHardware/games-in-go/internal/gameplay"
	"testing"
)

type mockGame struct {
}

func TestMainLoopExists(t *testing.T) {
	var m mockGame

	gameplay.MainLoop(m)
}
