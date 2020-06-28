package amplitude

import (
	"math"

	"github.com/felixrosatmetlla/audiodsplib/types"
)

//TODO: MultiChannel options

// ChangeGain modifies the signals gain using a linear value of Amplitude
func ChangeGain(inputSignal types.Signal, gain float64) types.Signal {
	var bufferSize = inputSignal.NumSamples * inputSignal.Channels
	var outputBuffer = make([]float64, bufferSize)

	for index := range outputBuffer {
		outputBuffer[index] = inputSignal.Data[index] * gain
	}

	outputSignal := types.Signal{
		Data:       outputBuffer,
		Channels:   inputSignal.Channels,
		Samplerate: inputSignal.Samplerate,
		NumSamples: inputSignal.NumSamples,
	}

	return outputSignal
}

// ChangeGaindB modifies the signals gain using a value in decibels as input
func ChangeGaindB(inputSignal types.Signal, dBChange float64) types.Signal {
	var scale = math.Pow(10, dBChange/20)

	outputSignal := ChangeGain(inputSignal, scale)

	return outputSignal
}

// InvertPolarity computes a polarly inverted signal of a mono signal
func InvertPolarity(inputSignal types.Signal) types.Signal {
	outputSignal := ChangeGain(inputSignal, -1)

	return outputSignal
}
