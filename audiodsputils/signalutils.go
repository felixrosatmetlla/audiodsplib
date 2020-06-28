package audiodsputils

import (
	"github.com/felixrosatmetlla/audiodsplib/types"
)

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
