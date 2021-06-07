package helper

func StringArrayContains(haystack []string, needle string) bool {

	if haystack == nil || len(haystack) == 0 {
		return false
	}

	for _, c := range haystack {
		if c == needle {
			return true
		}
	}

	return false
}
