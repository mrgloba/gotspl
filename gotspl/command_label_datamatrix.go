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
	DATAMATRIX_NAME     = "DMATRIX"
	DATAMATRIX_ROWS_MIN = 10
	DATAMATRIX_ROWS_MAX = 144
	DATAMATRIX_COLS_MIN = 10
	DATAMATRIX_COLS_MAX = 144
)

type DataMatrixImpl struct {
	xCoordinate             *int
	yCoordinate             *int
	width                   *int
	height                  *int
	escapeSequenceCharacter *int
	moduleSize              *int
	rotation                *int
	isRectangle             *bool
	numberCols              *int
	numberRows              *int
	content                 *string
	contentQuote			bool
}

type DataMatrixBuilder interface {
	TSPLCommand
	XCoordinate(x int) DataMatrixBuilder
	YCoordinate(y int) DataMatrixBuilder
	Width(width int) DataMatrixBuilder
	Height(height int) DataMatrixBuilder
	EscapeSequenceCharacter(chr int) DataMatrixBuilder
	ModuleSize(size int) DataMatrixBuilder
	Rotation(angle int) DataMatrixBuilder
	IsRectangle(rectangle bool) DataMatrixBuilder
	NumberCols(cols int) DataMatrixBuilder
	NumberRows(rows int) DataMatrixBuilder
	Content(content string, quote bool) DataMatrixBuilder
}

func DataMatrixCmd() DataMatrixBuilder {
	return DataMatrixImpl{}
}

func (d DataMatrixImpl) GetMessage() ([]byte, error) {
	if d.xCoordinate == nil || d.yCoordinate == nil || d.width == nil || d.height == nil || d.content == nil {
		return nil, errors.New("ParseError DATAMATRIX Command: " +
			"xCoordinate, yCoordinate, width, height and content should be specified")
	}

	if d.escapeSequenceCharacter != nil {
		if !(*d.escapeSequenceCharacter >= 0 && *d.escapeSequenceCharacter <= 255) {
			return nil, errors.New("ParseError DATAMATRIX Command: " +
				"escapeSequenceCharacter parameter must be between 0 and 255")
		}
	}

	if d.rotation != nil {
		if !findIntInSlice(ROTATION_ANGLES, *d.rotation) {

			var err_str string

			for _, v := range ROTATION_ANGLES {
				err_str += strconv.Itoa(v)
				err_str += ","
			}
			return nil, errors.New("ParseError DATAMATRIX Command: " +
				"rotation must be one of [" + err_str[:len(err_str)-1] + "]")
		}
	}

	if d.numberRows != nil {
		if !(*d.numberRows >= DATAMATRIX_ROWS_MIN && *d.numberRows <= DATAMATRIX_ROWS_MAX) {
			return nil, errors.New(fmt.Sprintf("ParseError DATAMATRIX Command: "+
				"numberRows parameter must be between %d and %d", DATAMATRIX_ROWS_MIN, DATAMATRIX_ROWS_MAX))
		}
	}

	if d.numberCols != nil {
		if !(*d.numberCols >= DATAMATRIX_COLS_MIN && *d.numberCols <= DATAMATRIX_COLS_MAX) {
			return nil, errors.New(fmt.Sprintf("ParseError DATAMATRIX Command: "+
				"numberCols parameter must be between %d and %d", DATAMATRIX_COLS_MIN, DATAMATRIX_COLS_MAX))
		}
	}

	if d.content == nil || len(*d.content) == 0 {
		return nil, errors.New("ParseError DATAMATRIX Command: content should be specified")
	}

	buf := bytes.NewBufferString(DATAMATRIX_NAME)
	buf.WriteString(EMPTY_SPACE)

	buf.WriteString(strconv.Itoa(*d.xCoordinate))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*d.yCoordinate))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*d.width))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*d.height))
	buf.WriteString(VALUE_SEPARATOR)

	if d.escapeSequenceCharacter != nil {
		buf.WriteString("c" + strconv.Itoa(*d.escapeSequenceCharacter))
		buf.WriteString(VALUE_SEPARATOR)
	}

	if d.moduleSize != nil {
		buf.WriteString("x" + strconv.Itoa(*d.moduleSize))
		buf.WriteString(VALUE_SEPARATOR)
	}

	if d.rotation != nil {
		buf.WriteString("r" + strconv.Itoa(*d.rotation))
		buf.WriteString(VALUE_SEPARATOR)
	}

	if d.isRectangle != nil {
		buf.WriteString("a")
		if *d.isRectangle {
			buf.WriteString("1")
		} else {
			buf.WriteString("0")
		}
		buf.WriteString(VALUE_SEPARATOR)
	}

	if d.numberRows != nil {
		buf.WriteString(strconv.Itoa(*d.numberRows))
		buf.WriteString(VALUE_SEPARATOR)
	}

	if d.numberCols != nil {
		buf.WriteString(strconv.Itoa(*d.numberCols))
		buf.WriteString(VALUE_SEPARATOR)
	}

	if d.contentQuote { buf.WriteString(DOUBLE_QUOTE) }
	buf.WriteString(*d.content)
	if d.contentQuote { buf.WriteString(DOUBLE_QUOTE) }
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (d DataMatrixImpl) XCoordinate(x int) DataMatrixBuilder {
	if d.xCoordinate == nil {
		d.xCoordinate = new(int)
	}
	*d.xCoordinate = x
	return d
}

func (d DataMatrixImpl) YCoordinate(y int) DataMatrixBuilder {
	if d.yCoordinate == nil {
		d.yCoordinate = new(int)
	}
	*d.yCoordinate = y
	return d
}

func (d DataMatrixImpl) Width(width int) DataMatrixBuilder {
	if d.width == nil {
		d.width = new(int)
	}
	*d.width = width
	return d
}

func (d DataMatrixImpl) Height(height int) DataMatrixBuilder {
	if d.height == nil {
		d.height = new(int)
	}
	*d.height = height
	return d
}

func (d DataMatrixImpl) EscapeSequenceCharacter(chr int) DataMatrixBuilder {
	if d.escapeSequenceCharacter == nil {
		d.escapeSequenceCharacter = new(int)
	}
	*d.escapeSequenceCharacter = chr
	return d
}

func (d DataMatrixImpl) ModuleSize(size int) DataMatrixBuilder {
	if d.moduleSize == nil {
		d.moduleSize = new(int)
	}
	*d.moduleSize = size
	return d
}

func (d DataMatrixImpl) Rotation(angle int) DataMatrixBuilder {
	if d.rotation == nil {
		d.rotation = new(int)
	}
	*d.rotation = angle
	return d
}

func (d DataMatrixImpl) IsRectangle(rectangle bool) DataMatrixBuilder {
	if d.isRectangle == nil {
		d.isRectangle = new(bool)
	}
	*d.isRectangle = rectangle
	return d
}

func (d DataMatrixImpl) NumberCols(cols int) DataMatrixBuilder {
	if d.numberCols == nil {
		d.numberCols = new(int)
	}
	*d.numberCols = cols
	return d
}

func (d DataMatrixImpl) NumberRows(rows int) DataMatrixBuilder {
	if d.numberRows == nil {
		d.numberRows = new(int)
	}
	*d.numberRows = rows
	return d
}

func (d DataMatrixImpl) Content(content string, quote bool) DataMatrixBuilder {
	if d.content == nil {
		d.content = new(string)
	}
	*d.content = content
	d.contentQuote = quote
	return d
}
