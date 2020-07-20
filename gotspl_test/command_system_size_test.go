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
package gotspl_test

import (
	"bytes"
	"github.com/mrgloba/gotspl/gotspl"
	"reflect"
	"testing"
)

func TestSizeCmd(t *testing.T) {
	tests := []struct {
		name string
		want gotspl.SizeBuilder
	}{
		{name: "Got SizeBuilder", want: gotspl.SizeCmd()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotspl.SizeCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SizeCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSizeImpl_GetMessage(t *testing.T) {
	type fields struct {
		labelWidth  float64
		labelLength float64
	}

	tests := []struct {
		name        string
		fields      fields
		want        []byte
		unitsSystem gotspl.MeasurementSystem
		wantErr     bool
	}{
		{"Build Size Command",
			fields{2.2, 3},
			append([]byte("SIZE 2.2,3"), gotspl.LINE_ENDING_BYTES...),
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			false,
		},

		{
			"Build Size Command without LabelLength",
			fields{2, -1},
			append([]byte("SIZE 2"), gotspl.LINE_ENDING_BYTES...),
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			false,
		},

		{
			"Build Size Command without LabelWidth should error",
			fields{-1, 2},
			nil,
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			true,
		},

		{
			"Build Size Command measurement units mm",
			fields{2, 2.2},
			append([]byte("SIZE 2 mm,2.2 mm"), gotspl.LINE_ENDING_BYTES...),
			gotspl.MEASUREMENT_SYSTEM_METRIC,
			false,
		},

		{
			"Build Size Command measurement units dot",
			fields{2, 2},
			append([]byte("SIZE 2 dot,2 dot"), gotspl.LINE_ENDING_BYTES...),
			gotspl.MEASUREMENT_SYSTEM_DOT,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotspl.TSPLInitialize(tt.unitsSystem)

			si := gotspl.SizeCmd()

			if tt.fields.labelWidth != -1 {
				si = si.LabelWidth(tt.fields.labelWidth)
			}

			if tt.fields.labelLength != -1 {
				si = si.LabelLength(tt.fields.labelLength)
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
