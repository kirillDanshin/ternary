package ternary

// Int returns int on any condition value
func Int(cond bool, onTrue int, onFalse int) int {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntStr returns int on true or string on false condition
func IntStr(cond bool, onTrue int, onFalse string) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntInt8 returns int on true or int8 on false condition
func IntInt8(cond bool, onTrue int, onFalse int8) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntInt16 returns int on true or int16 on false condition
func IntInt16(cond bool, onTrue int, onFalse int16) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntInt32 returns int on true or int32 on false condition
func IntInt32(cond bool, onTrue int, onFalse int32) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntInt64 returns int on true or int64 on false condition
func IntInt64(cond bool, onTrue int, onFalse int64) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntFloat32 returns int on true or float32 on false condition
func IntFloat32(cond bool, onTrue int, onFalse float32) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntFloat64 returns int on true or float64 on false condition
func IntFloat64(cond bool, onTrue int, onFalse float64) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntCmplx64 returns int on true or complex64 on false condition
func IntCmplx64(cond bool, onTrue int, onFalse complex64) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntCmplx128 returns int on true or complex128 on false condition
func IntCmplx128(cond bool, onTrue int, onFalse complex128) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntUint returns int on true or uint on false condition
func IntUint(cond bool, onTrue int, onFalse uint) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntUint8 returns int on true or uint8 on false condition
func IntUint8(cond bool, onTrue int, onFalse uint8) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntUint16 returns int on true or uint16 on false condition
func IntUint16(cond bool, onTrue int, onFalse uint16) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntUint32 returns int on true or uint32 on false condition
func IntUint32(cond bool, onTrue int, onFalse uint32) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntUint64 returns int on true or uint64 on false condition
func IntUint64(cond bool, onTrue int, onFalse uint64) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntUintptr returns int on true or uintptr on false condition
func IntUintptr(cond bool, onTrue int, onFalse uintptr) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntBool returns int on true or bool on false condition
func IntBool(cond bool, onTrue int, onFalse bool) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntIface returns int on true or interface on false condition
func IntIface(cond bool, onTrue int, onFalse interface{}) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}

// IntError returns int on true or interface on false condition
func IntError(cond bool, onTrue int, onFalse error) interface{} {
	if cond {
		return onTrue
	}
	return onFalse
}
