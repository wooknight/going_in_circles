// Code generated by "stringer --type StatusCode --trimprefix status"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[statusOK-200]
	_ = x[statusInternalServerErr-500]
}

const (
	_StatusCode_name_0 = "OK"
	_StatusCode_name_1 = "InternalServerErr"
)

func (i StatusCode) String() string {
	switch {
	case i == 200:
		return _StatusCode_name_0
	case i == 500:
		return _StatusCode_name_1
	default:
		return "StatusCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
