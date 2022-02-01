package helper

func StringSliceContains(haystack []string, needle string) bool {

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

func StringToInterfaceSlice(input []string) []interface{} {

	slice := make([]interface{}, len(input))
	for index, num := range input {
		slice[index] = num
	}

	return slice
}

func IntegerToInterfaceSlice(input []int) []interface{} {

	slice := make([]interface{}, len(input))
	for index, num := range input {
		slice[index] = num
	}

	return slice
}
