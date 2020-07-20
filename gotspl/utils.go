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
	"strconv"
)

func hasFloatDecimals(f float64) bool {
	return (f - float64(int(f))) != 0
}

func getUnits() string {
	switch measurementSystem {
	case MEASUREMENT_SYSTEM_METRIC:
		return EMPTY_SPACE + UNIT_MM
	case MEASUREMENT_SYSTEM_DOT:
		return EMPTY_SPACE + UNIT_DOT
	default:
		return ""
	}
}

func formatFloatWithUnits(value float64, useUnits bool) []byte {
	buf := bytes.Buffer{}
	if !hasFloatDecimals(value) {
		buf.WriteString(strconv.Itoa(int(value)))
	} else {
		buf.WriteString(strconv.FormatFloat(value, 'f', -1, 64))
	}

	if useUnits {
		buf.WriteString(getUnits())
	}

	return buf.Bytes()
}

func findIntInSlice(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}
