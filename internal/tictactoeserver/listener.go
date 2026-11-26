package tictactoeserver

import (
	"context"
	"errors"
	ttt "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tictactoe"
	"net"
)

type game interface {
	AssignPlayer(ctx context.Context, pn int, player ttt.Player)
	InitialPlayer(pn int)
}

type config struct {
	listener net.Listener
	printf   func(format string, v ...any)
	game     func() game
}

type player1Info struct {
	c              net.Conn
	haveSecondConn chan<- struct{}
	conf           config
}

func runGames(ctx context.Context, conf config) {
	var err error
	var c net.Conn
	first := true
	haveSecondConn := make(chan struct{})
	localCtx, cancel := context.WithCancel(ctx)

	for {
		c, err = conf.listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			cancel()
			return
		}
		if first {
			go player0(localCtx, c, haveSecondConn)
		} else {
			p1Info := player1Info{
				c:              c,
				haveSecondConn: haveSecondConn,
				conf:           conf,
			}
			go player1(localCtx, p1Info)
		}
		first = false
	}
}

func player0(ctx context.Context, c net.Conn, ready <-chan struct{}) {
	defer c.Close()
	msg := "110 You are the initial player; waiting for opponent.\r\n"
	c.Write([]byte(msg))
	select {
	case _, ok := <-ready:
		if !ok {
			return
		}
	case <-ctx.Done():
		return
	}
	msg = "111 All players are present.\r\n"
	c.Write([]byte(msg))
}

func player1(ctx context.Context, info player1Info) {
	c := info.c
	ready := info.haveSecondConn
	conf := info.conf

	defer c.Close()
	msg := "111 All players are present.\r\n"
	c.Write([]byte(msg))
	game := conf.game()

	var p0, p1 nullPlayer
	p0.playerNumber = -3

	game.InitialPlayer(0)
	game.AssignPlayer(context.Background(), 0, p0)
	game.AssignPlayer(context.Background(), 1, p1)

	select {
	case ready <- struct{}{}:
	case <-ctx.Done():
		return
	}
}
