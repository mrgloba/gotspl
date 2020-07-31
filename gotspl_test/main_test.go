/*
 * Copyright 2020 Anton Globa
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
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
		req, err := rw.ReadBytes(0x0A)
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
