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

const (
	DIRECTION_NAME = "DIRECTION"
)

type DirectionImpl struct {
	direction *bool
	mirror    *bool
}

type DirectionBuilder interface {
	TSPLCommand
	Direction(direction bool) DirectionBuilder
	Mirror(mirror bool) DirectionBuilder
}

func DirectionCmd() DirectionBuilder {
	return DirectionImpl{}
}

func (d DirectionImpl) GetMessage() ([]byte, error) {
	if d.direction == nil {
		return nil, errors.New("ParseError DIRECTION Command: direction should be specified")
	}

	buf := bytes.NewBufferString(DIRECTION_NAME)
	buf.WriteString(EMPTY_SPACE)
	if *d.direction {
		buf.WriteString("1")
	} else {
		buf.WriteString("0")
	}

	if d.mirror != nil {
		buf.WriteString(VALUE_SEPARATOR)
		if *d.mirror {
			buf.WriteString("1")
		} else {
			buf.WriteString("0")
		}
	}

	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (d DirectionImpl) Direction(direction bool) DirectionBuilder {
	if d.direction == nil {
		d.direction = new(bool)
	}
	*d.direction = direction
	return d
}

func (d DirectionImpl) Mirror(mirror bool) DirectionBuilder {
	if d.mirror == nil {
		d.mirror = new(bool)
	}
	*d.mirror = mirror
	return d
}
