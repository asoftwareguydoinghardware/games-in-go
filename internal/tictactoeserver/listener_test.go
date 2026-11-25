package tictactoeserver

import (
	"context"
	"fmt"
	"net"
	"testing"
	"testing/synctest"
)

func TestRunGamesAcceptsAConnection(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var listen mockListen2
		conf := config{
			listener: &listen,
			fatal:    nullFatal,
		}

		runGames(context.Background(), conf)

		if listen.acceptCount == 0 {
			t.Errorf("runGames() did not call listen.Accept(conf)")
		}
	})
}

func TestRunGamesAcceptsConnections(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		connectionCount := []int{2, 5, 1}

		for _, count := range connectionCount {
			listen := mockListen2{callsUntilShutdown: count}
			conf := config{
				listener: &listen,
				fatal:    nullFatal,
			}

			runGames(context.Background(), conf)

			if have, want := listen.acceptCount, count; have != want {
				t.Errorf("with mock listener, have %d calls to Accept(), want %d", have, want)
			}
		}
	})
}

func TestRunGamesContinuesAcceptingUntilNetErrClosed(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		const callCount = 3 // arbitrary value
		// The mockListen2 does not actually use the Close()
		// method, but simulates that by returning net.ErrClosed.  This
		// test is essentially making sure the loop (does not) terminate
		// when Accept() (does not) return net.ErrClosed.

		listen := mockListen2{callsUntilShutdown: callCount}
		conf := config{
			listener: &listen,
		}

		runGames(context.Background(), conf)

		if have, want := listen.acceptCount, callCount; have != want {
			t.Errorf("With mock listner returning net.ErrClosed after %d calls to Accept(), have %d calls, want %d", callCount, have, want)
		}
	})
}

func TestRunGamesHandlesWrappedNetErrClosed(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var (
			callCount  = 5
			wrappedErr = fmt.Errorf("%w", net.ErrClosed)
		)

		listen := mockListen2{
			callsUntilShutdown: callCount,
			shutdownErr:        wrappedErr,
		}
		conf := config{
			listener: &listen,
		}

		runGames(context.Background(), conf)

		if have, want := listen.acceptCount, callCount; have != want {
			t.Errorf("With mock listner returning net.ErrClosed after %d calls to Accept(), have %d calls, want %d", callCount, have, want)
		}
	})
}

func TestFirstConnectionGets110Response(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		cs, ss := net.Pipe()
		defer cs.Close()
		defer ss.Close()

		set2SecondTimeout(cs)
		listen := mockListen2{
			connections: []net.Conn{ss},
		}
		conf := config{
			listener: &listen,
			fatal:    nullFatal,
		}

		runGames(context.Background(), conf)

		buf := make([]byte, 100)
		n, err := cs.Read(buf)
		if err != nil {
			t.Errorf("client Read() returned error %v, want nil", err)
			return
		}

		response := string(buf[:n])
		checkResponse110(t, response)
	})
}

func TestSecondConnectionGets111Response(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var null nullNetConn
		cs, ss := net.Pipe()
		defer cs.Close()
		defer ss.Close()

		set2SecondTimeout(cs)
		listen := mockListen2{
			connections: []net.Conn{null, ss},
		}
		conf := config{
			listener: &listen,
			fatal:    nullFatal,
		}

		runGames(context.Background(), conf)

		buf := make([]byte, 100)
		n, err := cs.Read(buf)
		if err != nil {
			t.Errorf("client Read() returned error %v, want nil", err)
			return
		}

		response := string(buf[:n])
		checkResponse111(t, response)
	})
}

func TestFirstConnectionGets111ResponseAfterSecondConnectionEstablished(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var null nullNetConn

		cs, ss := net.Pipe()
		defer cs.Close()
		defer ss.Close()

		set2SecondTimeout(cs)
		listen := mockListen2{
			connections: []net.Conn{ss, null},
		}
		conf := config{
			listener: &listen,
			fatal:    nullFatal,
		}

		runGames(context.Background(), conf)

		buf := make([]byte, 100)
		cs.Read(buf) // discard 110 response
		n, err := cs.Read(buf)

		if err != nil {
			t.Errorf("client Read() returned error %v, want nil", err)
			return
		}

		response := string(buf[:n])
		checkResponse111(t, response)
	})
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

type mockListen2 struct {
	callsUntilShutdown int

	acceptCount int
	shutdownErr error

	connections    []net.Conn
	nextConnection int
}

func (m *mockListen2) Accept() (c net.Conn, err error) {
	var n nullNetConn

	c = n
	err = mockFunctionErr

	m.acceptCount++

	if m.nextConnection < len(m.connections) {
		c = m.connections[m.nextConnection]
		m.nextConnection++
		if m.callsUntilShutdown == 0 {
			m.callsUntilShutdown = len(m.connections) + 1
		}
	}
	m.callsUntilShutdown--
	if m.callsUntilShutdown < -10 {
		panic("excessive calls to Accept()")
	}
	if m.callsUntilShutdown <= 0 {
		if m.shutdownErr == nil {
			m.shutdownErr = net.ErrClosed
		}
		c = nil
		err = m.shutdownErr
	}

	return c, err
}

func (m *mockListen2) Close() (err error) {
	return err
}

func (m *mockListen2) Addr() (n net.Addr) {
	return n
}

func nullFatal(v ...any) {
}
