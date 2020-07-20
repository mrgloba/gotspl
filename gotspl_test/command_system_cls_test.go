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

func TestClsCmd(t *testing.T) {
	tests := []struct {
		name string
		want gotspl.ClsBuilder
	}{
		{name: "Got EndBuilder", want: gotspl.ClsCmd()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotspl.ClsCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClsCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClsImpl_GetMessage(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{"Build Cls Command",
			append([]byte("CLS"), gotspl.LINE_ENDING_BYTES...),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			si := gotspl.ClsCmd()

			got, _ := si.GetMessage()

			if bytes.Compare(got, tt.want) != 0 {
				t.Errorf("GetMessage() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
