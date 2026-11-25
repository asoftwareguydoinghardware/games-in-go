package tictactoeserver

import (
	"context"
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

	runGames(context.Background(), conf)

	if !listen.acceptCalled {
		t.Errorf("runGames() did not call listen.Accept(conf)")
	}
}

func TestRunGamesAcceptsConnections(t *testing.T) {
	connectionCount := []int{2 /*, 5, 1 */}

	for _, count := range connectionCount {
		listen := mockListen{numAcceptsRemaining: count}
		conf := config{
			listener: &listen,
			fatal:    nullFatal,
		}

		runGames(context.Background(), conf)

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

	runGamesOld(conf)

	if have, want := fatalCallCount, 1; have != want {
		t.Errorf("runGamesOld() called fatal %d times want %d", have, want)
	}
	have, want := fatalArgs, []any{mockFunctionErr}
	if !slices.Equal(have, want) {
		t.Errorf("runGamesOld() called fatal with args %v, want %v", have, want)
	}
}

func TestFirstConnectionGets110Response(t *testing.T) {
	cs, ss := net.Pipe()
	defer cs.Close()
	defer ss.Close()

	set2SecondTimeout(cs)
	listen := mockListen{
		connections: []net.Conn{ss},
	}
	conf := config{
		listener: &listen,
		fatal:    nullFatal,
	}

	runGamesOld(conf)

	buf := make([]byte, 100)
	n, err := cs.Read(buf)
	if err != nil {
		t.Errorf("client Read() returned error %v, want nil", err)
		return
	}

	response := string(buf[:n])
	checkResponse110(t, response)
}

func TestSecondConnectionGets111Response(t *testing.T) {
	var null nullNetConn
	cs, ss := net.Pipe()
	defer cs.Close()
	defer ss.Close()

	set2SecondTimeout(cs)
	listen := mockListen{
		connections: []net.Conn{null, ss},
	}
	conf := config{
		listener: &listen,
		fatal:    nullFatal,
	}

	runGamesOld(conf)

	buf := make([]byte, 100)
	n, err := cs.Read(buf)
	if err != nil {
		t.Errorf("client Read() returned error %v, want nil", err)
		return
	}

	response := string(buf[:n])
	checkResponse111(t, response)
}

func TestFirstConnectionGets111ResponseAfterSecondConnectionEstablished(t *testing.T) {
	var null nullNetConn

	cs, ss := net.Pipe()
	defer cs.Close()
	defer ss.Close()

	set2SecondTimeout(cs)
	listen := mockListen{
		connections: []net.Conn{ss, null},
	}
	conf := config{
		listener: &listen,
		fatal:    nullFatal,
	}

	runGamesOld(conf)

	buf := make([]byte, 100)
	cs.Read(buf) // discard 110 response
	n, err := cs.Read(buf)

	if err != nil {
		t.Errorf("client Read() returned error %v, want nil", err)
		return
	}

	response := string(buf[:n])
	checkResponse111(t, response)
}

type mockListen struct {
	acceptCalled        bool
	numAcceptsRemaining int
	acceptCallCount     int

	connections []net.Conn
	nextConn    int
}

var mockFunctionErr = fmt.Errorf("mock function")

func (m *mockListen) Accept() (c net.Conn, err error) {
	m.acceptCalled = true

	m.numAcceptsRemaining--
	m.acceptCallCount++

	var null nullNetConn

	if len(m.connections) > 0 {
		if m.nextConn < len(m.connections) {
			m.nextConn++
			return m.connections[m.nextConn-1], nil
		} else {
			return null, mockFunctionErr
		}
	} else if m.numAcceptsRemaining <= 0 {
		if m.numAcceptsRemaining < -10 {
			panic("excessive calls to mock Accept() after returning error condition")
		}
		return nil, mockFunctionErr
	} else {
		return null, nil
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
