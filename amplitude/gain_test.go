package amplitude

import (
	"math"
	"testing"

	"github.com/felixrosatmetlla/audiodsplib/audiodsputils"
	"github.com/felixrosatmetlla/audiodsplib/types"
)

func TestChangeGain(t *testing.T) {
	testData := []struct {
		inputSignal  types.Signal
		gain         float64
		outputSignal types.Signal
	}{
		{
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			2.0,
			types.Signal{
				Data:       []float64{0, 1, 2, 1, 0, -1, -2, -1, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
		},
		{
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			-0.5,
			types.Signal{
				Data:       []float64{0, -0.25, -0.5, -0.25, 0, 0.25, 0.5, 0.25, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
		},
	}

	for _, caseData := range testData {
		result := ChangeGain(caseData.inputSignal, caseData.gain)

		if !audiodsputils.CompareSignals(result, caseData.outputSignal) {
			t.Errorf("Gain change of signal %v was incorrect, got: %v, want: %v.", caseData.inputSignal, result, caseData.outputSignal)
		}
	}
}

func TestChangeGaindB(t *testing.T) {
	testData := []struct {
		inputSignal  []float64
		gainIndB     float64
		outputSignal []float64
	}{
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			6.0,
			[]float64{0, 0.5 * math.Pow(10, 6.0/20), 1 * math.Pow(10, 6.0/20), 0.5 * math.Pow(10, 6.0/20), 0, -0.5 * math.Pow(10, 6.0/20), -1 * math.Pow(10, 6.0/20), -0.5 * math.Pow(10, 6.0/20), 0},
		},
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			-3.0,
			[]float64{0, 0.5 * math.Pow(10, -3.0/20), 1 * math.Pow(10, -3.0/20), 0.5 * math.Pow(10, -3.0/20), 0, -0.5 * math.Pow(10, -3.0/20), -1 * math.Pow(10, -3.0/20), -0.5 * math.Pow(10, -3.0/20), 0},
		},
	}

	for _, caseData := range testData {
		result := ChangeGaindB(caseData.inputSignal, caseData.gainIndB)

		if !audiodsputils.CompareArrayValues(result, caseData.outputSignal) {
			t.Errorf("Gain change in dB of signal %v was incorrect, got: %v, want: %v.", caseData.inputSignal, result, caseData.outputSignal)
		}
	}
}

func TestInvertPolarity(t *testing.T) {
	testData := []struct {
		inputSignal    []float64
		invertedSignal []float64
	}{
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			[]float64{0, -0.5, -1, -0.5, 0, 0.5, 1, 0.5, 0},
		},
		{
			[]float64{},
			[]float64{},
		},
	}

	for _, caseData := range testData {
		result := InvertPolarity(caseData.inputSignal)

		if !audiodsputils.CompareArrayValues(result, caseData.invertedSignal) {
			t.Errorf("Inversion of signal %v was incorrect, got: %v, want: %v.", caseData.inputSignal, result, caseData.invertedSignal)
		}
	}
}
