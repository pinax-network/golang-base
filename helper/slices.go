package helper

// CompareStringSlices takes two string slices and compares its content. Returns true if both slices contain the same
// elements and the same number of elements. It does not check the order.
func CompareStringSlices(sliceA, sliceB []string) bool {

	if len(sliceA) != len(sliceB) {
		return false
	}

	for _, s := range sliceA {
		if !StringSliceContains(sliceB, s) {
			return false
		}
	}

	return true
}

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

func ToAnySlice[T any](input []T) []any {

	slice := make([]any, len(input))
	for index, num := range input {
		slice[index] = num
	}

	return slice
}

func SliceContains[T comparable](haystack []T, needle T) bool {

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

func EliminateDuplicates[T comparable](input []T) (result []T) {

	for _, e := range input {
		if !SliceContains(result, e) {
			result = append(result, e)
		}
	}

	return
}
