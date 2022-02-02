package vbcode

func prepend(buf []byte, b byte) []byte {
	ret := make([]byte, 0, 1)
	ret = append(ret, b)
	ret = append(ret, buf...)
	return ret
}

func Encode(n int) []byte {
	buf := make([]byte, 1)
	for {
		var mod = byte(n % 128)
		buf[0] += mod
		if n < 128 {
			break
		}
		n = n / 128
		buf = prepend(buf, 128)
	}

	return buf
}
