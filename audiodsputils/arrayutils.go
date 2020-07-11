package audiodsputils

// GetArrayMinMax returns the minimum and maximum value of an array
func GetArrayMinMax(array []float64) (min float64, max float64) {
	max = array[0]
	min = array[0]

	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}

	return min, max
}

// CompareArrayValues compares the values of 2 mono signals
// Returns a bool to indicate if the signals are equal or not
func CompareArrayValues(a, b []float64) bool {

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
