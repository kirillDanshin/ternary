package ternary

// Str returns string on true or string on false condition
func Str(cond bool, onTrue string, onFalse string) string {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrInt returns string on true or int on false condition
func StrInt(cond bool, onTrue string, onFalse int) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrInt8 returns string on true or int8 on false condition
func StrInt8(cond bool, onTrue string, onFalse int8) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrInt16 returns string on true or int16 on false condition
func StrInt16(cond bool, onTrue string, onFalse int16) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrInt32 returns string on true or int32 on false condition
func StrInt32(cond bool, onTrue string, onFalse int32) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrInt64 returns string on true or int64 on false condition
func StrInt64(cond bool, onTrue string, onFalse int64) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrFloat32 returns string on true or float32 on false condition
func StrFloat32(cond bool, onTrue string, onFalse float32) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrFloat64 returns string on true or float64 on false condition
func StrFloat64(cond bool, onTrue string, onFalse float64) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrCmplx64 returns string on true or complex64 on false condition
func StrCmplx64(cond bool, onTrue string, onFalse complex64) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrCmplx128 returns string on true or complex128 on false condition
func StrCmplx128(cond bool, onTrue string, onFalse complex128) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrUint returns string on true or uint on false condition
func StrUint(cond bool, onTrue string, onFalse uint) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrUint8 returns string on true or uint8 on false condition
func StrUint8(cond bool, onTrue string, onFalse uint8) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrUint16 returns string on true or uint16 on false condition
func StrUint16(cond bool, onTrue string, onFalse uint16) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrUint32 returns string on true or uint32 on false condition
func StrUint32(cond bool, onTrue string, onFalse uint32) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrUint64 returns string on true or uint64 on false condition
func StrUint64(cond bool, onTrue string, onFalse uint64) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrUintptr returns string on true or uintptr on false condition
func StrUintptr(cond bool, onTrue string, onFalse uintptr) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrBool returns string on true or bool on false condition
func StrBool(cond bool, onTrue string, onFalse bool) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrIface returns string on true or interface on false condition
func StrIface(cond bool, onTrue string, onFalse interface{}) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// StrError returns string on true or error on false condition
func StrError(cond bool, onTrue string, onFalse error) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}
