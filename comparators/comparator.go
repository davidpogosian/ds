package comparators

type Comparator[T any] func(a, b T) int

// ComparatorString is a comparator function for the string type.
// It compares two strings lexicographically.
func ComparatorString(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// ComparatorBool is a comparator function for the bool type.
// It considers false as less than true.
func ComparatorBool(a, b bool) int {
	if !a && b {
		return -1
	} else if a && !b {
		return 1
	}
	return 0
}

// ComparatorInt is a comparator function for the int type.
func ComparatorInt(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// ComparatorInt8 is a comparator function for the int8 type.
func ComparatorInt8(a, b int8) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// ComparatorInt16 is a comparator function for the int16 type.
func ComparatorInt16(a, b int16) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// ComparatorInt32 is a comparator function for the int32 type.
func ComparatorInt32(a, b int32) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// ComparatorInt64 is a comparator function for the int64 type.
func ComparatorInt64(a, b int64) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// ComparatorUint is a comparator function for the uint type.
func ComparatorUint(a, b uint) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// ComparatorUint8 is a comparator function for the uint8 type.
func ComparatorUint8(a, b uint8) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// ComparatorUint16 is a comparator function for the uint16 type.
func ComparatorUint16(a, b uint16) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// ComparatorUint32 is a comparator function for the uint32 type.
func ComparatorUint32(a, b uint32) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// ComparatorUint64 is a comparator function for the uint64 type.
func ComparatorUint64(a, b uint64) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// ComparatorByte is a comparator function for the byte type.
// It compares two bytes numerically.
func ComparatorByte(a, b byte) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// ComparatorFloat32 is a comparator function for the float32 type.
func ComparatorFloat32(a, b float32) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// ComparatorFloat64 is a comparator function for the float64 type.
func ComparatorFloat64(a, b float64) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
