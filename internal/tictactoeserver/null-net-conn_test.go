package tictactoeserver

import (
	"io"
	"net"
	"time"
)

type nullNetConn struct {
	simpleWriteHook func()
}

func (nullNetConn) Read([]byte) (int, error) {
	return 0, io.EOF
}

func (n nullNetConn) Write(b []byte) (int, error) {
	if n.simpleWriteHook != nil {
		n.simpleWriteHook()
	}
	return len(b), nil
}

func (nullNetConn) Close() error {
	return nil
}

func (nullNetConn) LocalAddr() (a net.Addr) {
	return a
}

func (nullNetConn) RemoteAddr() (a net.Addr) {
	return a
}

func (nullNetConn) SetDeadline(time.Time) error {
	return nil
}

func (nullNetConn) SetReadDeadline(time.Time) error {
	return nil
}

func (nullNetConn) SetWriteDeadline(time.Time) error {
	return nil
}
