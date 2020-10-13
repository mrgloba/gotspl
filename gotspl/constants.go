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

const (
	MEASUREMENT_SYSTEM_METRIC MeasurementSystem = iota + 1
	MEASUREMENT_SYSTEM_DOT
	MEASUREMENT_SYSTEM_ENGLISH

	EMPTY_SPACE           = " "
	VALUE_SEPARATOR       = ","
	EQUAL_SYMBOL          = "="
	DOUBLE_QUOTE          = "\""
	STATUS_COMMAND_PREFIX = "!"

	UNIT_MM  = "mm"
	UNIT_DOT = "dot"
	LF       = 0x0A
	CR       = 0x0D
	ESC      = 0x1B
)

var (
	LINE_ENDING_BYTES = []byte{LF}
    ROTATION_ANGLES = []int{0, 90, 180, 270}
)
