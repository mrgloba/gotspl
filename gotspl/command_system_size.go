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
