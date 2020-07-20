/*
 * Copyright 2020 Anton Globa
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
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
