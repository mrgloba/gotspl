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

func TestDataMatrixCmd(t *testing.T) {
	tests := []struct {
		name string
		want gotspl.DataMatrixBuilder
	}{
		{name: "Got SizeBuilder", want: gotspl.DataMatrixCmd()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotspl.DataMatrixCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SizeCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataMatrixImpl_GetMessage(t *testing.T) {
	type fields struct {
		xCoordinate             int
		yCoordinate             int
		width                   int
		height                  int
		escapeSequenceCharacter int
		moduleSize              int
		rotation                int
		isRectangle             int

		numberCols int
		numberRows int
		content    string
	}

	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{"Build Size Command",
			fields{
				10,
				10,
				100,
				100,
				-1,
				-1,
				-1,
				-1,
				-1,
				-1,
				"TEST CONTENT"},
			append([]byte("DATAMATRIX 10,10,100,100, \"TEST CONTENT\""), gotspl.LINE_ENDING_BYTES...),
			false,
		},

		{"Build Size Command with escapeSequenceCharacter",
			fields{
				10,
				10,
				100,
				100,
				126,
				-1,
				-1,
				-1,
				-1,
				-1,
				"TEST CONTENT"},
			append([]byte("DATAMATRIX 10,10,100,100,c126, \"TEST CONTENT\""), gotspl.LINE_ENDING_BYTES...),
			false,
		},

		{"Build Size Command with moduleSize",
			fields{
				10,
				10,
				100,
				100,
				-1,
				18,
				-1,
				-1,
				-1,
				-1,
				"TEST CONTENT"},
			append([]byte("DATAMATRIX 10,10,100,100,x18, \"TEST CONTENT\""), gotspl.LINE_ENDING_BYTES...),
			false,
		},

		{"Build Size Command with rotation",
			fields{
				10,
				10,
				100,
				100,
				-1,
				-1,
				90,
				-1,
				-1,
				-1,
				"TEST CONTENT"},
			append([]byte("DATAMATRIX 10,10,100,100,r90, \"TEST CONTENT\""), gotspl.LINE_ENDING_BYTES...),
			false,
		},

		{"Build Size Command with rectangle",
			fields{
				10,
				10,
				100,
				100,
				-1,
				-1,
				-1,
				1,
				-1,
				-1,
				"TEST CONTENT"},
			append([]byte("DATAMATRIX 10,10,100,100,a1, \"TEST CONTENT\""), gotspl.LINE_ENDING_BYTES...),
			false,
		},

		{"Build Size Command with rectangle negative",
			fields{
				10,
				10,
				100,
				100,
				-1,
				-1,
				-1,
				0,
				-1,
				-1,
				"TEST CONTENT"},
			append([]byte("DATAMATRIX 10,10,100,100,a0, \"TEST CONTENT\""), gotspl.LINE_ENDING_BYTES...),
			false,
		},

		{"Build Size Command with cols and rows",
			fields{
				10,
				10,
				100,
				100,
				-1,
				-1,
				-1,
				-1,
				11,
				10,
				"TEST CONTENT"},
			append([]byte("DATAMATRIX 10,10,100,100,10,11, \"TEST CONTENT\""), gotspl.LINE_ENDING_BYTES...),
			false,
		},

		{"Build Size Command with all fields",
			fields{
				10,
				10,
				100,
				100,
				126,
				18,
				180,
				1,
				11,
				10,
				"TEST CONTENT"},
			append([]byte("DATAMATRIX 10,10,100,100,c126,x18,r180,a1,10,11, \"TEST CONTENT\""), gotspl.LINE_ENDING_BYTES...),
			false,
		},

		{"Build Size Command without xCoordinate should error",
			fields{
				-1,
				10,
				100,
				100,
				126,
				18,
				180,
				1,
				11,
				10,
				"TEST CONTENT"},
			nil,
			true,
		},

		{"Build Size Command without yCoordinate should error",
			fields{
				10,
				-1,
				100,
				100,
				126,
				18,
				180,
				1,
				11,
				10,
				"TEST CONTENT"},
			nil,
			true,
		},

		{"Build Size Command without width should error",
			fields{
				10,
				10,
				-1,
				100,
				126,
				18,
				180,
				1,
				11,
				10,
				"TEST CONTENT"},
			nil,
			true,
		},

		{"Build Size Command without height should error",
			fields{
				10,
				10,
				100,
				-1,
				126,
				18,
				180,
				1,
				11,
				10,
				"TEST CONTENT"},
			nil,
			true,
		},

		{"Build Size Command with wrong value escapeSequenceCharacter should error",
			fields{
				10,
				10,
				100,
				100,
				256,
				18,
				180,
				1,
				11,
				10,
				"TEST CONTENT"},
			nil,
			true,
		},

		{"Build Size Command with wrong value rotation should error",
			fields{
				10,
				10,
				100,
				100,
				126,
				18,
				91,
				1,
				11,
				10,
				"TEST CONTENT"},
			nil,
			true,
		},

		{"Build Size Command with wrong value numberCols should error",
			fields{
				10,
				10,
				100,
				100,
				126,
				18,
				90,
				1,
				gotspl.DATAMATRIX_COLS_MAX + 1,
				10,
				"TEST CONTENT"},
			nil,
			true,
		},

		{"Build Size Command with wrong value numberRows should error",
			fields{
				10,
				10,
				100,
				100,
				126,
				18,
				90,
				1,
				10,
				gotspl.DATAMATRIX_ROWS_MAX + 1,
				"TEST CONTENT"},
			nil,
			true,
		},

		{"Build Size Command without content should error",
			fields{
				10,
				10,
				100,
				100,
				126,
				18,
				90,
				1,
				10,
				10,
				"-1"},
			nil,
			true,
		},

		{"Build Size Command with wrong value content should error",
			fields{
				10,
				10,
				100,
				100,
				126,
				18,
				90,
				1,
				10,
				10,
				""},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			si := gotspl.DataMatrixCmd()

			if tt.fields.xCoordinate != -1 {
				si = si.XCoordinate(tt.fields.xCoordinate)
			}

			if tt.fields.yCoordinate != -1 {
				si = si.YCoordinate(tt.fields.yCoordinate)
			}

			if tt.fields.width != -1 {
				si = si.Width(tt.fields.width)
			}

			if tt.fields.height != -1 {
				si = si.Height(tt.fields.height)
			}

			if tt.fields.escapeSequenceCharacter != -1 {
				si = si.EscapeSequenceCharacter(tt.fields.escapeSequenceCharacter)
			}

			if tt.fields.moduleSize != -1 {
				si = si.ModuleSize(tt.fields.moduleSize)
			}

			if tt.fields.rotation != -1 {
				si = si.Rotation(tt.fields.rotation)
			}

			if tt.fields.isRectangle != -1 {
				si = si.IsRectangle(tt.fields.isRectangle == 1)
			}

			if tt.fields.numberRows != -1 {
				si = si.NumberRows(tt.fields.numberRows)
			}

			if tt.fields.numberCols != -1 {
				si = si.NumberCols(tt.fields.numberCols)
			}

			if tt.fields.content != "-1" {
				si = si.Content(tt.fields.content)
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
