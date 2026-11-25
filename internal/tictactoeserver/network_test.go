package tictactoeserver

import (
	"net"
	"strings"
	"testing"
	"time"
	"unicode"
)

func TestCreatePlayerConfirmsConnection(t *testing.T) {
	c1, c2 := net.Pipe()
	defer c1.Close()
	defer c2.Close()

	createPlayer(c1)

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

	response := string(buf[:n])
	checkResponse(t, response, "110")

	wantList := []string{
		"you are",
		"initial player",
		"waiting for opponent",
	}
	checkMessageContains(t, response, wantList)
}

func TestAddSecondPlayerAnnounceEveryonePresent(t *testing.T) {
	var c1 nullNetConn
	p2c, p2s := net.Pipe()
	defer p2c.Close()
	defer p2s.Close()

	p1 := createPlayer(c1)
	go p1.addSecondPlayer(p2s)

	buf := make([]byte, 100)
	n, err := p2c.Read(buf)

	if err != nil {
		t.Errorf("client Read() returned error %v, want nil", err)
		return
	}
	if n < 3 {
		t.Errorf("client Read() returned n = %d, want at least 3", n)
		return
	}
}

func checkResponse(t *testing.T, response string, wantCode string) {
	if have, want := response[:3], wantCode; have != want {
		t.Errorf("client Read() returned (first 3 bytes) %q, want %q", have, want)
	}

	msg := response[3:]
	theRest := msg[len(msg)-2:]
	if have, want := theRest, "\r\n"; have != want {
		t.Errorf("Message %q does not end with %q", msg, want)
	}

	lastChar := rune(msg[len(msg)-3])
	if have, want := lastChar, '.'; have != want {
		t.Errorf("Want sentences in messages, msg %q does not end with %q", msg, want)
	}

	firstCharOfMsg := rune(msg[0])
	if !unicode.IsSpace(firstCharOfMsg) {
		t.Errorf("Want a whitespace character between response code and human readable message, first character is %q want a space character", firstCharOfMsg)
	}

	msg = strings.TrimLeftFunc(msg, unicode.IsSpace)
	first := rune(msg[0])
	if !unicode.IsUpper(first) {
		t.Errorf("Want sentences in messages, message %q does not start with an upper case letter (%q)", msg, first)
	}
}

func checkMessageContains(t *testing.T, msg string, wantList []string) {
	msg = strings.ToLower(msg[3:])
	for _, want := range wantList {
		if !strings.Contains(msg, want) {
			t.Errorf("Want message %q to contain %q", msg, want)
		}
	}
}

func set2SecondTimeout(c net.Conn) (err error) {
	deadline := time.Now().Add(2 * time.Second)
	err = c.SetDeadline(deadline)
	return err
}
