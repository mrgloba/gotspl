package gotspl

import (
	"bytes"
)

const (
	RESPONSE_OFF  ConfigParam = "RESPONSE OFF"
	RESPONSE_ON   ConfigParam = "RESPONSE ON"
	PEEL_ON       ConfigParam = "PEEL ON"
	PEEL_OFF      ConfigParam = "PEEL ON"
	BACK_ON       ConfigParam = "BACK ON"
	BACK_OFF      ConfigParam = "BACK OFF"
	TEAR_ON       ConfigParam = "TEAR ON"
	TEAR_OFF      ConfigParam = "TEAR OFF"
	STRIPER_ON    ConfigParam = "STRIPER ON"
	STRIPER_OFF   ConfigParam = "STRIPER OFF"
	REWIND_ON     ConfigParam = "REWIND ON"
	REWIND_OFF    ConfigParam = "REWIND OFF"
	REWIND_RS232  ConfigParam = "REWIND RS232"
	BLINE_REVERSE ConfigParam = "BLINE REVERSE"
	BLINE_OBVERSE ConfigParam = "BLINE OBVERSE"
)

type ConfigParam string

func (r ConfigParam) GetMessage() ([]byte, error) {
	buf := bytes.NewBufferString("SET")
	buf.WriteString(EMPTY_SPACE)
	buf.WriteString(string(r))
	buf.Write(LINE_ENDING_BYTES)

	return buf.Bytes(), nil
}
