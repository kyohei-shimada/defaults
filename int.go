package defaults

func Int8(value, defaultValue int8) int8 {
	if value == 0 {
		return defaultValue
	}
	return value
}

func Int16(value, defaultValue int16) int16 {
	if value == 0 {
		return defaultValue
	}
	return value
}

func Int32(value, defaultValue int32) int32 {
	if value == 0 {
		return defaultValue
	}
	return value
}

func Int64(value, defaultValue int64) int64 {
	if value == 0 {
		return defaultValue
	}
	return value
}

func Int(value, defaultValue int) int {
	if value == 0 {
		return defaultValue
	}
	return value
}
