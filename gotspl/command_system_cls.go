package gotspl

import "bytes"

const CLS_NAME = "CLS"

type ClsImpl struct {
}

type ClsBuilder interface {
	TSPLCommand
}

func ClsCmd() ClsBuilder {
	return ClsImpl{}
}

func (c ClsImpl) GetMessage() ([]byte, error) {
	buf := bytes.NewBufferString(CLS_NAME)
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}
