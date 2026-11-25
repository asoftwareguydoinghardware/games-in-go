package tictactoeserver

import (
	"fmt"
	"net"
	"testing"
)

func TestRunGamesAcceptsAConnection(t *testing.T) {
	var listen mockListen
	conf := config{listener: &listen}

	runGames(conf)

	if !listen.acceptCalled {
		t.Errorf("runGames() did not call listen.Accept(conf)")
	}
}

func TestRunGamesAcceptsConnections(t *testing.T) {
	connectionCount := []int{2, 5}

	for _, count := range connectionCount {
		listen := mockListen{numAcceptsRemaining: count}
		conf := config{listener: &listen}

		runGames(conf)

		if have, want := listen.acceptCallCount, count; have != want {
			t.Errorf("with mock listener, have %d calls to Accept(), want %d", have, want)
		}
	}
}

type mockListen struct {
	acceptCalled        bool
	numAcceptsRemaining int
	acceptCallCount     int
}

func (m *mockListen) Accept() (c net.Conn, err error) {
	m.acceptCalled = true

	m.numAcceptsRemaining--
	m.acceptCallCount++

	if m.numAcceptsRemaining <= 0 {
		if m.numAcceptsRemaining < -10 {
			panic("excessive calls to mock Accept() after returning error condition")
		}
		return nil, fmt.Errorf("mock function")
	} else {
		return nil, nil
	}
}

func (m *mockListen) Addr() (n net.Addr) {
	return n
}

func (m *mockListen) Close() (err error) {
	return err
}
