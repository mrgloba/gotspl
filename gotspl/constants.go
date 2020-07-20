package gotspl

const (
	MEASUREMENT_SYSTEM_METRIC MeasurementSystem = iota + 1
	MEASUREMENT_SYSTEM_DOT
	MEASUREMENT_SYSTEM_ENGLISH

	EMPTY_SPACE     = " "
	VALUE_SEPARATOR = ","
	DOUBLE_QUOTE    = "\""

	UNIT_MM  = "mm"
	UNIT_DOT = "dot"
	LF       = 0x13
)

var (
	LINE_ENDING_BYTES = []byte{LF}
)
