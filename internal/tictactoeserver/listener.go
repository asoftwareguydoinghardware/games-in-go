package tictactoeserver

import (
	"net"
)

type config struct {
	listener net.Listener
	fatal    func(v ...any)
}

func runGames(conf config) {
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
			msg = "110 You are the initial player, waiting for opponent.\r\n"
			first = false
		} else {
			msg = "111 All players are present.\r\n"
		}
		bytes := []byte(msg)
		go conn.Write(bytes)
	}
}
