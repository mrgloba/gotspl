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

func TestGapCmd(t *testing.T) {
	tests := []struct {
		name string
		want gotspl.GapBuilder
	}{
		{name: "Got GapBuilder", want: gotspl.GapCmd()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotspl.GapCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GapCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGapImpl_GetMessage(t *testing.T) {
	type fields struct {
		labelDistance       float64
		labelOffsetDistance float64
	}

	tests := []struct {
		name        string
		fields      fields
		want        []byte
		unitsSystem gotspl.MeasurementSystem
		wantErr     bool
	}{
		{"Build Gap Command",
			fields{2.2, 0},
			append([]byte("GAP 2.2,0"), gotspl.LINE_ENDING_BYTES...),
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			false,
		},

		{
			"Build Gap Command without LabelOffsetDistance should error",
			fields{2, -1},
			nil,
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			true,
		},

		{
			"Build Gap Command without LabelDistance should error",
			fields{-1, 2},
			nil,
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			true,
		},

		{
			"Build GAP Command measurement units mm",
			fields{2, 2.2},
			append([]byte("GAP 2 mm,2.2 mm"), gotspl.LINE_ENDING_BYTES...),
			gotspl.MEASUREMENT_SYSTEM_METRIC,
			false,
		},

		{
			"Build GAP Command measurement units dot",
			fields{2, 2},
			append([]byte("GAP 2 dot,2 dot"), gotspl.LINE_ENDING_BYTES...),
			gotspl.MEASUREMENT_SYSTEM_DOT,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotspl.TSPLInitialize(tt.unitsSystem)

			gi := gotspl.GapCmd()

			if tt.fields.labelDistance != -1 {
				gi = gi.LabelDistance(tt.fields.labelDistance)
			}

			if tt.fields.labelOffsetDistance != -1 {
				gi = gi.LabelOffsetDistance(tt.fields.labelOffsetDistance)
			}

			got, err := gi.GetMessage()
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
