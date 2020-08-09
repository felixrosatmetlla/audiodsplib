package types

import (
	"testing"

	"github.com/felixrosatmetlla/audiodsplib/audiodsputils"
)

func TestCreateSignal(t *testing.T) {
	testData := []struct {
		data         []float64
		channels     int
		samplerate   float64
		numSamples   int
		outputSignal Signal
	}{
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			1,
			44100.0,
			9,
			Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
		},
		{
			[]float64{0, -0.5, -1, -0.5, 0, 0.5, 1, 0.5, 0},
			-1,
			44100.0,
			9,
			Signal{
				Data:       []float64{},
				Channels:   -1,
				Samplerate: 44100.0,
				NumSamples: 0,
			},
		},
	}

	for _, caseData := range testData {
		result, err := CreateSignal(caseData.data, caseData.channels, caseData.samplerate, caseData.numSamples)

		if result.NumSamples == 0 && err == nil {
			t.Errorf("Error message informing of operation failure was expected and got: %v", err)
		}

		if result.NumSamples != 0 && err != nil {
			t.Errorf("No error message was expected, and got %v", err)
		}

		if !audiodsputils.CompareSignals(result, caseData.outputSignal) {
			t.Errorf("Creation of Signal was incorrect, got: %v, want: %v.", result, caseData.outputSignal)
		}
	}
}
