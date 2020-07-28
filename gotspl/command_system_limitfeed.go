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
	LIMITFEED_NAME = "LIMITFEED"
)

type LimitFeedImpl struct {
	maxLenSensorDetect *float64
	maxLenPaper        *float64
	maxLenGap          *float64
}

type LimitFeedBuilder interface {
	TSPLCommand
	LengthSensorDetect(length float64) LimitFeedBuilder
	LengthPaper(length float64) LimitFeedBuilder
	LengthGap(length float64) LimitFeedBuilder
}

func LimitFeedCmd() LimitFeedBuilder {
	return LimitFeedImpl{}
}

func (l LimitFeedImpl) GetMessage() ([]byte, error) {
	if l.maxLenSensorDetect == nil {
		return nil, errors.New("ParseError LIMITFEED Command: LengthSensorDetect should be specified")
	}

	buf := bytes.NewBufferString(LIMITFEED_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.Write(formatFloatWithUnits(*l.maxLenSensorDetect, true))

	if l.maxLenPaper != nil && l.maxLenGap != nil {
		buf.WriteString(VALUE_SEPARATOR)
		buf.Write(formatFloatWithUnits(*l.maxLenPaper, true))
		buf.WriteString(VALUE_SEPARATOR)
		buf.Write(formatFloatWithUnits(*l.maxLenGap, true))
	}

	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (l LimitFeedImpl) LengthSensorDetect(length float64) LimitFeedBuilder {
	if l.maxLenSensorDetect == nil {
		l.maxLenSensorDetect = new(float64)
	}
	*l.maxLenSensorDetect = length
	return l
}

func (l LimitFeedImpl) LengthPaper(length float64) LimitFeedBuilder {
	if l.maxLenPaper == nil {
		l.maxLenPaper = new(float64)
	}
	*l.maxLenPaper = length
	return l
}

func (l LimitFeedImpl) LengthGap(length float64) LimitFeedBuilder {
	if l.maxLenGap == nil {
		l.maxLenGap = new(float64)
	}
	*l.maxLenGap = length
	return l
}
