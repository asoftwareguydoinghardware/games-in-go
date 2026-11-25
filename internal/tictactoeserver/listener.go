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

	for {
		c, err = conf.listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			close(secondConnReady)
			return
		}
		if first {
			go func(c net.Conn) {
				msg := "110 You are the initial player; waiting for opponent.\r\n"
				c.Write([]byte(msg))
				_, ok := <-secondConnReady
				if !ok {
					return
				}
				msg = "111 All players are present.\r\n"
				c.Write([]byte(msg))
			}(c)
		} else {
			msg := "111 All players are present.\r\n"
			go c.Write([]byte(msg))
			secondConnReady <- struct{}{}
		}
		first = false
	}
}
