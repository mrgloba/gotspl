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
