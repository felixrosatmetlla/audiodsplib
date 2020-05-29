package amplitude

import (
	"math"

	"github.com/felixrosatmetlla/audiodsplib/audiodsputils"
)

// PeakNormalization normalizes a signal using its peak value
func PeakNormalization(signal []float64) []float64 {
	var bufferSize = len(signal)
	var output = make([]float64, bufferSize)

	var _, maxValue = audiodsputils.GetArrayMinMax(signal)
	for index := range output {
		output[index] = signal[index] / maxValue
	}

	return output
}

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
