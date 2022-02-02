package vbcode

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	testCase := []struct {
		in  int
		out []byte
	}{
		{
			1,
			[]byte{0b1},
		},
		{
			7,
			[]byte{0b00000111},
		},
		{
			127,
			[]byte{0b01111111},
		},
		{
			128,
			[]byte{0b10000001, 0b00000000},
		},
	}

	for i, tc := range testCase {
		ret := Encode(tc.in)
		if !bytes.Equal(ret, tc.out) {
			t.Errorf("#%d: input: %v, want: %v, result: %v", i, tc.in, tc.out, ret)
		}
	}
}
