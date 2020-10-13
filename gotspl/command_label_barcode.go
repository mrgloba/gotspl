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
	BARCODE_NAME = "BARCODE"
	BARCODE_HUMAN_READABLE_MIN = 0
	BARCODE_HUMAN_READABLE_MAX = 3
	BARCODE_ALIGNMENT_MIN = 0
	BARCODE_ALIGNMENT_MAX = 3
)

type BarcodeImpl struct {
	xCoordinate             *int
	yCoordinate             *int
	codeType				*string
	height					*int
	humanReadable			*int
	rotation                *int
	narrow					*int
	wide					*int
	alignment				*int
	content                 *string
	contentQuote			bool
}

type BarcodeBuilder interface {
	TSPLCommand
	XCoordinate(x int) BarcodeBuilder
	YCoordinate(y int) BarcodeBuilder
	Height(height int) BarcodeBuilder
	CodeType(codeType string) BarcodeBuilder
	HumanReadable(humanReadable int) BarcodeBuilder
	Rotation(angle int) BarcodeBuilder
	Narrow( narrow int) BarcodeBuilder
	Wide(wide int) BarcodeBuilder
	Alignment(align int) BarcodeBuilder
	Content(content string, quote bool) BarcodeBuilder
}


func BarcodeCmd() BarcodeBuilder {
	return BarcodeImpl{}
}

func (b BarcodeImpl) GetMessage() ([]byte, error) {
	if b.xCoordinate == nil || b.yCoordinate == nil || b.codeType == nil ||
		b.humanReadable == nil || b.rotation == nil || b.narrow == nil || b.wide == nil ||
		b.height == nil || b.content == nil {
		return nil, errors.New("ParseError BARCODE Command: " +
			"xCoordinate, yCoordinate, codeType, height, humanReadable, rotation, narrow, wide  and content should be specified")
	}

	//TODO: check codeType by manual

	if b.rotation != nil {
		if !findIntInSlice(ROTATION_ANGLES, *b.rotation) {

			var err_str string

			for _, v := range ROTATION_ANGLES {
				err_str += strconv.Itoa(v)
				err_str += ","
			}
			return nil, errors.New("ParseError BARCODE Command: " +
				"rotation must be one of [" + err_str[:len(err_str)-1] + "]")
		}
	}

	if b.humanReadable != nil {
		if !(*b.humanReadable >= BARCODE_HUMAN_READABLE_MIN && *b.humanReadable <= BARCODE_HUMAN_READABLE_MAX) {
			return nil, errors.New(fmt.Sprintf("ParseError BARCODE Command: "+
				"humanReadable parameter must be between %d and %d", BARCODE_HUMAN_READABLE_MIN, BARCODE_HUMAN_READABLE_MAX))
		}
	}

	if b.alignment != nil {
		if !(*b.alignment >= BARCODE_ALIGNMENT_MIN && *b.humanReadable <= BARCODE_ALIGNMENT_MAX) {
			return nil, errors.New(fmt.Sprintf("ParseError BARCODE Command: "+
				"alignment parameter must be between %d and %d", BARCODE_ALIGNMENT_MIN, BARCODE_ALIGNMENT_MAX))
		}
	}


	if b.content == nil || len(*b.content) == 0 {
		return nil, errors.New("ParseError BARCODE Command: content should be specified")
	}

	buf := bytes.NewBufferString(BARCODE_NAME)
	buf.WriteString(EMPTY_SPACE)

	buf.WriteString(strconv.Itoa(*b.xCoordinate))
	buf.WriteString(VALUE_SEPARATOR)
	buf.WriteString(strconv.Itoa(*b.yCoordinate))
	buf.WriteString(VALUE_SEPARATOR)

	buf.WriteString(DOUBLE_QUOTE)
	buf.WriteString(*b.codeType)
	buf.WriteString(DOUBLE_QUOTE)
	buf.WriteString(VALUE_SEPARATOR)

	buf.WriteString(strconv.Itoa(*b.height))
	buf.WriteString(VALUE_SEPARATOR)

	buf.WriteString(strconv.Itoa(*b.humanReadable))
	buf.WriteString(VALUE_SEPARATOR)

	buf.WriteString(strconv.Itoa(*b.rotation))
	buf.WriteString(VALUE_SEPARATOR)

	buf.WriteString(strconv.Itoa(*b.narrow))
	buf.WriteString(VALUE_SEPARATOR)

	buf.WriteString(strconv.Itoa(*b.wide))
	buf.WriteString(VALUE_SEPARATOR)

	if b.alignment != nil {
		buf.WriteString(strconv.Itoa(*b.alignment))
		buf.WriteString(VALUE_SEPARATOR)
	}

	buf.WriteString(EMPTY_SPACE)

	if b.contentQuote { buf.WriteString(DOUBLE_QUOTE) }
	buf.WriteString(*b.content)
	if b.contentQuote { buf.WriteString(DOUBLE_QUOTE) }
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (b BarcodeImpl) XCoordinate(x int) BarcodeBuilder {
	if b.xCoordinate == nil {
		b.xCoordinate = new(int)
	}
	*b.xCoordinate = x
	return b
}

func (b BarcodeImpl) YCoordinate(y int) BarcodeBuilder {
	if b.yCoordinate == nil {
		b.yCoordinate = new(int)
	}
	*b.yCoordinate = y
	return b
}


func (b BarcodeImpl) Height(height int) BarcodeBuilder {
	if b.height == nil {
		b.height = new(int)
	}
	*b.height = height
	return b
}

func (b BarcodeImpl) Rotation(angle int) BarcodeBuilder {
	if b.rotation == nil {
		b.rotation = new(int)
	}
	*b.rotation = angle
	return b
}

func (b BarcodeImpl) Content(content string, quote bool) BarcodeBuilder {
	if b.content == nil {
		b.content = new(string)
	}
	*b.content = content
	b.contentQuote = quote
	return b
}

func (b BarcodeImpl) CodeType(codeType string) BarcodeBuilder {
	if b.codeType == nil {
		b.codeType = new(string)
	}
	*b.codeType = codeType
	return b
}

func (b BarcodeImpl) HumanReadable(humanReadable int) BarcodeBuilder {
	if b.humanReadable == nil {
		b.humanReadable = new(int)
	}
	*b.humanReadable = humanReadable
	return b
}

func (b BarcodeImpl) Narrow(narrow int) BarcodeBuilder {
	if b.narrow == nil {
		b.narrow = new(int)
	}
	*b.narrow = narrow
	return b
}

func (b BarcodeImpl) Wide(wide int) BarcodeBuilder {
	if b.wide == nil {
		b.wide = new(int)
	}
	*b.wide = wide
	return b
}

func (b BarcodeImpl) Alignment(align int) BarcodeBuilder {
	if b.alignment == nil {
		b.alignment = new(int)
	}
	*b.alignment = align
	return b
}
