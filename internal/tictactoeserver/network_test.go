package tictactoeserver

import (
	"net"
	"strings"
	"testing"
	"testing/synctest"
	"time"
	"unicode"
)

func TestCreatePlayerConfirmsConnection(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
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
		checkResponse110(t, response)
	})
}

func TestAddSecondPlayerAnnounceEveryonePresent(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var c1 nullNetConn
		p2c, p2s := net.Pipe()
		defer p2c.Close()
		defer p2s.Close()

		p1 := createPlayer(c1)
		p1.addSecondPlayer(p2s)

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

		response := string(buf[:n])
		checkResponse111(t, response)
	})
}

func TestAddSecondPlayerAnnouncesEveryonePresentToFirstPlayer(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		t.Skip("There is a race condition in the code related to the order of Write() calls in two different goroutines, see comments in code")

		/*
			createPlayer invokes go c.Write() for the 110 message.
			addSecondPlayer invokes go p.c.Write() for the 111 message.

			The go was needed to prevent test code from blocking since the
			net.Comm associated with the Write method is from a net.Pipe()
			which synchronizes the Read()/Write() on the two ends of the pipe.

			It seems reasonable to write a test case using at least one mock
			connection to hit the race condition by using a mutex or similar
			so that the two Write()s occur in the "wrong" order.  Unfortunately,
			that test would block once the code was fixed because the fix seems
			to be to do the two Writes() sequentially in the same goroutine
			(since the mock connection would essentially stall the first write
			until after the second occurred; and that can not happen with
			sequential code).
		*/

		var p2s nullNetConn
		p1c, p1s := net.Pipe()
		defer p1c.Close()
		defer p1s.Close()

		p1 := createPlayer(p1s)
		p1.addSecondPlayer(p2s)

		buf := make([]byte, 100)
		p1c.Read(buf)
		n, err := p1c.Read(buf)

		if err != nil {
			t.Errorf("client Read() returned error %v, want nil", err)
			return
		}
		if n < 3 {
			t.Errorf("client Read() returned n = %d, want at least 3", n)
			return
		}

		response := string(buf[:n])
		checkResponse111(t, response)
	})
}

func checkResponse(t *testing.T, response string, wantCode string) {
	if have, min := len(response), 6; have < min {
		t.Errorf("Expected response of more than %d bytes, have %d bytes", min, have)
		return
	}
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
	if len(msg) == 0 {
		t.Errorf("Want text after whitespace in reponse %q", response)
		return
	}
	first := rune(msg[0])
	if !unicode.IsUpper(first) {
		t.Errorf("Want sentences in messages, message %q does not start with an upper case letter (%q)", msg, first)
	}
}

func checkMessageContains(t *testing.T, msg string, wantList []string) {
	if have, want := len(msg), 3; have <= want {
		t.Errorf("message %q too short, need at least %d characters have %d", msg, want, have)
		return
	}
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

func checkResponse110(t *testing.T, response string) {
	checkResponse(t, response, "110")

	wantList := []string{
		"you are",
		"initial player",
		"waiting for opponent",
	}
	checkMessageContains(t, response, wantList)
}

func checkResponse111(t *testing.T, response string) {
	checkResponse(t, response, "111")

	wantList := []string{
		"all players are present",
	}
	checkMessageContains(t, response, wantList)
}
