package defaults

func Uint8(value, defaultValue uint8) uint8 {
	if value == 0 {
		return defaultValue
	}
	return value
}

func Uint16(value, defaultValue uint16) uint16 {
	if value == 0 {
		return defaultValue
	}
	return value
}

func Uint32(value, defaultValue uint32) uint32 {
	if value == 0 {
		return defaultValue
	}
	return value
}

func Uint64(value, defaultValue uint64) uint64 {
	if value == 0 {
		return defaultValue
	}
	return value
}

func Uint(value, defaultValue uint) uint {
	if value == 0 {
		return defaultValue
	}
	return value
}
