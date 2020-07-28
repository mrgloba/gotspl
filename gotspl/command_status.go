package gotspl

import "bytes"

const (
	PAUSE           StatusCommand = "P"
	CANCEL_PAUSE    StatusCommand = "O"
	CANCEL_PRINTING StatusCommand = "."
)

type StatusCommand string

func (s StatusCommand) GetMessage() ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteByte(ESC)
	buf.WriteString(STATUS_COMMAND_PREFIX)
	buf.WriteString(string(s))
	return buf.Bytes(), nil
}
