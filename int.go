package defaults

func Int8(first int8, others ...int8) int8 {
	if first != 0 {
		return first
	}

	for _, o := range others {
		if o != 0 {
			return o
		}
	}

	return 0
}

func Int16(first int16, others ...int16) int16 {
	if first != 0 {
		return first
	}

	for _, o := range others {
		if o != 0 {
			return o
		}
	}

	return 0
}

func Int32(first int32, others ...int32) int32 {
	if first != 0 {
		return first
	}

	for _, o := range others {
		if o != 0 {
			return o
		}
	}

	return 0
}

func Int64(first int64, others ...int64) int64 {
	if first != 0 {
		return first
	}

	for _, o := range others {
		if o != 0 {
			return o
		}
	}

	return 0
}

func Int(first int, others ...int) int {
	if first != 0 {
		return first
	}

	for _, o := range others {
		if o != 0 {
			return o
		}
	}

	return 0
}
