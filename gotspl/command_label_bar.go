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
	BAR_NAME = "BAR"
)

type BarImpl struct {
	xCoordinate *int
	yCoordinate *int
	width       *int
	height      *int
}

type BarBuilder interface {
	TSPLCommand
	XCoordinate(x int) BarBuilder
	YCoordinate(y int) BarBuilder
	Width(width int) BarBuilder
	Height(height int) BarBuilder
}

func BarCmd() BarBuilder {
	return BarImpl{}
}

func (b BarImpl) GetMessage() ([]byte, error) {
	if b.xCoordinate == nil || b.yCoordinate == nil || b.width == nil || b.height == nil {
		return nil, errors.New("ParseError BAR Command: " +
			"xCoordinate, yCoordinate, width, height should be specified")
	}

	buf := bytes.NewBufferString(BAR_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.WriteString(strconv.Itoa(*b.xCoordinate))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*b.yCoordinate))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*b.width))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*b.height))
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (b BarImpl) XCoordinate(x int) BarBuilder {
	if b.xCoordinate == nil {
		b.xCoordinate = new(int)
	}
	*b.xCoordinate = x
	return b
}

func (b BarImpl) YCoordinate(y int) BarBuilder {
	if b.yCoordinate == nil {
		b.yCoordinate = new(int)
	}
	*b.yCoordinate = y
	return b
}

func (b BarImpl) Width(width int) BarBuilder {
	if b.width == nil {
		b.width = new(int)
	}
	*b.width = width
	return b
}

func (b BarImpl) Height(height int) BarBuilder {
	if b.height == nil {
		b.height = new(int)
	}
	*b.height = height
	return b
}
