package types

import (
	"errors"

	"github.com/felixrosatmetlla/audiodsplib/audiodsputils"
)

// Signal struct defines the type used to represent the signals
// and its propertys in the library
//
// Fields:
//  Data: samples of the signal
//  	  represented with a 1D slice where all channels data is put consecutively in order
//  Channels: number of the channels the signal has
//  Samplerate: signal samplerate in samples/s
//  NumSamples: Number of samples per channel
type Signal struct {
	Data       []float64
	Channels   int
	Samplerate float64
	NumSamples int
}

// CreateSignal is a constructor to create safely a Signal type instance
//
// Input variables:
//  data: samples of the signal
//  	  represented with a 1D slice where all channels data is put consecutively in order
//  channels: number of the channels the signal has
//  samplerate: signal samplerate in samples/s
//  numSamples: Number of samples per channel
func CreateSignal(data []float64, channels int, samplerate float64, numSamples int) (Signal, error) {
	var outputSignal = Signal{
		Data:       data,
		Channels:   channels,
		Samplerate: samplerate,
		NumSamples: numSamples,
	}

	if !audiodsputils.IsSignalValid(outputSignal) {
		outputSignal = Signal{
			Data:       []float64{},
			Channels:   channels,
			Samplerate: samplerate,
			NumSamples: 0,
		}

		return outputSignal, errors.New("Signal: Invalid parameters to create a signal")
	}

	outputSignal = Signal{
		Data:       data,
		Channels:   channels,
		Samplerate: samplerate,
		NumSamples: numSamples,
	}

	return outputSignal, nil
}
