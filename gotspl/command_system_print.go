package gotspl

import (
	"bytes"
	"errors"
	"strconv"
)

const (
	PRINT_NAME             = "PRINT"
	PRINT_NUMBERLABELS_MIN = 1
	PRINT_NUMBERLABELS_MAX = 999999999
	PRINT_NUMBERCOPIES_MIN = 0
	PRINT_NUMBERCOPIES_MAX = 999999999
)

type PrintImpl struct {
	numberLabels *int
	numberCopies *int
}

type PrintBuilder interface {
	TSPLCommand
	NumberLabels(numberLabels int) PrintBuilder
	NumberCopies(numberCopies int) PrintBuilder
}

func PrintCmd() PrintBuilder {
	return PrintImpl{}
}

func (p PrintImpl) GetMessage() ([]byte, error) {
	if p.numberLabels == nil {
		return nil, errors.New("ParseError PRINT Command: numberLabels should be specified")
	}

	if !(*p.numberLabels < PRINT_NUMBERLABELS_MAX && *p.numberLabels > PRINT_NUMBERLABELS_MIN) {
		return nil, errors.New("ParseError PRINT Command: numberLabels parameter must be between " +
			strconv.Itoa(PRINT_NUMBERLABELS_MIN) + " and " + strconv.Itoa(PRINT_NUMBERLABELS_MAX))
	}

	if p.numberCopies != nil {
		if !(*p.numberCopies < PRINT_NUMBERCOPIES_MAX && *p.numberCopies > PRINT_NUMBERCOPIES_MIN) {
			return nil, errors.New("ParseError PRINT Command: numberCopies parameter must be between " +
				strconv.Itoa(PRINT_NUMBERCOPIES_MIN) + " and " + strconv.Itoa(PRINT_NUMBERCOPIES_MAX))
		}
	}

	buf := bytes.NewBufferString(PRINT_NAME)
	buf.WriteString(EMPTY_SPACE)

	buf.WriteString(strconv.Itoa(*p.numberLabels))

	if p.numberCopies != nil {
		buf.WriteString(VALUE_SEPARATOR)
		buf.WriteString(strconv.Itoa(*p.numberCopies))
	}

	buf.Write(LINE_ENDING_BYTES)

	return buf.Bytes(), nil
}

func (p PrintImpl) NumberLabels(numberLabels int) PrintBuilder {
	if p.numberLabels == nil {
		p.numberLabels = new(int)
	}

	*p.numberLabels = numberLabels
	return p
}

func (p PrintImpl) NumberCopies(numberCopies int) PrintBuilder {
	if p.numberCopies == nil {
		p.numberCopies = new(int)
	}
	*p.numberCopies = numberCopies
	return p
}
