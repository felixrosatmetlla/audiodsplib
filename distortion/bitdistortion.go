package distortion

import "math"

func BitReduction(signal []float64, numberBits int) []float64 {
	var bufferSize = len(signal)
	var output = make([]float64, bufferSize)

	amplitudeValues := math.Pow(2, float64(numberBits))

	for index, value := range signal {
		auxScaledInput := 0.5*value + 0.5
		scaledInput := amplitudeValues * auxScaledInput
		roundedInput := math.Round(scaledInput)

		auxOutput := roundedInput / amplitudeValues
		output[index] = 2*auxOutput - 1
	}

	return output
}