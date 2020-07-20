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
