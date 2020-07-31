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
	"bytes"
	"github.com/mrgloba/gotspl/gotspl"
	"testing"
	"time"
)

var server TCPServer

func init() {
	server = NewTCPServer(":3234")
	go func() {
		server.Run()
	}()
}

func TestEthernetTSPLClient(t *testing.T) {
	time.Sleep(1 * time.Second)
	client := gotspl.NewEthernetTSPLClient(":3234")
	if client.IsConnected() {
		t.Error("EthernetTSPLClient: IsConnected should be false")
	}

	err := client.Connect()
	defer client.Disconnect()

	if err != nil {
		t.Fatalf("EthernetTSPLClient: Connect() error: %s", err.Error())
	}

	if !client.IsConnected() {
		t.Fatal("EthernetTSPLClient: IsConnected should be true")
	}

	data := []byte("TEST")
	data = append(data, byte(0x0A))
	err = client.SendData(data)

	if err != nil {
		t.Fatalf("EthernetTSPLClient: SendData() error: %s", err.Error())
	}

	buf := make([]byte, 11)
	err = client.ReadData(buf)
	if err != nil {
		t.Fatalf("EthernetTSPLClient: ReadData() error: %s", err.Error())
	}

	if bytes.Compare(buf[:len(buf)-1], []byte("RCVD: TEST")) != 0 {
		t.Fatalf("ReadData() got = %v, want %v", string(buf), "RCVD: TEST")
	}

	seq := gotspl.NewTSPLLabel()
	seq = seq.Cmd(gotspl.EndCmd())

	err = client.SendCommandSequence(seq)

	if err != nil {
		t.Fatalf("EthernetTSPLClient: SendCommandSequence() error: %s", err.Error())
	}

	buf2 := make([]byte, 10)
	err = client.ReadData(buf2)
	if err != nil {
		t.Fatalf("EthernetTSPLClient: ReadData() error: %s", err.Error())
	}

	if bytes.Compare(buf2[:len(buf2)-1], []byte("RCVD: END")) != 0 {
		t.Fatalf("SendCommandSequence() got = %v, want %v", string(buf2), "RCVD: END")
	}

}
