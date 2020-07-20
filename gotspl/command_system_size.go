package gotspl

import (
	"bytes"
	"errors"
)

const (
	SIZE_NAME = "SIZE"
)

type SizeImpl struct {
	labelWidth  *float64
	labelLength *float64
}

type SizeBuilder interface {
	TSPLCommand
	LabelWidth(labelWidth float64) SizeBuilder
	LabelLength(labelLength float64) SizeBuilder
}

func SizeCmd() SizeBuilder {
	return SizeImpl{}
}

func (si SizeImpl) LabelWidth(labelWidth float64) SizeBuilder {
	if si.labelWidth == nil {
		si.labelWidth = new(float64)
	}
	*si.labelWidth = labelWidth
	return si
}

func (si SizeImpl) LabelLength(labelLength float64) SizeBuilder {
	if si.labelLength == nil {
		si.labelLength = new(float64)
	}
	*si.labelLength = labelLength
	return si
}

func (si SizeImpl) GetMessage() ([]byte, error) {
	if si.labelWidth == nil {
		return nil, errors.New("ParseError SIZE Command: LabelWidth should be specified")
	}

	buf := bytes.NewBufferString(SIZE_NAME)

	buf.WriteString(EMPTY_SPACE)

	buf.Write(formatFloatWithUnits(*si.labelWidth, true))

	if si.labelLength != nil {
		buf.WriteString(VALUE_SEPARATOR)
		buf.Write(formatFloatWithUnits(*si.labelLength, true))
	}

	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}
