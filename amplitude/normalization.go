package amplitude

import (
	"math"

	"github.com/felixrosatmetlla/audiodsplib/audiodsputils"
	"github.com/felixrosatmetlla/audiodsplib/types"
)

// PeakNormalization normalizes a signal using its peak value
func PeakNormalization(inputSignal types.Signal) types.Signal {
	var bufferSize = inputSignal.NumSamples * inputSignal.Channels
	var outputBuffer = make([]float64, bufferSize)

	var _, maxValue = audiodsputils.GetArrayMinMax(inputSignal.Data)
	for index := range outputBuffer {
		outputBuffer[index] = inputSignal.Data[index] / maxValue
	}

	outputSignal := types.Signal{
		Data:       outputBuffer,
		Channels:   inputSignal.Channels,
		Samplerate: inputSignal.Samplerate,
		NumSamples: inputSignal.NumSamples,
	}

	return outputSignal
}

//TODO: Check if samplesSquaredSum is for channel or total

// RMSNormalization normalizes a signal amplitude to a specified RMS (Root Mean Squared) amplitude value
func RMSNormalization(signal []float64, rmsAmplitude float64) []float64 {
	var signalLength = len(signal)
	var output = make([]float64, signalLength)

	var samplesSquaredSum float64 = 0
	for index := range signal {
		samplesSquaredSum = samplesSquaredSum + math.Pow(signal[index], 2)
	}

	var scale = math.Sqrt((float64(signalLength) * rmsAmplitude) / samplesSquaredSum)

	for index := range output {
		output[index] = signal[index] * scale
	}

	return output
}
