package gotspl

import (
	"bytes"
	"errors"
)

const GAP_NAME = "GAP"

type GapImpl struct {
	labelDistance       *float64
	labelOffsetDistance *float64
}

type GapBuilder interface {
	TSPLCommand
	LabelDistance(labelDistance float64) GapBuilder
	LabelOffsetDistance(labelOffsetDistance float64) GapBuilder
}

func GapCmd() GapBuilder {
	return GapImpl{}
}

func (gi GapImpl) LabelDistance(labelDistance float64) GapBuilder {
	if gi.labelDistance == nil {
		gi.labelDistance = new(float64)
	}
	*gi.labelDistance = labelDistance
	return gi
}

func (gi GapImpl) LabelOffsetDistance(labelOffsetDistance float64) GapBuilder {
	if gi.labelOffsetDistance == nil {
		gi.labelOffsetDistance = new(float64)
	}
	*gi.labelOffsetDistance = labelOffsetDistance
	return gi
}

func (gi GapImpl) GetMessage() ([]byte, error) {
	if gi.labelDistance == nil || gi.labelOffsetDistance == nil {
		return nil, errors.New("ParseError GAP Command: labelDistance and labelOffsetDistance should be specified")
	}

	buf := bytes.NewBufferString(GAP_NAME)

	buf.WriteString(EMPTY_SPACE)

	buf.Write(formatFloatWithUnits(*gi.labelDistance, true))
	buf.WriteString(VALUE_SEPARATOR)
	buf.Write(formatFloatWithUnits(*gi.labelOffsetDistance, true))
	buf.Write(LINE_ENDING_BYTES)

	return buf.Bytes(), nil
}
