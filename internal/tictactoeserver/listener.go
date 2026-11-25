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
	secondConnReady := make(chan struct{})
	localCtx, cancel := context.WithCancel(ctx)

	for {
		c, err = conf.listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			cancel()
			return
		}
		if first {
			go func(ctx context.Context, c net.Conn) {
				defer c.Close()
				msg := "110 You are the initial player; waiting for opponent.\r\n"
				c.Write([]byte(msg))
				select {
				case _, ok := <-secondConnReady:
					if !ok {
						return
					}
				case <-ctx.Done():
					return
				}
				msg = "111 All players are present.\r\n"
				c.Write([]byte(msg))
			}(localCtx, c)
		} else {
			go func(ctx context.Context, c net.Conn) {
				defer c.Close()
				msg := "111 All players are present.\r\n"
				c.Write([]byte(msg))
				select {
				case secondConnReady <- struct{}{}:
				case <-ctx.Done():
					return
				}
			}(localCtx, c)
		}
		first = false
	}
}
