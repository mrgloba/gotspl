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
	FEED_NAME = "FEED"
	FEED_MIN  = 1
	FEED_MAX  = 9999
)

type FeedImpl struct {
	length *int
}

type FeedBuilder interface {
	TSPLCommand
	Length(length int) FeedBuilder
}

func FeedCmd() FeedBuilder {
	return FeedImpl{}
}

func (f FeedImpl) GetMessage() ([]byte, error) {
	if f.length == nil {
		return nil, errors.New("ParseError FEED Command: length should be specified")
	}

	if !(*f.length >= FEED_MIN && *f.length <= FEED_MAX) {
		return nil, errors.New("ParseError FEED Command: length parameter must be between " +
			strconv.Itoa(FEED_MIN) + " and " + strconv.Itoa(FEED_MAX))
	}

	buf := bytes.NewBufferString(FEED_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.WriteString(strconv.Itoa(*f.length))
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (f FeedImpl) Length(length int) FeedBuilder {
	if f.length == nil {
		f.length = new(int)
	}
	*f.length = length
	return f
}
