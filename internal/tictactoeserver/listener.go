package tictactoeserver

import (
	"context"
	"errors"
	"net"
)

type config struct {
	listener net.Listener
	printf   func(format string, v ...any)
}

func runGames(ctx context.Context, conf config) {
	var err error
	var c net.Conn
	first := true
	haveSecondConn := make(chan struct{})
	localCtx, cancel := context.WithCancel(ctx)

	for {
		c, err = conf.listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			cancel()
			return
		}
		if first {
			go player0(localCtx, c, haveSecondConn)
		} else {
			go player1(localCtx, c, haveSecondConn)
		}
		first = false
	}
}

func player0(ctx context.Context, c net.Conn, ready <-chan struct{}) {
	defer c.Close()
	msg := "110 You are the initial player; waiting for opponent.\r\n"
	c.Write([]byte(msg))
	select {
	case _, ok := <-ready:
		if !ok {
			return
		}
	case <-ctx.Done():
		return
	}
	msg = "111 All players are present.\r\n"
	c.Write([]byte(msg))
}

func player1(ctx context.Context, c net.Conn, ready chan<- struct{}) {
	defer c.Close()
	msg := "111 All players are present.\r\n"
	c.Write([]byte(msg))
	select {
	case ready <- struct{}{}:
	case <-ctx.Done():
		return
	}
}
