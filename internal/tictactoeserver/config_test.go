package tictactoeserver

import (
	"net"
)

func newConfig(l net.Listener) (conf config) {
	conf.listener = l
	conf.printf = nil

	return conf
}
