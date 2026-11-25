package tictactoeserver

import (
	"net"
)

type config struct {
	listener net.Listener
}

func runGames(conf config) {
	for _, err := conf.listener.Accept(); err == nil; _, err = conf.listener.Accept() {
		// nothing for now
	}
}
