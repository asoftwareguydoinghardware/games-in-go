package tictactoeserver

import (
	"context"
	"errors"
	"net"
)

type config struct {
	listener net.Listener
	fatal    func(v ...any)
	printf   func(format string, v ...any)
}

func runGamesOld(conf config) {
	accept := conf.listener.Accept
	fatal := conf.fatal

	first := true
	for {
		conn, err := accept()
		if err != nil {
			fatal(err)
			return
		}

		var msg string

		if first {
			go handleFirstConn(conn)
			first = false
			continue
		} else {
			msg = "111 All players are present.\r\n"
		}
		bytes := []byte(msg)
		go conn.Write(bytes)
	}
}

func runGames(ctx context.Context, conf config) {
	var err error
	var c net.Conn
	first := true

	for {
		c, err = conf.listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			return
		}
		if first {
			go func(c net.Conn) {
				msg := "110 You are the initial player; waiting for opponent.\r\n"
				c.Write([]byte(msg))
				msg = "111 All players are present.\r\n"
				c.Write([]byte(msg))
			}(c)
		} else {
			msg := "111 All players are present.\r\n"
			go c.Write([]byte(msg))
		}
		first = false
	}
}

func handleFirstConn(conn net.Conn) {
	msg := "110 You are the initial player, waiting for opponent.\r\n"
	bytes := []byte(msg)
	conn.Write(bytes)

	msg = "111 All players are present.\r\n"
	bytes = []byte(msg)
	conn.Write(bytes)
}
