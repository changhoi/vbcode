package vbcode

func Decode(vb []byte) int {
	var ret int

	for _, v := range vb {
		ret += int(v & 127)
		if v&128 == 128 {
			ret <<= 7
		}
	}

	return ret
}
