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
	"strings"
)

type VariableImpl struct {
	name  *string
	value *string
	quote bool
}

type VariableBuilder interface {
	TSPLCommand
	Name(name string) VariableBuilder
	Value(value string, quote bool) VariableBuilder
}

func VariableCmd() VariableBuilder {
	return VariableImpl{quote: false}
}

func Variable(name string, value string, quote bool) VariableBuilder {
	return VariableCmd().Name(name).Value(value, quote)
}

func StringVariable(name string, value string) VariableBuilder {
	return VariableCmd().Name(name).Value(value, true)
}

func IntVariable(name string, value int) VariableBuilder {
	return VariableCmd().Name(name).Value(strconv.Itoa(value), false)
}

func (v VariableImpl) GetMessage() ([]byte, error) {
	if v.name == nil || len(*v.name) == 0 || v.value == nil || len(*v.value) == 0 {
		return nil, errors.New("ParseError VALUE Command: name and value should be specified")
	}

	buf := bytes.NewBufferString(strings.ToUpper(*v.name))
	buf.WriteString(EQUAL_SYMBOL)
	if v.quote {
		buf.WriteString(DOUBLE_QUOTE)
	}
	buf.WriteString(*v.value)
	if v.quote {
		buf.WriteString(DOUBLE_QUOTE)
	}

	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (v VariableImpl) Name(name string) VariableBuilder {
	if v.name == nil {
		v.name = new(string)
	}
	*v.name = name
	return v
}

func (v VariableImpl) Value(value string, quote bool) VariableBuilder {
	if v.value == nil {
		v.value = new(string)
	}
	*v.value = value
	v.quote = quote
	return v
}
