package gotspl

import "bytes"

const END_NAME = "END"

type EndImpl struct {
}

type EndBuilder interface {
	TSPLCommand
}

func EndCmd() EndBuilder {
	return EndImpl{}
}

func (c EndImpl) GetMessage() ([]byte, error) {
	buf := bytes.NewBufferString(END_NAME)
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}
