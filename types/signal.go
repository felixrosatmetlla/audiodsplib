package types

import (
	"errors"

	"github.com/felixrosatmetlla/audiodsplib/audiodsputils"
	"github.com/felixrosatmetlla/audiodsplib/types"
)

// Signal struct defines the type used to represent the signals
// and its propertys in the library
//
// Fields:
//  Data: samples of the signal
//  	  represented with a 1D slice where all channels data is put consecutively in order
//  Channels: number of the channels the signal has
//  Samplerate: signal samplerate in samples/s
type Signal struct {
	Data       []float64
	Channels   int
	Samplerate float64
	NumSamples int
}

// TODO: Make constructor method

func CreateSignal(data []float64, channels int, samplerate float64, numSamples int) (types.Signal, error) {
	var outputSignal = types.Signal{
		Data:       data,
		Channels:   channels,
		Samplerate: samplerate,
		NumSamples: numSamples,
	}

	if !audiodsputils.IsSignalValid(outputSignal) {
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
