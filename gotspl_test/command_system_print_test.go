package gotspl_test

import (
	"bytes"
	"github.com/mrgloba/gotspl/gotspl"
	"reflect"
	"testing"
)

func TestPrintCmd(t *testing.T) {
	tests := []struct {
		name string
		want gotspl.PrintBuilder
	}{
		{name: "Got SoundBuilder", want: gotspl.PrintCmd()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotspl.PrintCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrintCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintImpl_GetMessage(t *testing.T) {
	type fields struct {
		numberLabels int
		numberCopies int
	}

	tests := []struct {
		name        string
		fields      fields
		want        []byte
		unitsSystem gotspl.MeasurementSystem
		wantErr     bool
	}{
		{"Build Print Command",
			fields{2, 1},
			append([]byte("PRINT 2,1"), gotspl.LINE_ENDING_BYTES...),
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			false,
		},

		{
			"Build Print Command without numberLabels should error",
			fields{-1, 1},
			nil,
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			true,
		},

		{
			"Build Print Command without numberCopies",
			fields{2, -1},
			append([]byte("PRINT 2"), gotspl.LINE_ENDING_BYTES...),
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			false,
		},

		{
			"Build Print Command should be without measurement units mm",
			fields{2, 2},
			append([]byte("PRINT 2,2"), gotspl.LINE_ENDING_BYTES...),
			gotspl.MEASUREMENT_SYSTEM_METRIC,
			false,
		},

		{
			"Build Print Command numberLabels out of range",
			fields{gotspl.PRINT_NUMBERLABELS_MAX + 1, 1},
			nil,
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			true,
		},

		{
			"Build Print Command numberCopies out out range",
			fields{2, gotspl.PRINT_NUMBERCOPIES_MAX + 1},
			nil,
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotspl.TSPLInitialize(tt.unitsSystem)

			si := gotspl.PrintCmd()

			if tt.fields.numberLabels != -1 {
				si = si.NumberLabels(tt.fields.numberLabels)
			}

			if tt.fields.numberCopies != -1 {
				si = si.NumberCopies(tt.fields.numberCopies)
			}

			got, err := si.GetMessage()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if bytes.Compare(got, tt.want) != 0 {
				t.Errorf("GetMessage() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
