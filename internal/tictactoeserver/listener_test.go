package tictactoeserver

import (
	"fmt"
	"net"
	"slices"
	"testing"
)

func TestRunGamesAcceptsAConnection(t *testing.T) {
	var listen mockListen
	conf := config{
		listener: &listen,
		fatal:    nullFatal,
	}

	runGames(conf)

	if !listen.acceptCalled {
		t.Errorf("runGames() did not call listen.Accept(conf)")
	}
}

func TestRunGamesAcceptsConnections(t *testing.T) {
	connectionCount := []int{2, 5, 1}

	for _, count := range connectionCount {
		listen := mockListen{numAcceptsRemaining: count}
		conf := config{
			listener: &listen,
			fatal:    nullFatal,
		}

		runGames(conf)

		if have, want := listen.acceptCallCount, count; have != want {
			t.Errorf("with mock listener, have %d calls to Accept(), want %d", have, want)
		}
	}
}

func TestRunGamesCallsFatalOnError(t *testing.T) {
	var count = 3
	var fatalArgs []any
	var fatalCallCount int

	listen := mockListen{numAcceptsRemaining: count}
	conf := config{listener: &listen,
		fatal: func(v ...any) {
			fatalArgs = v
			fatalCallCount++
		},
	}

	runGames(conf)

	if have, want := fatalCallCount, 1; have != want {
		t.Errorf("runGames() called fatal %d times want %d", have, want)
	}
	have, want := fatalArgs, []any{mockFunctionErr}
	if !slices.Equal(have, want) {
		t.Errorf("runGames() called fatal with args %v, want %v", have, want)
	}
}

type mockListen struct {
	acceptCalled        bool
	numAcceptsRemaining int
	acceptCallCount     int
}

var mockFunctionErr = fmt.Errorf("mock function")

func (m *mockListen) Accept() (c net.Conn, err error) {
	m.acceptCalled = true

	m.numAcceptsRemaining--
	m.acceptCallCount++

	if m.numAcceptsRemaining <= 0 {
		if m.numAcceptsRemaining < -10 {
			panic("excessive calls to mock Accept() after returning error condition")
		}
		return nil, mockFunctionErr
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

func nullFatal(v ...any) {
}
