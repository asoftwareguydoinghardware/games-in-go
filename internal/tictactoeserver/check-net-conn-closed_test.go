package tictactoeserver

import (
	"io"
	"net"
	"time"
)

// simpler version of this which just had a bool value did not work as
// desired because there was essentially a race condition related to where
// the bool was set and the code checking the bool.
type checkClosedNetConn struct {
	sender, receiver net.Conn
}

func (*checkClosedNetConn) Read([]byte) (int, error) {
	return 0, io.EOF
}

func (*checkClosedNetConn) Write(b []byte) (int, error) {
	return len(b), nil
}

func (n *checkClosedNetConn) Close() error {
	n.sender.Write([]byte{})
	return nil
}

func (*checkClosedNetConn) LocalAddr() (a net.Addr) {
	return a
}

func (*checkClosedNetConn) RemoteAddr() (a net.Addr) {
	return a
}

func (*checkClosedNetConn) SetDeadline(time.Time) error {
	return nil
}

func (*checkClosedNetConn) SetReadDeadline(time.Time) error {
	return nil
}

func (*checkClosedNetConn) SetWriteDeadline(time.Time) error {
	return nil
}
