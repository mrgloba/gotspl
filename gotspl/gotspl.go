package gotspl

type MeasurementSystem int

var measurementSystem MeasurementSystem = MEASUREMENT_SYSTEM_ENGLISH

func TSPLInitialize(mSystem MeasurementSystem) {
	measurementSystem = mSystem
}

func NewTcpTSPLClient(address string, measurementSystem MeasurementSystem) TSPLClient {
	return &EthernetTSPLClient{address: address}
}
