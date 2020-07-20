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

func TestSoundCmd(t *testing.T) {
	tests := []struct {
		name string
		want gotspl.SoundBuilder
	}{
		{name: "Got SoundBuilder", want: gotspl.SoundCmd()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotspl.SoundCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SoundCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSoundImpl_GetMessage(t *testing.T) {
	type fields struct {
		volumeLevel    int
		timingInterval int
	}

	tests := []struct {
		name        string
		fields      fields
		want        []byte
		unitsSystem gotspl.MeasurementSystem
		wantErr     bool
	}{
		{"Build Sound Command",
			fields{2, 500},
			append([]byte("SOUND 2,500"), gotspl.LINE_ENDING_BYTES...),
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			false,
		},

		{
			"Build Sound Command without volumeLevel should error",
			fields{-1, 500},
			nil,
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			true,
		},

		{
			"Build Sound Command without timingLevel should error",
			fields{2, -1},
			nil,
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			true,
		},

		{
			"Build Sound Command measurement units mm",
			fields{2, 500},
			append([]byte("SOUND 2,500"), gotspl.LINE_ENDING_BYTES...),
			gotspl.MEASUREMENT_SYSTEM_METRIC,
			false,
		},

		{
			"Build Sound Command volumeLevel out if range",
			fields{gotspl.SOUND_LEVEL_MAX + 1, 500},
			nil,
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			true,
		},

		{
			"Build Sound Command timingInterval out if range",
			fields{2, gotspl.SOUND_INTERVAL_MAX + 1},
			nil,
			gotspl.MEASUREMENT_SYSTEM_ENGLISH,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotspl.TSPLInitialize(tt.unitsSystem)

			si := gotspl.SoundCmd()

			if tt.fields.volumeLevel != -1 {
				si = si.VolumeLevel(tt.fields.volumeLevel)
			}

			if tt.fields.timingInterval != -1 {
				si = si.TimingInterval(tt.fields.timingInterval)
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
