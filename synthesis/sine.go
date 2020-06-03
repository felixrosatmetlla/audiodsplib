package synthesis

import (
	"math"
)

// SynthMonoSinus synthesizes a mono sinus signal
// frequency: signal frequency in Hz
// phase: initial phase offset of the signal
// duration: signal duration in seconds
// sampleRate: signal sample rate in samples/s
func SynthMonoSinus(frequency float64, amplitude float64, phase float64, duration float64, sampleRate float64) []float64 {
	var samples int = int(duration * sampleRate)
	var sinBuffer = make([]float64, samples)
	var Ts = 1 / sampleRate

	for index := range sinBuffer {
		var time float64 = float64(index) * Ts
		sinBuffer[index] = amplitude * math.Sin(2*math.Pi*frequency*time+phase)
	}

	return sinBuffer
}
