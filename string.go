package defaults

func String(first string, others ...string) string {
	if first != "" {
		return first
	}

	for _, o := range others {
		if o != "" {
			return o
		}
	}

	return ""
}
