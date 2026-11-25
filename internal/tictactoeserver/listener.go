package tictactoeserver

import (
	"net"
)

type config struct {
	listener net.Listener
}

func runGames(conf config) {
	accept := conf.listener.Accept

	for {
		_, err := accept()
		if err != nil {
			return
		}
	}
}
