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

const SPEED_NAME = "SPEED"

type SpeedImpl struct {
	printSpeed *float64
}

type SpeedBuilder interface {
	TSPLCommand
	PrintSpeed(printSpeed float64) SpeedBuilder
}

func (s SpeedImpl) GetMessage() ([]byte, error) {
	if s.printSpeed == nil {
		return nil, errors.New("ParseError SOUND Command: PrintSpeed should be specified")
	}

	buf := bytes.NewBufferString(SPEED_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.Write(formatFloatWithUnits(*s.printSpeed, false))
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (s SpeedImpl) PrintSpeed(printSpeed float64) SpeedBuilder {
	if s.printSpeed == nil {
		s.printSpeed = new(float64)
	}
	*s.printSpeed = printSpeed
	return s
}

func SpeedCmd() SpeedBuilder {
	return SpeedImpl{}
}
