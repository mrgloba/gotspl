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
	"fmt"
	"strconv"
)

const (
	TEXT_NAME = "TEXT"

	TEXT_ALIGNMENT_DEFAILT TextAlignment = iota
	TEXT_ALIGNMENT_LEFT
	TEXT_ALIGNMENT_CENTER
	TEXT_ALIGNMENT_RIGHT

	TEXT_MULTIPLIER_MIN = 1
	TEXT_MULTIPLIER_MAX = 10
)

type TextAlignment int

type TextImpl struct {
	xCoordinate             *int
	yCoordinate             *int
	fontName				*string
	rotation                *int
	xMultiplication			*float64
	yMultiplication			*float64
	alignment				*int
	content                 *string
	contentQuote			bool
}

type TextBuilder interface {
	TSPLCommand
	XCoordinate(x int) TextBuilder
	YCoordinate(y int) TextBuilder
	FontName(name string) TextBuilder
	Rotation(angle int) TextBuilder
	XMultiplier(xm float64) TextBuilder
	YMultiplier(ym float64) TextBuilder
	Alignment(align TextAlignment) TextBuilder
	Content(content string, quote bool) TextBuilder

}

func Text() TextBuilder {
	return TextImpl{}
}

func (t TextImpl) GetMessage() ([]byte, error) {
	if t.xCoordinate == nil ||
		t.yCoordinate == nil ||
		t.fontName == nil ||
		t.rotation == nil ||
		t.content == nil ||
		t.xMultiplication == nil ||
		t.yMultiplication == nil{
		return nil, errors.New("ParseError TEXT Command: " +
			"xCoordinate, yCoordinate, fontName, rotation, xMultiplication, yMultiplication and content should be specified")
	}

	if t.rotation != nil {
		if !findIntInSlice(ROTATION_ANGLES, *t.rotation) {

			var err_str string

			for _, v := range ROTATION_ANGLES {
				err_str += strconv.Itoa(v)
				err_str += ","
			}
			return nil, errors.New("ParseError TEXT Command: " +
				"rotation must be one of [" + err_str[:len(err_str)-1] + "]")
		}
	}

	if t.xMultiplication != nil {
		if !(*t.xMultiplication >= TEXT_MULTIPLIER_MIN && *t.xMultiplication <= TEXT_MULTIPLIER_MAX) {
			return nil, errors.New(fmt.Sprintf("ParseError TEXT Command: "+
				"xMultiplication parameter must be between %d and %d", TEXT_MULTIPLIER_MIN, TEXT_MULTIPLIER_MAX))
		}
	}

	if t.yMultiplication != nil {
		if !(*t.yMultiplication >= TEXT_MULTIPLIER_MIN && *t.yMultiplication <= TEXT_MULTIPLIER_MAX) {
			return nil, errors.New(fmt.Sprintf("ParseError TEXT Command: "+
				"yMultiplication parameter must be between %d and %d", TEXT_MULTIPLIER_MIN, TEXT_MULTIPLIER_MAX))
		}
	}

	buf := bytes.NewBufferString(TEXT_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.WriteString(strconv.Itoa(*t.xCoordinate))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*t.yCoordinate))
	buf.WriteString(VALUE_SEPARATOR+DOUBLE_QUOTE)
	buf.WriteString(*t.fontName)
	buf.WriteString(DOUBLE_QUOTE+VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*t.rotation))
	buf.WriteString(VALUE_SEPARATOR)
	buf.Write(formatFloatWithUnits(*t.xMultiplication, false))
	buf.WriteString(VALUE_SEPARATOR)
	buf.Write(formatFloatWithUnits(*t.yMultiplication, false))
	buf.WriteString(VALUE_SEPARATOR)
	if t.alignment != nil {
		buf.WriteString(strconv.Itoa(*t.alignment))
		buf.WriteString(VALUE_SEPARATOR)
	}

	if t.contentQuote { buf.WriteString(DOUBLE_QUOTE) }
	buf.WriteString(*t.content)
	if t.contentQuote { buf.WriteString(DOUBLE_QUOTE) }
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (t TextImpl) XCoordinate(x int) TextBuilder {
	if t.xCoordinate == nil { t.xCoordinate = new(int) }
	*t.xCoordinate = x
	return t
}

func (t TextImpl) YCoordinate(y int) TextBuilder {
	if t.yCoordinate == nil { t.yCoordinate = new(int) }
	*t.yCoordinate = y
	return t
}

func (t TextImpl) FontName(name string) TextBuilder {
	if t.fontName == nil { t.fontName = new(string) }
	*t.fontName = name
	return t
}

func (t TextImpl) Rotation(angle int) TextBuilder {
	if t.rotation == nil { t.rotation = new(int) }
	*t.rotation = angle
	return t
}

func (t TextImpl) XMultiplier(xm float64) TextBuilder {
	if t.xMultiplication == nil { t.xMultiplication = new(float64) }
	*t.xMultiplication = xm
	return t
}

func (t TextImpl) YMultiplier(ym float64) TextBuilder {
	if t.yMultiplication == nil { t.yMultiplication = new(float64) }
	*t.yMultiplication = ym
	return t
}

func (t TextImpl) Alignment(align TextAlignment) TextBuilder {
	if t.alignment == nil { t.alignment = new(int) }
	*t.alignment = int(align)
	return t
}

func (t TextImpl) Content(content string, quote bool) TextBuilder {
	if t.content == nil { t.content = new(string) }
	*t.content = content
	t.contentQuote = quote
	return t
}
