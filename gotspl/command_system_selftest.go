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

import "bytes"

const (
	SELFTEST_NAME                     = "SELFTEST"
	SELFTEST          SelfTestCommand = ""
	SELFTEST_PATTERN  SelfTestCommand = "PATTERN"
	SELFTEST_ETHERNET SelfTestCommand = "ETHERNET"
	SELFTEST_WLAN     SelfTestCommand = "WLAN"
	SELFTEST_RS232    SelfTestCommand = "RS232"
	SELFTEST_SYSTEM   SelfTestCommand = "SYSTEM"
	SELFTEST_Z        SelfTestCommand = "Z"
	SELFTEST_BT       SelfTestCommand = "BT"
)

type SelfTestCommand string

func (st SelfTestCommand) GetMessage() ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString(SELFTEST_NAME)
	if len(string(st)) > 0 {
		buf.WriteString(EMPTY_SPACE)
		buf.WriteString(string(st))
	}
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}
