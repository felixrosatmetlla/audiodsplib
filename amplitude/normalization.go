package amplitude

import (
	"audiodsplib/audiodsputils"
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

//TODO: Implement RMS normalization
