package amplitude

import (
	"math"

	"github.com/felixrosatmetlla/audiodsplib/types"
)

//TODO: MultiChannel options

// ChangeGain modifies the signals gain using a linear value of Amplitude
func ChangeGain(signal types.Signal, gain float64) types.Signal {
	var bufferSize = signal.NumSamples * signal.Channels
	var outputBuffer = make([]float64, bufferSize)

	for index := range outputBuffer {
		outputBuffer[index] = signal.Data[index] * gain
	}

	outputSignal := types.Signal{
		Data:       outputBuffer,
		Channels:   signal.Channels,
		Samplerate: signal.Samplerate,
		NumSamples: signal.NumSamples,
	}

	return outputSignal
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

// InvertPolarity computes a polarly inverted signal of a mono signal
func InvertPolarity(signal []float64) []float64 {
	var bufferSize = len(signal)
	var output = make([]float64, bufferSize)

	for index := range output {
		output[index] = signal[index] * -1
	}

	return output
}
