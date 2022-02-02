package vbcode

import "testing"

func TestDecode(t *testing.T) {
	testCase := []struct {
		in  []byte
		out int
	}{
		{

			[]byte{0b1},
			1,
		},
		{

			[]byte{0b00000111},
			7,
		},
		{

			[]byte{0b01111111},
			127,
		},
		{
			[]byte{0b10000001, 0b00000000},
			128,
		},
	}
	for i, tc := range testCase {
		ret := Decode(tc.in)
		if ret != tc.out {
			t.Errorf("%d: input: %v, want: %v, result:%v", i, tc.in, tc.out, ret)
		}
	}

}
