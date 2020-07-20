package gotspl

import "bytes"

type TSPLLabel struct {
	commandList []TSPLCommand
}

type TSPLLabelBuilder interface {
	TSPLCommandSequence
	Cmd(command TSPLCommand) TSPLLabelBuilder
}

func NewTSPLLabel() TSPLLabelBuilder {
	return TSPLLabel{}
}

func (T TSPLLabel) getTSPLCode() ([]byte, error) {
	var buf bytes.Buffer
	for _, c := range T.commandList {
		msg, err := c.GetMessage()
		if err != nil {
			return nil, err
		}

		buf.Write(msg)
	}

	return buf.Bytes(), nil
}

func (T TSPLLabel) Cmd(command TSPLCommand) TSPLLabelBuilder {
	if command != nil {
		T.commandList = append(T.commandList, command)
	}

	return T
}
