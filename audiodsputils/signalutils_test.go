package audiodsputils

import (
	"testing"

	"github.com/felixrosatmetlla/audiodsplib/types"
)

func TestCompareSignals(t *testing.T) {
	testData := []struct {
		firstSignal     types.Signal
		secondSignal    types.Signal
		areSignalsEqual bool
	}{
		{
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			true,
		},
		{
			types.Signal{
				Data:       []float64{0, -0.5, -2, -1.5, 0, 0.5, 2, 1.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			types.Signal{
				Data:       []float64{0, -0.25, -1, -0.75, 0, 0.25, 1, 0.75, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			false,
		},
		{
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   2,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			false,
		},
		{
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 48000.0,
				NumSamples: 9,
			},
			false,
		},
		{
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 6,
			},
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			false,
		},
	}

	for _, caseData := range testData {
		result := CompareSignals(caseData.firstSignal, caseData.secondSignal)

		if result != caseData.areSignalsEqual {
			t.Errorf("Comparison of signals %v and %v was incorrect, got: %t, want: %t.", caseData.firstSignal, caseData.secondSignal, result, caseData.areSignalsEqual)
		}
	}
}

func TestIsSignalValid(t *testing.T) {
	testData := []struct {
		inputSignal   types.Signal
		isSignalValid bool
	}{
		{
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 48000.0,
				NumSamples: 9,
			},
			true,
		},
		{
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   0,
				Samplerate: 48000.0,
				NumSamples: 9,
			},
			false,
		},
		{
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: -48000.0,
				NumSamples: 9,
			},
			false,
		},
		{
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 48000.0,
				NumSamples: -9,
			},
			false,
		},
	}

	for _, caseData := range testData {
		result := IsSignalValid(caseData.inputSignal)

		if result != caseData.isSignalValid {
			t.Errorf("Validity check of signal %v was incorrect, expected %v and got %v", caseData.inputSignal, caseData.isSignalValid, result)
		}
	}
}
