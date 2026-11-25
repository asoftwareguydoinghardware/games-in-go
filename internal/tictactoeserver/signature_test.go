package tictactoeserver_test

import (
	server "github.com/ASoftwareGuyDoingHardware/games-in-go/internal/tictactoeserver"
	"net"
	"testing"
	"testing/synctest"
)

func TestPlayGameExists(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var p1, p2 net.Conn

		server.PlayGame(p1, p2)
	})
}
