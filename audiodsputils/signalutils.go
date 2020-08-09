package audiodsputils

import (
	"errors"

	"github.com/felixrosatmetlla/audiodsplib/types"
)

// CreateSignal is a constructor to create safely a Signal type instance
//
// Input variables:
//  data: samples of the signal
//  	  represented with a 1D slice where all channels data is put consecutively in order
//  channels: number of the channels the signal has
//  samplerate: signal samplerate in samples/s
//  numSamples: Number of samples per channel
func CreateSignal(data []float64, channels int, samplerate float64, numSamples int) (types.Signal, error) {
	var outputSignal = types.Signal{
		Data:       data,
		Channels:   channels,
		Samplerate: samplerate,
		NumSamples: numSamples,
	}

	if !IsSignalValid(outputSignal) {
		outputSignal = types.Signal{
			Data:       []float64{},
			Channels:   channels,
			Samplerate: samplerate,
			NumSamples: 0,
		}

		return outputSignal, errors.New("Signal: Invalid parameters to create a signal")
	}

	outputSignal = types.Signal{
		Data:       data,
		Channels:   channels,
		Samplerate: samplerate,
		NumSamples: numSamples,
	}

	return outputSignal, nil
}

// CompareSignals compares two Signal type variables
//
// Input:
//  signalA: First signal to compare
//  signalB: Second signal to compare
// Output:
//  result: boolean with the comparison, if Signals are equal returns true
func CompareSignals(signalA, signalB types.Signal) bool {

	//TODO: Think of returning message errors or error type
	if !CompareArrayValues(signalA.Data, signalB.Data) {
		return false
	}

	if !(signalA.Channels == signalB.Channels) {
		return false
	} else if !(signalA.NumSamples == signalB.NumSamples) {
		return false
	} else if !(signalA.Samplerate == signalB.Samplerate) {
		return false
	}

	return true
}

// IsSignalValid checks if a Signal type variable has valid field values
//
// Input:
//  signal: signal to check fields validity
// Output:
//  result: boolean with the result, if Signal is valid returns true
func IsSignalValid(signal types.Signal) bool {
	if signal.Channels < 1 {
		return false
	} else if signal.NumSamples < 0 {
		return false
	} else if signal.Samplerate < 0 {
		return false
	}

	return true
}
