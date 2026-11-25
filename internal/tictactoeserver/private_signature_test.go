package tictactoeserver

import (
	"context"
	"log"
	"net"
	"testing"
)

func TestCreatePlayerExists(t *testing.T) {
	var p *player
	var c net.Conn
	var null nullNetConn

	c = null
	p = createPlayer(c)
	_ = p
}

func TestAddSecondPlayerExists(t *testing.T) {
	var n1, n3 nullNetConn
	var c1, c3 net.Conn

	c1, c3 = n1, n3

	p := createPlayer(c1)
	p.addSecondPlayer(c3)
}

func TestRunGamesOldExists(t *testing.T) {
	var mock mockListen
	conf := config{
		listener: &mock,
		fatal:    nullFatal,
	}

	runGamesOld(conf)
}

func TestRunGamesExists(t *testing.T) {
	var mock mockListen2
	var ctx context.Context = context.Background()
	conf := config{
		listener: &mock,
		fatal:    nullFatal,
	}

	runGames(ctx, conf)
}

func TestMembersOfConfigType(t *testing.T) {
	var conf config
	var l net.Listener

	conf.listener = l
	conf.fatal = log.Fatal
	conf.printf = log.Printf
	_ = conf
}
