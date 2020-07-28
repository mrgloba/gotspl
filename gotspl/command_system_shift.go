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
	SHIFT_NAME = "SHIFT"
	SHIFT_MIN  = -300
	SHIFT_MAX  = -300
)

type ShiftImpl struct {
	horizontal *int
	vertical   *int
}

type ShiftBuilder interface {
	TSPLCommand
	Horizontal(horizontal int) ShiftBuilder
	Vertical(vertical int) ShiftBuilder
}

func ShiftCmd() ShiftBuilder {
	return ShiftImpl{}
}

func (s ShiftImpl) GetMessage() ([]byte, error) {
	if s.vertical == nil {
		return nil, errors.New("ParseError SHIFT Command: vertical should be specified")
	}

	if !(*s.vertical >= SHIFT_MIN && *s.vertical <= SHIFT_MAX) {
		return nil, errors.New("ParseError SHIFT Command: vertical parameter must be between " +
			strconv.Itoa(SHIFT_MIN) + " and " + strconv.Itoa(SHIFT_MAX))
	}

	if !(s.horizontal != nil && (*s.horizontal >= SHIFT_MIN && *s.horizontal <= SHIFT_MAX)) {
		return nil, errors.New("ParseError SHIFT Command: horizontal parameter must be between " +
			strconv.Itoa(SHIFT_MIN) + " and " + strconv.Itoa(SHIFT_MAX))
	}

	buf := bytes.NewBufferString(SHIFT_NAME)
	buf.WriteString(EMPTY_SPACE)

	if s.horizontal != nil {
		buf.WriteString(strconv.Itoa(*s.horizontal))
		buf.WriteString(VALUE_SEPARATOR)
	}

	buf.WriteString(strconv.Itoa(*s.vertical))
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (s ShiftImpl) Horizontal(horizontal int) ShiftBuilder {
	if s.horizontal == nil {
		s.horizontal = new(int)
	}
	*s.horizontal = horizontal
	return s
}

func (s ShiftImpl) Vertical(vertical int) ShiftBuilder {
	if s.vertical == nil {
		s.vertical = new(int)
	}
	*s.vertical = vertical
	return s
}
