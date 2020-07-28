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

const OFFSET_NAME = "OFFSET"

type OffSetImpl struct {
	distance *float64
}

type OffSetBuilder interface {
	TSPLCommand
	Distance(distance float64) OffSetBuilder
}

func OffSetCmd() OffSetBuilder {
	return OffSetImpl{}
}

func (os OffSetImpl) Distance(distance float64) OffSetBuilder {
	if os.distance == nil {
		os.distance = new(float64)
	}
	*os.distance = distance
	return os
}

func (os OffSetImpl) GetMessage() ([]byte, error) {
	if os.distance == nil {
		return nil, errors.New("ParseError OFFSET Command: distance should be specified")
	}

	buf := bytes.NewBufferString(OFFSET_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.Write(formatFloatWithUnits(*os.distance, true))
	buf.Write(LINE_ENDING_BYTES)

	return buf.Bytes(), nil
}
