package amplitude

import (
	"math"
)

//TODO: Find a good way to do overload in Golang or not duplicate so much code

// ChangeGain modifies the signals gain using a linear value of Amplitude
func ChangeGain(signal []float64, gain float64) []float64 {
	var bufferSize = len(signal)
	var output = make([]float64, bufferSize)

	for index := range output {
		output[index] = signal[index] * gain
	}

	return output
}

// ChangeGaindB modifies the signals gain using a value in decibels as input
func ChangeGaindB(signal []float64, dBChange float64) []float64 {
	var bufferSize = len(signal)
	var output = make([]float64, bufferSize)
	var scale = math.Pow(10, dBChange/20)

	for index := range output {
		output[index] = signal[index] * scale
	}

	return output
}

//TODO: MultiChannel options

// InvertPolarity computes a polarly inverted signal of a mono signal
func InvertPolarity(signal []float64) []float64 {
	var bufferSize = len(signal)
	var output = make([]float64, bufferSize)

	for index := range output {
		output[index] = signal[index] * -1
	}

	return output
}
