package tictactoeserver

import (
	"net"
)

type player struct{}

func PlayGame(p1, p2 net.Conn) {
}

func createPlayer(c net.Conn) (p *player) {
	msg := "110 You are the initial player, waiting for opponent.\r\n"
	bytes := []byte(msg)
	go c.Write(bytes)
	return p
}

func (p *player) addSecondPlayer(c net.Conn) {
	msg := "111 All players are present.\r\n"
	bytes := []byte(msg)
	c.Write(bytes)
}
