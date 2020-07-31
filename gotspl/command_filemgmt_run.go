package gotspl

import (
	"bytes"
	"errors"
	"strings"
)

const RUN_NAME = "RUN"

type RunImpl struct {
	file *string
}

type RunBuilder interface {
	TSPLCommand
	File(name string) RunBuilder
}

func RunCmd() RunBuilder {
	return RunImpl{}
}

func (r RunImpl) GetMessage() ([]byte, error) {
	if r.file == nil || len(*r.file) == 0 {
		return nil, errors.New("ParseError RUN Command: file should be specified")
	}

	buf := bytes.NewBufferString(RUN_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.WriteString(DOUBLE_QUOTE)
	buf.WriteString(strings.ToUpper(*r.file))
	buf.WriteString(DOUBLE_QUOTE)
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (r RunImpl) File(name string) RunBuilder {
	if r.file == nil {
		r.file = new(string)
	}
	*r.file = name
	return r
}
