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

	for {
		conn, err := accept()
		if err != nil {
			fatal(err)
			return
		}
		msg := "110 X.\r\n"
		bytes := []byte(msg)
		go conn.Write(bytes)
	}
}
