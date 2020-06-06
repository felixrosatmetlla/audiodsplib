package distortion

import (
	"errors"
	"math"
)

// InfiniteClipping distorts a signal clipping it to values 1 and -1
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

// HardClipping distorts a signal clipping it the value indicated by the threshold parameter
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

// CubicDistortion distorts a signal using a cubic function
//
// Input variables:
//  signal: Input signal to distort
//  amplitude: The drive amount of the distortion. Range from [0, 1]:
//   0: no distortion
//   1: maximum amount of distortion
func CubicDistortion(signal []float64, amplitude float64) []float64 {
	var bufferSize = len(signal)
	var output = make([]float64, bufferSize)

	for index, value := range signal {
		output[index] = value - amplitude*(1./3.)*math.Pow(value, 3)
	}

	return output
}

// ArctangentDistortion distorts a signal using an arctangent function
//
// Input variables:
//  signal: Input signal to distort
//  alpha: The drive amount of the distortion
//   Range from [1, 10]: higher -> more distortion
func ArctangentDistortion(signal []float64, alpha float64) ([]float64, error) {
	if alpha < 1 || alpha > 10 {
		return []float64{}, errors.New("distortion: invalid alpha range")
	}

	var bufferSize = len(signal)
	var output = make([]float64, bufferSize)

	for index, value := range signal {
		output[index] = (2 / math.Pi) * math.Atan(value*alpha)
	}

	return output, nil
}

func SineDistortion(signal []float64) []float64 {
	var bufferSize = len(signal)
	var output = make([]float64, bufferSize)

	for index, value := range signal {
		output[index] = math.Sin(2 * math.Pi * value)
	}

	return output
}
