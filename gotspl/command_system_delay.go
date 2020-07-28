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
	"strconv"
)

const (
	DELAY_NAME = "DELAY"
)

type DelayImpl struct {
	duration *int
}

type DelayBuilder interface {
	TSPLCommand
	Duration(duration int) DelayBuilder
}

func DelayCmd() DelayBuilder {
	return DelayImpl{}
}

func (d DelayImpl) GetMessage() ([]byte, error) {
	if d.duration == nil {
		return nil, errors.New("ParseError DELAY Command: duration should specified")
	}

	buf := bytes.NewBufferString(DELAY_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.WriteString(strconv.Itoa(*d.duration))
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (d DelayImpl) Duration(duration int) DelayBuilder {
	if d.duration == nil {
		d.duration = new(int)
	}
	*d.duration = duration
	return d
}
