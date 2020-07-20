package gotspl_test

import (
	"bytes"
	"github.com/mrgloba/gotspl/gotspl"
	"reflect"
	"testing"
)

func TestEndCmd(t *testing.T) {
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

func TestEndImpl_GetMessage(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{"Build End Command",
			append([]byte("END"), gotspl.LINE_ENDING_BYTES...),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			si := gotspl.EndCmd()

			got, _ := si.GetMessage()

			if bytes.Compare(got, tt.want) != 0 {
				t.Errorf("GetMessage() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
