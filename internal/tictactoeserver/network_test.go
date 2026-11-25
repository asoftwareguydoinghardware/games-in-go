package tictactoeserver

import (
	"net"
	"strings"
	"testing"
	"unicode"
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

	checkMessage(t, buf, "110")

	msg := strings.ToLower(string(buf[3:n]))
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

	/* Ideally I would check len(msg) >= 2 after stripping the first
	three characters.  However, the sequence in which I am writing
	the test cases means that test would never fail, and strict TDD
	forbids me from writing that.

	In the unlikely case, len(msg) < 2, the code below will panic,
	so that is likely sufficient.

	OK, the comment is getting out of date as the check should now
	be three.  However, the underlying point remains:  I acknowledge
	the checks should be better, but strict TDD prevents that.
	*/
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

	msg = strings.TrimLeftFunc(string(buf[3:n]), unicode.IsSpace)
	first := rune(msg[0])
	if !unicode.IsUpper(first) {
		t.Errorf("Want sentences in messages, message %q does not start with an upper case letter (%q)", msg, first)
	}
}

func checkMessage(t *testing.T, buf []byte, wantCode string) {
	if have, want := string(buf[:3]), wantCode; have != want {
		t.Errorf("client Read() returned (first 3 bytes) %q, want %q", have, want)
	}
}
