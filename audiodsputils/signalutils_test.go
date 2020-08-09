package audiodsputils

import (
	"testing"

	"github.com/felixrosatmetlla/audiodsplib/audiodsputils"
	"github.com/felixrosatmetlla/audiodsplib/types"
)

func TestCreateSignal(t *testing.T) {
	testData := []struct {
		data         []float64
		channels     int
		samplerate   float64
		numSamples   int
		outputSignal types.Signal
	}{
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			1,
			44100.0,
			9,
			types.Signal{
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
			types.Signal{
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
