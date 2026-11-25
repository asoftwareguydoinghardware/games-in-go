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

func TestAddSecondPlayerExists(t *testing.T) {
	c1, c2 := net.Pipe()
	defer c1.Close()
	defer c2.Close()

	c3, c4 := net.Pipe()
	defer c3.Close()
	defer c4.Close()

	p := createPlayer(c1)
	p.addSecondPlayer(c3)
}
