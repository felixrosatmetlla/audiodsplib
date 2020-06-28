package synthesis

import (
	"math"

	"github.com/felixrosatmetlla/audiodsplib/types"
)

// SynthSinus synthesizes a mono sinus signal
//
// Input variables:
//  frequency: signal frequency in Hz
//  amplitude: amplitude of the output signal
//  phase: initial phase offset of the signal
//  duration: signal duration in seconds
//  sampleRate: signal sample rate in samples/s
//  channels: number of channels of the output signal
func SynthSinus(frequency float64, amplitude float64, phase float64, duration float64, sampleRate float64, channels int) types.Signal {
	var numSamplesChannel int = int(duration * sampleRate)
	var numSamplesSignal int = numSamplesChannel * channels
	var sinBuffer = make([]float64, numSamplesSignal)
	var Ts = 1 / sampleRate

	for channel := 0; channel < channels; channel++ {
		for index := 0; index < numSamplesChannel; index++ {
			var sample int = index + channel*numSamplesChannel

			var time float64 = float64(sample) * Ts
			sinBuffer[sample] = amplitude * math.Sin(2*math.Pi*frequency*time+phase)
		}
	}

	sinus := types.Signal{
		Data:       sinBuffer,
		Channels:   channels,
		Samplerate: sampleRate,
		NumSamples: numSamplesChannel,
	}

	return sinus
}
