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
	CODEPAGE_NAME = "CODEPAGE"
)

type CodePageImpl struct {
	codePage *string
}

type CodePageBuilder interface {
	TSPLCommand
	CodePage(codePage string) CodePageBuilder
}

func CodePageCmd() CodePageBuilder {
	return CodePageImpl{}
}

func (c CodePageImpl) GetMessage() ([]byte, error) {
	if c.codePage == nil {
		return nil, errors.New("ParseError CODEPAGE Command: codepage should be specified")
	}

	buf := bytes.NewBufferString(CODEPAGE_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.WriteString(*c.codePage)
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (c CodePageImpl) CodePage(codePage string) CodePageBuilder {
	if c.codePage == nil {
		c.codePage = new(string)
	}
	*c.codePage = codePage
	return c
}
