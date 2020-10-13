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
package gotspl

import (
	"context"
	"net"
	"sync"
	"time"
)

const NET_TCP_NETWORK = "tcp"

type EthernetTSPLClient struct {
	address     string
	isConnected bool
	tcpclient   net.Conn
	exit bool
	listeners []chan *RawResponseEvent
	wg sync.WaitGroup
}

func NewEthernetTSPLClient(address string) TSPLClient {
	return &EthernetTSPLClient{address: address}
}

func (c *EthernetTSPLClient) AddResponseListener(listener chan *RawResponseEvent) {
	c.listeners = append(c.listeners, listener)
}

func (c *EthernetTSPLClient) emitListeners(response *RawResponseEvent) {
	for _, ch := range c.listeners {
		ch <- response
	}
}

func (c *EthernetTSPLClient) Connect() error {
	c.exit = false
	var err error = nil
	c.tcpclient, err = net.Dial(NET_TCP_NETWORK, c.address)
	if err == nil {
		c.isConnected = true
	}

	c.wg.Add(1)
	go c.dataReader()
	return err
}

func (c *EthernetTSPLClient) Disconnect() error {
	c.exit = true
	c.wg.Wait()
	c.isConnected = false
	return c.tcpclient.Close()
}

func (c *EthernetTSPLClient) IsConnected() bool {
	return c.isConnected
}

func (c *EthernetTSPLClient) SendData(data []byte) error {
	c.tcpclient.SetWriteDeadline(time.Now().Add(time.Millisecond*1000))
	_, err := c.tcpclient.Write(data)

	return err
}

func (c *EthernetTSPLClient) ReadData(data []byte) (int, error) {
	c.tcpclient.SetReadDeadline(time.Now().Add(time.Millisecond*1000))
	n, err := c.tcpclient.Read(data)
	return n, err
}

func (c *EthernetTSPLClient) dataReader() {
	for {
		if c.isConnected {
			data  := make([]byte,1024)
			n, err := c.ReadData(data)
			if err != nil {
				break
			}
			c.emitListeners(&RawResponseEvent{
				Size: n,
				Data: data,
			})
		} else {
			break
		}
	}
	c.wg.Done()
}

func (c *EthernetTSPLClient) SendCommandSequence(commandSequence TSPLCommandSequence) error {
	seq, err := commandSequence.getTSPLCode()
	if err != nil {
		return err
	}

	return c.SendData(seq)
}

func (c *EthernetTSPLClient) SendCommand(command TSPLCommand) error {
	cmd, err := command.GetMessage()
	if err != nil {
		return err
	}

	return c.SendData(cmd)
}

func (c *EthernetTSPLClient) printerDataPuller(ctx context.Context, dataChan <-chan int) {
	dataChan = make(chan int)

}
