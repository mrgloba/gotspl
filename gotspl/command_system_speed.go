package gotspl

import (
	"bytes"
	"errors"
)

const SPEED_NAME = "SPEED"

type SpeedImpl struct {
	printSpeed *float64
}

type SpeedBuilder interface {
	TSPLCommand
	PrintSpeed(printSpeed float64) SpeedBuilder
}

func (s SpeedImpl) GetMessage() ([]byte, error) {
	if s.printSpeed == nil {
		return nil, errors.New("ParseError SOUND Command: PrintSpeed should be specified")
	}

	buf := bytes.NewBufferString(SPEED_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.Write(formatFloatWithUnits(*s.printSpeed, false))
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (s SpeedImpl) PrintSpeed(printSpeed float64) SpeedBuilder {
	if s.printSpeed == nil {
		s.printSpeed = new(float64)
	}
	*s.printSpeed = printSpeed
	return s
}

func SpeedCmd() SpeedBuilder {
	return SpeedImpl{}
}
