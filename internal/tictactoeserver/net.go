package tictactoeserver

import (
	"net"
)

type player struct{}

func PlayGame(p1, p2 net.Conn) {
}

func createPlayer(c net.Conn) (p *player) {
	msg := "110you are the initial player, waiting for opponent.\r\n"
	bytes := []byte(msg)
	c.Write(bytes)
	return p
}

func (p *player) addSecondPlayer(c net.Conn) {
}
