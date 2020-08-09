package synthesis

import (
	"math"

	"github.com/felixrosatmetlla/audiodsplib/types"
)

// SynthSinus synthesizes a mono sinus signal
//
// Input:
//  frequency: signal frequency in Hz
//  amplitude: amplitude of the output signal
//  phase: initial phase offset of the signal
//  duration: signal duration in seconds
//  sampleRate: signal sample rate in samples/s
//  channels: number of channels of the output signal
// Output:
//  Signal: the output sinus Signal
func SynthSinus(frequency float64, amplitude float64, phase float64, duration float64, sampleRate float64, channels int) types.Signal {
	var numSamples int = int(duration * sampleRate)
	var bufferSize int = numSamples * channels
	var outputBuffer = make([]float64, bufferSize)
	var Ts = 1 / sampleRate

	for channel := 0; channel < channels; channel++ {
		for index := 0; index < numSamples; index++ {
			var sample int = index + channel*numSamples

			var time float64 = float64(sample) * Ts
			outputBuffer[sample] = amplitude * math.Sin(2*math.Pi*frequency*time+phase)
		}
	}

	sinus := types.Signal{
		Data:       outputBuffer,
		Channels:   channels,
		Samplerate: sampleRate,
		NumSamples: numSamples,
	}

	return sinus
}
