package qry

// appendJoin appends vals to dst, separated by sep, without building an
// intermediate joined string.
func appendJoin(dst []byte, sep string, vals []string) []byte {
	for i, v := range vals {
		if i > 0 {
			dst = append(dst, sep...)
		}
		dst = append(dst, v...)
	}
	return dst
}

// joinSize reports how many bytes appendJoin writes for vals and sep.
func joinSize(sep string, vals []string) int {
	if len(vals) == 0 {
		return 0
	}

	n := len(sep) * (len(vals) - 1)
	for _, v := range vals {
		n += len(v)
	}
	return n
}

// intSize reports how many bytes the base-10 representation of v needs.
func intSize(v int) int {
	if v < 0 {
		return 1 + uintSize(uint64(-(v+1))+1)
	}
	return uintSize(uint64(v))
}

// uintSize reports how many bytes the base-10 representation of v needs.
func uintSize(v uint64) int {
	n := 1
	for v >= 10 {
		v /= 10
		n++
	}
	return n
}
