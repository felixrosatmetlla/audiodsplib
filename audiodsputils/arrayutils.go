package audiodsputils

func transformRange(valueToTransform float64, actualMin float64, actualMax float64, newMin float64, newMax float64) float64 {
	var newValue float64 = ((valueToTransform-actualMin)/(actualMax-actualMin))*(newMax-newMin) + newMin

	return newValue
}

// GetArrayMinMax returns the minimum and maximum value of an array
func GetArrayMinMax(array []float64) (float64, float64) {
	var max float64 = array[0]
	var min float64 = array[0]

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

// CompareMonoSignals compares the values of 2 mono signals
// Returns a bool to indicate if the signals are equal or not
func CompareMonoSignals(a, b []float64) bool {

	// If one is nil, the other must also be nil.
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
