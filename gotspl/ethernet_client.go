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

import "net"

const NET_TCP_NETWORK = "tcp"

type EthernetTSPLClient struct {
	address     string
	isConnected bool
	tcpclient   net.Conn
}

func NewEthernetTSPLClient(address string) TSPLClient {
	return &EthernetTSPLClient{address: address}
}

func (c *EthernetTSPLClient) Connect() error {
	var err error = nil
	c.tcpclient, err = net.Dial(NET_TCP_NETWORK, c.address)
	if err == nil {
		c.isConnected = true
	}
	return err
}

func (c *EthernetTSPLClient) Disconnect() error {
	c.isConnected = false
	return c.tcpclient.Close()
}

func (c *EthernetTSPLClient) IsConnected() bool {
	return c.isConnected
}

func (c *EthernetTSPLClient) SendData(data []byte) error {
	_, err := c.tcpclient.Write(data)
	return err
}

func (c *EthernetTSPLClient) ReadData(data []byte) error {
	_, err := c.tcpclient.Read(data)
	return err
}

func (c *EthernetTSPLClient) SendCommandSequence(commandSequence TSPLCommandSequence) error {
	seq, err := commandSequence.getTSPLCode()
	if err != nil {
		return err
	}

	return c.SendData(seq)
}
