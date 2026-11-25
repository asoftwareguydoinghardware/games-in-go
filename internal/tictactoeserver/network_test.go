package tictactoeserver

import (
	"net"
	"testing"
)

func TestCreatePlayerConfirmsConnection(t *testing.T) {
	c1, c2 := net.Pipe()
	defer c1.Close()
	defer c2.Close()

	go createPlayer(c1)

	buf := make([]byte, 100)
	n, err := c2.Read(buf)

	if err != nil {
		t.Errorf("client Read() returned error %v, want nil", err)
		return
	}
	if n < 3 {
		t.Errorf("client Read() returned n = %d, want at least 3", n)
		return
	}

	if have, want := string(buf[:3]), "110"; have != want {
		t.Errorf("client Read() returned (first 3 bytes) %q, want %q", have, want)
	}
}
