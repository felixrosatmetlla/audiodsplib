package amplitude

import (
	"math"

	"github.com/felixrosatmetlla/audiodsplib/types"
)

// ChangeGain modifies the signals gain using a linear value of Amplitude
//
// Input:
//  inputSignal: Input signal to modify amplitude
//  gain: Value of gain that will change the signal
// Output:
//  Signal: the output signal with the modifications
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
//
// Input:
//  inputSignal: Input signal to modify amplitude
//  dBChange: Value of decibels that will change the input signal
// Output:
//  Signal: the output signal with the modifications
func ChangeGaindB(inputSignal types.Signal, dBChange float64) types.Signal {
	var scale = math.Pow(10, dBChange/20)

	outputSignal := ChangeGain(inputSignal, scale)

	return outputSignal
}

// InvertPolarity computes a polarly inverted signal of a mono signal
//
// Input:
//  inputSignal: Input signal to modify amplitude
// Output:
//  Signal: the output signal with the modifications
func InvertPolarity(inputSignal types.Signal) types.Signal {
	outputSignal := ChangeGain(inputSignal, -1)

	return outputSignal
}
