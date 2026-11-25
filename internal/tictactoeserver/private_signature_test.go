package tictactoeserver

import (
	"context"
	"log"
	"net"
	"testing"
	"testing/synctest"
)

func TestCreatePlayerExists(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var p *player
		var c net.Conn
		var null nullNetConn

		c = null
		p = createPlayer(c)
		_ = p
	})
}

func TestAddSecondPlayerExists(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var n1, n3 nullNetConn
		var c1, c3 net.Conn

		c1, c3 = n1, n3

		p := createPlayer(c1)
		p.addSecondPlayer(c3)
	})
}

func TestRunGamesExists(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var mock mockListen2
		var ctx context.Context = context.Background()
		conf := config{
			listener: &mock,
		}

		runGames(ctx, conf)
	})
}

func TestMembersOfConfigType(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var conf config
		var l net.Listener

		conf.listener = l
		conf.printf = log.Printf
		_ = conf
	})
}
