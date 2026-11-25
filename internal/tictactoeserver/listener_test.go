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

type mockListen struct {
	acceptCalled bool
}

func (m *mockListen) Accept() (c net.Conn, err error) {
	m.acceptCalled = true

	return nil, fmt.Errorf("mock function")
}

func (m *mockListen) Addr() (n net.Addr) {
	return n
}

func (m *mockListen) Close() (err error) {
	return err
}
