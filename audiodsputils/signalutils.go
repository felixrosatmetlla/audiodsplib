package audiodsputils

import (
	"github.com/felixrosatmetlla/audiodsplib/types"
)

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

//TODO: Method(Signal, output N-D slice(empty)) return slice?

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
