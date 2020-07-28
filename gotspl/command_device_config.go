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
)

const (
	RESPONSE_OFF   ConfigParam = "RESPONSE OFF"
	RESPONSE_ON    ConfigParam = "RESPONSE ON"
	PEEL_ON        ConfigParam = "PEEL ON"
	PEEL_OFF       ConfigParam = "PEEL ON"
	BACK_ON        ConfigParam = "BACK ON"
	BACK_OFF       ConfigParam = "BACK OFF"
	TEAR_ON        ConfigParam = "TEAR ON"
	TEAR_OFF       ConfigParam = "TEAR OFF"
	STRIPER_ON     ConfigParam = "STRIPER ON"
	STRIPER_OFF    ConfigParam = "STRIPER OFF"
	REWIND_ON      ConfigParam = "REWIND ON"
	REWIND_OFF     ConfigParam = "REWIND OFF"
	REWIND_RS232   ConfigParam = "REWIND RS232"
	BLINE_REVERSE  ConfigParam = "BLINE REVERSE"
	BLINE_OBVERSE  ConfigParam = "BLINE OBVERSE"
	HEAD_OFF       ConfigParam = "HEAD OFF"
	HEAD_ON        ConfigParam = "HEAD ON"
	RIBBON_ON      ConfigParam = "RIBBON ON"
	RIBBON_OFF     ConfigParam = "RIBBON OFF"
	RIBBON_INSIDE  ConfigParam = "RIBBON INSIDE"
	RIBBON_OUTSIDE ConfigParam = "RIBBON OUTSIDE"
)

type ConfigParam string

func (r ConfigParam) GetMessage() ([]byte, error) {
	buf := bytes.NewBufferString("SET")
	buf.WriteString(EMPTY_SPACE)
	buf.WriteString(string(r))
	buf.Write(LINE_ENDING_BYTES)

	return buf.Bytes(), nil
}
