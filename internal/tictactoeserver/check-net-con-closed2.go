package tictactoeserver

import (
	"io"
	"net"
	"time"
)

// simpler version of this which just had a bool value did not work as
// desired because there was essentially a race condition related to where
// the bool was set and the code checking the bool.
type checkClosedNetConn2 struct {
	closed bool
}

func (*checkClosedNetConn2) Read([]byte) (int, error) {
	return 0, io.EOF
}

func (*checkClosedNetConn2) Write(b []byte) (int, error) {
	return len(b), nil
}

func (n *checkClosedNetConn2) Close() error {
	n.closed = true
	return nil
}

func (*checkClosedNetConn2) LocalAddr() (a net.Addr) {
	return a
}

func (*checkClosedNetConn2) RemoteAddr() (a net.Addr) {
	return a
}

func (*checkClosedNetConn2) SetDeadline(time.Time) error {
	return nil
}

func (*checkClosedNetConn2) SetReadDeadline(time.Time) error {
	return nil
}

func (*checkClosedNetConn2) SetWriteDeadline(time.Time) error {
	return nil
}
