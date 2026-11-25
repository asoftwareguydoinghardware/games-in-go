package tictactoeserver

import (
	"net"
	"strings"
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

	msg := strings.ToLower(strings.Trim(string(buf[3:]), "\000"))
	wantList := []string{
		"you are",
		"initial player",
		"waiting for opponent",
	}
	for _, want := range wantList {
		if !strings.Contains(msg, want) {
			t.Errorf("Want message %q to contain %q", msg, want)
		}
	}
}
