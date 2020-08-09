package audiodsputils

// GetArrayMinMax returns the minimum and maximum value of an array
//
// Input:
//  array: Input array from which min and max are wanted
// Output:
//  min: minimum value of the signal
//  max: maximum value of the signal
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
//
// Input:
//  arrayA: Input array to compare
//  arrayB: Second input array to compare
// Output:
//  result: true if the signals are equal
func CompareArrayValues(arrayA, arrayB []float64) bool {

	if (arrayA == nil) != (arrayB == nil) {
		return false
	}

	if len(arrayA) != len(arrayB) {
		return false
	}

	for i := range arrayA {
		if arrayA[i] != arrayB[i] {
			return false
		}
	}

	return true
}
