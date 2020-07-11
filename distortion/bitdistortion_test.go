package distortion

import (
	"testing"

	"github.com/felixrosatmetlla/audiodsplib/audiodsputils"
	"github.com/felixrosatmetlla/audiodsplib/types"
)

func TestBitReduction(t *testing.T) {
	testData := []struct {
		inputSignal  types.Signal
		numberBits   int
		outputSignal types.Signal
	}{
		{
			types.Signal{
				Data:       []float64{0, 0.25, 0.5, 0.750, 1, 0.750, 0.5, 0.25, 0, -0.25, -0.5, -0.75, -1, -0.75, -0.5, -0.25, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 17,
			},
			1,
			types.Signal{
				Data:       []float64{0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, -1, -1, -1, 0, 0, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 17,
			},
		},
		{
			types.Signal{
				Data:       []float64{0, 0.25, 0.5, 0.750, 1, 0.750, 0.5, 0.25, 0, -0.25, -0.5, -0.75, -1, -0.75, -0.5, -0.25, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 17,
			},
			8,
			types.Signal{
				Data:       []float64{0, 0.25, 0.5, 0.750, 1, 0.750, 0.5, 0.25, 0, -0.25, -0.5, -0.75, -1, -0.75, -0.5, -0.25, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 17,
			},
		},
		{
			types.Signal{
				Data:       []float64{0, 0.25, 0.5, 0.750, 1, 0.750, 0.5, 0.25, 0, -0.25, -0.5, -0.75, -1, -0.75, -0.5, -0.25, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 17,
			},
			-1,
			types.Signal{
				Data:       []float64{},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 0,
			},
		},
	}

	for _, caseData := range testData {
		result, err := BitReduction(caseData.inputSignal, caseData.numberBits)

		if result.NumSamples == 0 && err == nil {
			t.Errorf("Error message informing of operation failure was expected and got: %v", err)
		}

		if result.NumSamples != 0 && err != nil {
			t.Errorf("No error message was expected, and got %v", err)
		}

		if !audiodsputils.CompareSignals(result, caseData.outputSignal) {
			t.Errorf("Bit reduction of signal %v was incorrect, got: %v, want: %v.", caseData.inputSignal, result, caseData.outputSignal)
		}
	}
}
