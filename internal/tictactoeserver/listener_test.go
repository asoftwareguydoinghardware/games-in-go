package tictactoeserver

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"testing"
	"testing/synctest"
	"time"
)

func TestRunGamesAcceptsAConnection(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var listen mockListen2
		conf := newConfig(&listen)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		runGames(ctx, conf)

		if listen.acceptCount == 0 {
			t.Errorf("runGames() did not call listen.Accept(conf)")
		}
	})
}

func TestRunGamesAcceptsConnections(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		connectionCount := []int{2, 5, 1}

		for _, count := range connectionCount {
			if count > 2 {
				t.Skip("skipping more than 2 connections")
			}
			t.Logf("testing %d connections", count)
			listen := mockListen2{callsUntilShutdown: count}
			conf := newConfig(&listen)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			runGames(ctx, conf)

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
		conf := newConfig(&listen)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		runGames(ctx, conf)

		if have, want := listen.acceptCount, callCount; have != want {
			t.Errorf("With mock listner returning net.ErrClosed after %d calls to Accept(), have %d calls, want %d", callCount, have, want)
		}
	})
}

func TestRunGamesHandlesWrappedNetErrClosed(t *testing.T) {
	t.Skip("skipping tests using more than 2 connections")
	synctest.Test(t, func(t *testing.T) {
		var (
			callCount  = 5
			wrappedErr = fmt.Errorf("%w", net.ErrClosed)
		)

		listen := mockListen2{
			callsUntilShutdown: callCount,
			shutdownErr:        wrappedErr,
		}
		conf := newConfig(&listen)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		runGames(ctx, conf)

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
		conf := newConfig(&listen)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		runGames(ctx, conf)

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
		conf := newConfig(&listen)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		runGames(ctx, conf)

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
		conf := newConfig(&listen)

		// net.Pipe connections block until there is a matched
		// set of read/write requests on the two ends of the
		// pipe.  runGames() must be run in its own goroutine
		// so that its i/o on the net.Pipe does not block
		// waiting for the corresponding read/write that occurs
		// after runGames() returns.
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go runGames(ctx, conf)

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

func TestFirstConnectionDoesNotGet111ResponseWhenThereIsNoSecondConnection(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		cs, ss := net.Pipe()
		defer cs.Close()
		defer ss.Close()

		set2SecondTimeout(cs)
		listen := mockListen2{
			connections: []net.Conn{ss},
		}
		conf := newConfig(&listen)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		runGames(ctx, conf)

		buf := make([]byte, 100)
		cs.Read(buf) // discard 110 response
		n, err := cs.Read(buf)
		if have, want := n, 0; have != want {
			t.Errorf("With a single connection, second read returns n = %d, want %d", have, want)
		}
		if err == nil {
			t.Errorf("with a single connection read, second read returned nil error, want a non-nil error")
		}
	})
}

func TestFirstConnectionClosed(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		checkClosed := new(checkClosedNetConn)
		checkClosed.sender, checkClosed.receiver = net.Pipe()

		var null nullNetConn

		listen := mockListen2{
			connections: []net.Conn{checkClosed, null},
		}
		conf := newConfig(&listen)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		runGames(ctx, conf)

		// Would deadlock before change associated with this test
		// was added.
		checkClosed.receiver.Read(nil)
	})
}

func TestSecondConnectionClosed(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		checkClosed := new(checkClosedNetConn)
		checkClosed.sender, checkClosed.receiver = net.Pipe()

		var null nullNetConn

		listen := mockListen2{
			connections: []net.Conn{null, checkClosed},
		}
		conf := newConfig(&listen)

		// needs to run in a seperate goroutine so that the
		// Write() in checkClosedNetConn.Close() will not be
		// blocked; the corresponding Read() occurs after
		// the goroutine is started.
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go runGames(ctx, conf)

		// Would deadlock before change associated with this test
		// was added.
		checkClosed.receiver.Read(nil)
	})
}

func TestFirstconnectionClosedIfContextCancelled(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		checkClosed := new(checkClosedNetConn2)

		listen := mockListen3{
			connections: []net.Conn{checkClosed},
		}

		conf := newConfig(&listen)

		ctx, cancel := context.WithCancel(context.Background())
		go runGames(ctx, conf)
		synctest.Wait()
		cancel()
		synctest.Wait()

		if !checkClosed.closed {
			t.Errorf("when context cancelled net.Conn not closed")
		}
	})
}

func TestSecondConnectionGetsNewGame(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var n1, n2 nullNetConn
		var gameWasCalled bool

		listen := mockListen3{
			connections: []net.Conn{n1, n2},
		}

		ng := func() game {
			var n nullGame

			gameWasCalled = true
			return &n
		}

		conf := newConfig(&listen)
		conf.game = ng

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go runGames(ctx, conf)

		synctest.Wait()

		if !gameWasCalled {
			t.Errorf("runGames() did not call game()")
		}
	})
}

func TestNewGameHappensAfterInitialResponse(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		type state int
		const (
			start state = iota
			gotResponse
			gotRunGame
		)
		sequence := make(chan state, 10)

		var n1, n2 nullNetConn
		n2.simpleWriteHook = func() {
			sequence <- gotResponse
		}

		listen := mockListen3{
			connections: []net.Conn{n1, n2},
		}
		conf := newConfig(&listen)
		conf.game = func() game {
			var ng nullGame

			sequence <- gotRunGame
			return &ng
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		runGames(ctx, conf)

		synctest.Wait()

		s1 := <-sequence
		if s1 != gotResponse {
			t.Errorf("game() called before sending initial response")
		}
	})
}

func TestFistPlayerAssigned(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var n1, n2 nullNetConn
		listen := mockListen3{
			connections: []net.Conn{n1, n2},
		}
		conf := newConfig(&listen)
		var g gameObserver

		conf.game = func() game {
			return &g
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		runGames(ctx, conf)

		synctest.Wait()

		if have := len(g.assignPlayerArgs); have == 0 {
			t.Errorf("AssignPlayer not called")
			return
		}
		if have, want := g.assignPlayerArgs[0].playerNum, 0; have != want {
			t.Errorf("first call to AssignPlayer() called with player number %d, want %d", have, want)
		}
	})
}

func TestSecondPlayerAssigned(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var n1, n2 nullNetConn
		listen := mockListen3{
			connections: []net.Conn{n1, n2},
		}
		conf := newConfig(&listen)
		var g gameObserver

		conf.game = func() game {
			return &g
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		runGames(ctx, conf)

		synctest.Wait()

		if have, want := len(g.assignPlayerArgs), 2; have != want {
			t.Errorf("AssignPlayer called %d times, want %d", have, want)
			return
		}
		if have, want := g.assignPlayerArgs[1].playerNum, 1; have != want {
			t.Errorf("second call to AssignPlayer() called with player number %d, want %d", have, want)
		}
	})
}

func TestAssignedPlayersAreUniqueAndNonNil(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var n1, n2 nullNetConn
		listen := mockListen3{
			connections: []net.Conn{n1, n2},
		}
		conf := newConfig(&listen)
		var g gameObserver

		conf.game = func() game {
			return &g
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		runGames(ctx, conf)

		synctest.Wait()

		if have, want := len(g.assignPlayerArgs), 2; have != want {
			t.Skip("Not enough calls to AssignPlayer()")
			return
		}

		p0 := g.assignPlayerArgs[0].player
		p1 := g.assignPlayerArgs[1].player

		if p0 == nil {
			t.Errorf("AssignPlayer() for player 0 called with nil player argument")
		}
		if p1 == nil {
			t.Errorf("AssignPlayer() for player 1 called with nil player argument")
		}
		if p0 == p1 {
			t.Errorf("AssignPlayer() called with the same player argument for player0 and player1")
		}
	})
}

func TestInitialPlayerIs0ForFirstGame(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var n1, n2 nullNetConn
		listen := mockListen3{
			connections: []net.Conn{n1, n2},
		}
		conf := newConfig(&listen)
		var g gameObserver

		conf.game = func() game {
			return &g
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		runGames(ctx, conf)

		synctest.Wait()

		if have, want := len(g.initialPlayer), 1; have != want {
			t.Errorf("InitialPlayer() called %d times, want %d", have, want)
			return
		}
		if have, want := g.initialPlayer[0], 0; have != want {
			t.Errorf("InitialPlayer() called with %d, want %d", have, want)
		}

	})
}

func TestFirstConnectionGets190Response(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var ns nullNetConn
		c1, s1 := net.Pipe()

		listen := mockListen3{
			connections:       []net.Conn{s1, ns},
			sleepBeforeAccept: true,
		}
		conf := newConfig(&listen)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go runGames(ctx, conf)

		buf := make([]byte, 100)
		c1.Read(buf) // ignored 110 response
		c1.Read(buf) // ignored 111 response
		n, err := c1.Read(buf)
		defer drainConn(c1)

		if have := err; have != nil {
			t.Errorf("Read of third response from server got err %v, want nill", have)
		}

		response := string(buf[:n])
		checkResponse(t, response, "190")

		wantList := []string{
			"you are player 0",
		}
		checkMessageContains(t, response, wantList)

		time.Sleep(2 * time.Second) // Allow Accept() to unblock and terminate
	})
}

func TestSecondConnectionGets191Response(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var ns nullNetConn
		c2, s2 := net.Pipe()

		listen := mockListen3{
			connections:       []net.Conn{ns, s2},
			sleepBeforeAccept: true,
		}
		conf := newConfig(&listen)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go runGames(ctx, conf)

		buf := make([]byte, 100)
		c2.Read(buf) // ignored 111 response
		n, err := c2.Read(buf)
		drainConn(c2)

		if have := err; have != nil {
			t.Errorf("Read on second connection of third response from server got err %v, want nil", have)
		}

		response := string(buf[:n])
		checkResponse(t, response, "191")

		wantList := []string{
			"you are player 1",
		}
		checkMessageContains(t, response, wantList)

		time.Sleep(2 * time.Second) // Allow Accept() to unblock and terminate
	})
}

func TestPlayIsCalled(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var n1, n2 nullNetConn
		listen := mockListen3{
			connections: []net.Conn{n1, n2},
		}
		conf := newConfig(&listen)
		var g gameObserver

		conf.game = func() game {
			return &g
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		runGames(ctx, conf)

		synctest.Wait()

		if !g.playCalled {
			t.Errorf("Play() method was not called")
		}
	})
}

func TestPlayIsCalledAfterInitialPlayerIsCalled(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var n1, n2 nullNetConn
		listen := mockListen3{
			connections: []net.Conn{n1, n2},
		}
		conf := newConfig(&listen)
		var g gameObserver

		conf.game = func() game {
			return &g
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		runGames(ctx, conf)

		synctest.Wait()

		if !g.playCalledAfterInitialPlayerCalled {
			t.Errorf("Play() method must be called after InitialPlayer()")
		}
	})
}

func TestPlayIsCalledAfterAllPlayersAssigned(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var n1, n2 nullNetConn
		listen := mockListen3{
			connections: []net.Conn{n1, n2},
		}
		conf := newConfig(&listen)
		var g gameObserver

		conf.game = func() game {
			return &g
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		runGames(ctx, conf)

		synctest.Wait()

		// Note: other tests verify that the first and second calls
		// to AssignPlayer() specify player number 0 and 1.
		if have, want := len(g.playersAssignedBeforePlay), 2; have != want {
			t.Errorf("Number of players assigned before Play() called = %d, want %d", have, want)
		}
	})
}

func TestFirstPlayerGets120Response(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var ns nullNetConn
		c1, s1 := net.Pipe()

		listen := mockListen3{
			connections:       []net.Conn{s1, ns},
			sleepBeforeAccept: true,
		}
		conf := newConfig(&listen)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go runGames(ctx, conf)

		buf := make([]byte, 100)
		c1.Read(buf) // ignored 110 response
		c1.Read(buf) // ignored 111 response
		c1.Read(buf) // ignored 190 response
		n, err := c1.Read(buf)

		if have := err; have != nil {
			t.Errorf("Read on first connection of fourth response from server got err %v, want nil", have)
		}

		response := string(buf[:n])
		checkResponse(t, response, "120")

		wantList := []string{
			"you get the initial move",
		}
		checkMessageContains(t, response, wantList)

		time.Sleep(2 * time.Second) // Allow Accept() to unblock and terminate
	})
}

func TestSecondPlayerGets121Response(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var ns nullNetConn
		c2, s2 := net.Pipe()

		listen := mockListen3{
			connections:       []net.Conn{ns, s2},
			sleepBeforeAccept: true,
		}
		conf := newConfig(&listen)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go runGames(ctx, conf)

		buf := make([]byte, 100)
		c2.Read(buf) // ignored 111 response
		c2.Read(buf) // ignored 191 response
		n, err := c2.Read(buf)

		if have := err; have != nil {
			t.Errorf("Read on second connection of third response from server got err %v, want nil", have)
		}

		response := string(buf[:n])
		checkResponse(t, response, "121")

		wantList := []string{
			"you get the second move",
		}
		checkMessageContains(t, response, wantList)

		time.Sleep(2 * time.Second) // Allow Accept() to unblock and terminate
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

	synctest.Wait()
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

type mockListen3 struct {
	connections    []net.Conn
	nextConnection int
	closed         bool

	// Setting sleepBeforeAccept forces Accept() to stall until all
	// goroutines are durably blocked.  If this is not done, then it
	// is possible that Accept() will return a wrapped net.ErrClosed
	// which causes runGames() to cancel the context in which the
	// other goroutines are running.  Cancelling the other goroutines
	// early means that they may not handle all the i/o that they
	// should have handled.
	sleepBeforeAccept bool
}

func (m *mockListen3) Accept() (c net.Conn, err error) {

	if m.sleepBeforeAccept {
		time.Sleep(time.Second)
		synctest.Wait()
	}
	if m.nextConnection < len(m.connections) {
		c = m.connections[m.nextConnection]
		m.nextConnection++
		err = nil
	} else {
		c = nil
		err = fmt.Errorf("%w", net.ErrClosed)
	}

	return c, err
}

func (m *mockListen3) Close() (err error) {
	m.closed = true
	return err
}

func (m *mockListen3) Addr() (n net.Addr) {
	return n
}

func nullFatal(v ...any) {
}

func drainConn(c net.Conn) {
	readCount := 0

	buf := make([]byte, 100)
	for {
		n, err := c.Read(buf)
		if n == 0 && errors.Is(err, io.EOF) {
			return
		}
		readCount++
		if readCount > 100 {
			panic("drainConn called read more than 100 times")
		}
	}
}
