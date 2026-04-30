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

func TestDoneCalledAfterInitialize(t *testing.T) {
	var m mockGame
	m.doneReturnValues = []bool{true}

	gameplay.MainLoop(&m)

	if !m.calledDone {
		t.Errorf("Done() not called from MainLoop()")
	}
}

func TestGetValidMoveIsCalled(t *testing.T) {
	var m mockGame
	m.doneReturnValues = []bool{false}

	gameplay.MainLoop(&m)

	if !m.calledHandleValidMoveFromPlayer {
		t.Errorf("MainLoop() did not call HandleValidMoveFromPlayer()")
	}
}

func TestFirstMoveRequestedFromPlayerZero(t *testing.T) {
	var m mockGame
	m.doneReturnValues = []bool{false}
	m.havePlayerForMove = make([]int, 1)

	gameplay.MainLoop(&m)

	want := 0
	have := m.havePlayerForMove[0]

	if have != want {
		t.Errorf("Player %d was asked for the first move, want %d", have, want)
	}
}

func TestSecondMoveRequestedFromPlayerOne(t *testing.T) {
	var m mockGame
	m.doneReturnValues = []bool{false, false}
	m.havePlayerForMove = make([]int, 2)

	gameplay.MainLoop(&m)

	want := 1
	have := m.havePlayerForMove[1]

	if have != want {
		t.Errorf("Player %d was asked for the second move, want %d", have, want)
	}
}
