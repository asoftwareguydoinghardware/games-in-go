package tictactoeserver

import (
	"net"
	"testing"
)

func TestCreatePlayerExists(t *testing.T) {
	var p *player
	var c net.Conn

	p = createPlayer(c)
	_ = p
}
