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

const BLINE_NAME = "BLINE"

type BlineImpl struct {
	lineHeight    *float64
	feedingLength *float64
}

type BlineBuilder interface {
	TSPLCommand
	LineHeight(lineHeight float64) BlineBuilder
	FeedingLength(feedingLength float64) BlineBuilder
}

func BlineCmd() BlineBuilder {
	return BlineImpl{}
}

func (gi BlineImpl) LineHeight(lineHeight float64) BlineBuilder {
	if gi.lineHeight == nil {
		gi.lineHeight = new(float64)
	}
	*gi.lineHeight = lineHeight
	return gi
}

func (gi BlineImpl) FeedingLength(feedingLength float64) BlineBuilder {
	if gi.feedingLength == nil {
		gi.feedingLength = new(float64)
	}
	*gi.feedingLength = feedingLength
	return gi
}

func (gi BlineImpl) GetMessage() ([]byte, error) {
	if gi.lineHeight == nil || gi.feedingLength == nil {
		return nil, errors.New("ParseError BLINE Command: lineHeight and feedingLength should be specified")
	}

	buf := bytes.NewBufferString(BLINE_NAME)

	buf.WriteString(EMPTY_SPACE)

	buf.Write(formatFloatWithUnits(*gi.lineHeight, true))
	buf.WriteString(VALUE_SEPARATOR)
	buf.Write(formatFloatWithUnits(*gi.feedingLength, true))
	buf.Write(LINE_ENDING_BYTES)

	return buf.Bytes(), nil
}
