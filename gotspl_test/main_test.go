package gotspl_test

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

type TCPServer struct {
	server  net.Listener
	address string
}

func NewTCPServer(address string) TCPServer {
	return TCPServer{address: address}
}

func (t *TCPServer) Run() (err error) {
	t.server, err = net.Listen("tcp", t.address)
	if err != nil {
		return
	}
	for {
		return t.handleConnections()
	}
	return
}

func (t *TCPServer) handleConnections() error {
	for {
		conn, err := t.server.Accept()
		if err != nil || conn == nil {
			err = errors.New("could not accept connection")
			break
		}

		go t.handleConnection(conn)
	}
	return nil
}

func (t *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {
		req, err := rw.ReadBytes(0x13)
		if err != nil {
			rw.WriteString("failed to read input")
			rw.Flush()
			return
		}

		rw.WriteString(fmt.Sprintf("RCVD: %s", req))
		rw.Flush()
	}
}

// Close shuts down the TCP Server
func (t *TCPServer) Close() (err error) {
	return t.server.Close()
}
