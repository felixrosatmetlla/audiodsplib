package distortion

import (
	"math"
)

func InfiniteClipping(signal []float64) []float64 {
	var bufferSize = len(signal)
	var output = make([]float64, bufferSize)

	for index, value := range signal {
		if value >= 0 {
			output[index] = 1
		} else {
			output[index] = -1
		}
	}

	return output
}

func HardClipping(signal []float64, threshold float64) []float64 {
	var bufferSize = len(signal)
	var output = make([]float64, bufferSize)

	for index, value := range signal {
		if value >= threshold {
			output[index] = threshold

		} else if value <= -threshold {
			output[index] = -threshold

		} else {
			output[index] = value
		}
	}

	return output
}

func CubicDistortion(signal []float64, amplitude float64) []float64 {
	var bufferSize = len(signal)
	var output = make([]float64, bufferSize)

	for index, value := range signal {
		output[index] = value - amplitude*(1/3)*math.Pow(value, 3)
	}

	return output
}

func ArctangentDistortion(signal []float64, alpha float64) []float64 {
	var bufferSize = len(signal)
	var output = make([]float64, bufferSize)

	for index, value := range signal {
		output[index] = (2 / math.Pi) * math.Atan(value*alpha)
	}

	return output
}
