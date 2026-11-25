package tictactoeserver

import (
	"net"
)

func newConfig(l net.Listener) (conf config) {
	var ng nullGame

	conf.listener = l
	conf.printf = nil
	conf.game = func() game { return &ng }

	return conf
}
