package defaults

func Uint8(first uint8, others ...uint8) uint8 {
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

func Uint16(first uint16, others ...uint16) uint16 {
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

func Uint32(first uint32, others ...uint32) uint32 {
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

func Uint64(first uint64, others ...uint64) uint64 {
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

func Uint(first uint, others ...uint) uint {
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
