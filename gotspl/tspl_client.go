package gotspl

type TSPLClient interface {
	Connect() error
	Disconnect() error
	SendData(data []byte) error
	ReadData(data []byte) error
	SendCommandSequence(commandSequence TSPLCommandSequence) error
	IsConnected() bool
}
