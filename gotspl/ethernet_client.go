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
