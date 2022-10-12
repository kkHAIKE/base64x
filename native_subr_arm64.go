// +build !noasm !appengine
// Code generated by nocgo, DO NOT EDIT.

package base64x

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __native_entry__() uintptr

func alignEntry() uintptr {
	r := __native_entry__()
	if r&4095 == 0 {
		return r
	}

	r += 6144
	if r&4095 != 0 {
		panic("oops")
	}

	return r
}

var (
	_subr__b64decode = alignEntry() + 404
	_subr__b64encode = alignEntry() + 12
)

const (
	_stack__b64decode = 48
	_stack__b64encode = 0
)

var (
	_ = _subr__b64decode
	_ = _subr__b64encode
)

const (
	_ = _stack__b64decode
	_ = _stack__b64encode
)