package gotspl_test

import (
	"bytes"
	"github.com/mrgloba/gotspl/gotspl"
	"reflect"
	"testing"
)

func TestSpeedCmd(t *testing.T) {
	tests := []struct {
		name string
		want gotspl.SpeedBuilder
	}{
		{name: "Got SpeedBuilder", want: gotspl.SpeedCmd()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotspl.SpeedCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpeedCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpeedImpl_GetMessage(t *testing.T) {
	type fields struct {
		printSpeed float64
	}

	tests := []struct {
		name        string
		fields      fields
		want        []byte
		unitsSystem gotspl.MeasurementSystem
		wantErr     bool
	}{
		{"Build Speed Command",
			fields{2.5},
			append([]byte("SPEED 2.5"), gotspl.LINE_ENDING_BYTES...),
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			false,
		},

		{
			"Build Speed Command without printSpeed should error",
			fields{-1},
			nil,
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			true,
		},

		{
			"Build Speed Command measurement units mm",
			fields{2},
			append([]byte("SPEED 2"), gotspl.LINE_ENDING_BYTES...),
			gotspl.MEASUREMENT_SYSTEM_METRIC,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotspl.TSPLInitialize(tt.unitsSystem)

			si := gotspl.SpeedCmd()

			if tt.fields.printSpeed != -1 {
				si = si.PrintSpeed(tt.fields.printSpeed)
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
