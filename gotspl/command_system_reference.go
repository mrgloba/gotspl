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
	REFERENCE_NAME = "REFERENCE"
)

type ReferenceImpl struct {
	horizontal *int
	vertical   *int
}

type ReferenceBuilder interface {
	TSPLCommand
	Horizontal(horizontal int) ReferenceBuilder
	Vertical(vertical int) ReferenceBuilder
}

func ReferenceCmd() ReferenceBuilder {
	return ReferenceImpl{}
}

func (r ReferenceImpl) GetMessage() ([]byte, error) {
	if r.horizontal == nil || r.vertical == nil {
		return nil, errors.New("ParseError REFERENCE Command: horizontal and vertical should be specified")
	}

	buf := bytes.NewBufferString(REFERENCE_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.WriteString(strconv.Itoa(*r.horizontal))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*r.vertical))
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (r ReferenceImpl) Horizontal(horizontal int) ReferenceBuilder {
	if r.horizontal == nil {
		r.horizontal = new(int)
	}
	*r.horizontal = horizontal
	return r
}

func (r ReferenceImpl) Vertical(vertical int) ReferenceBuilder {
	if r.vertical == nil {
		r.vertical = new(int)
	}
	*r.vertical = vertical
	return r
}
