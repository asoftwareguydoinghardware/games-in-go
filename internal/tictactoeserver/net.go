package tictactoeserver

import (
	"net"
)

type player struct{
	c net.Conn
}

func PlayGame(p1, p2 net.Conn) {
}

func createPlayer(c net.Conn) (p *player) {
	p = new(player)
	p.c = c
	msg := "110 You are the initial player, waiting for opponent.\r\n"
	bytes := []byte(msg)
	go c.Write(bytes)
	return p
}

func (p *player) addSecondPlayer(c net.Conn) {
	msg := "111 All players are present.\r\n"
	bytes := []byte(msg)
	go c.Write(bytes)

	bytes = make([]byte, 3)
	go p.c.Write(bytes)
}
