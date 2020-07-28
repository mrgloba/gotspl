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
	DENSITY_NAME         = "DENSITY"
	DENSITY_DARKNESS_MIN = 0
	DENSITY_DARKNESS_MAX = 15
)

type DensityImpl struct {
	darkness *int
}

type DensityBuilder interface {
	TSPLCommand
	Darkness(darkness int) DensityBuilder
}

func DensityCmd() DensityBuilder {
	return DensityImpl{}
}

func (di DensityImpl) Darkness(darkness int) DensityBuilder {
	if di.darkness == nil {
		di.darkness = new(int)
	}
	*di.darkness = darkness
	return di
}

func (di DensityImpl) GetMessage() ([]byte, error) {
	if di.darkness == nil {
		return nil, errors.New("ParseError DENSITY Command: darkness should be specified")
	}

	if !(*di.darkness >= DENSITY_DARKNESS_MIN || *di.darkness <= DENSITY_DARKNESS_MAX) {
		return nil, errors.New("ParseError DENSITY Command: darkness parameter must be between " +
			strconv.Itoa(DENSITY_DARKNESS_MIN) + " and " + strconv.Itoa(DENSITY_DARKNESS_MAX))
	}

	buf := bytes.NewBufferString(DENSITY_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.WriteString(strconv.Itoa(*di.darkness))
	buf.Write(LINE_ENDING_BYTES)

	return buf.Bytes(), nil
}
