package distortion

import (
	"errors"
	"math"

	"github.com/felixrosatmetlla/audiodsplib/types"
)

// BitReduction reduces the signals bit resolution
//
// Input:
//  inputSignal: Input signal to reduce bits resolution
//  numberBits: Number of bits to reduce to
// Output:
//  Signal: the output signal with the modifications
//  error: has a value if numberBits has an invalid value
func BitReduction(inputSignal types.Signal, numberBits int) (types.Signal, error) {
	var outputSignal types.Signal

	if numberBits < 0 {
		outputSignal = types.Signal{
			Data:       []float64{},
			Channels:   inputSignal.Channels,
			Samplerate: inputSignal.Samplerate,
			NumSamples: 0,
		}

		return outputSignal, errors.New("distortion: Invalid number of bits value. Valid value: numberBits >= 0")
	}

	var bufferSize = inputSignal.NumSamples * inputSignal.Channels
	var outputBuffer = make([]float64, bufferSize)

	bitsAmplitudeValue := math.Pow(2, float64(numberBits))

	for index, value := range inputSignal.Data {
		auxScaledInput := 0.5*value + 0.5
		scaledInput := bitsAmplitudeValue * auxScaledInput
		roundedInput := math.Round(scaledInput)

		auxOutput := roundedInput / bitsAmplitudeValue
		outputBuffer[index] = 2*auxOutput - 1
	}

	outputSignal = types.Signal{
		Data:       outputBuffer,
		Channels:   inputSignal.Channels,
		Samplerate: inputSignal.Samplerate,
		NumSamples: inputSignal.NumSamples,
	}

	return outputSignal, nil
}
