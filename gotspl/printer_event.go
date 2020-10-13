package gotspl

import "reflect"

type PrinterEvent interface {
	RawValue() []byte
	EventType() reflect.Type
	Value() interface{}
}

type RawResponseEvent struct {
	Size int
	Data []byte
}