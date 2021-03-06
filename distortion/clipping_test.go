package distortion

import (
	"testing"

	"github.com/felixrosatmetlla/audiodsplib/audiodsputils"
	"github.com/felixrosatmetlla/audiodsplib/types"
)

func TestInfiniteClipping(t *testing.T) {
	testData := []struct {
		inputSignal  types.Signal
		outputSignal types.Signal
	}{
		{
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			types.Signal{
				Data:       []float64{1, 1, 1, 1, 1, -1, -1, -1, 1},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
		},
		{
			types.Signal{
				Data:       []float64{0, -0.5, -2, -1.5, 0, 0.5, 2, 1.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			types.Signal{
				Data:       []float64{1, -1, -1, -1, 1, 1, 1, 1, 1},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
		},
	}

	for _, caseData := range testData {
		result := InfiniteClipping(caseData.inputSignal)

		if !audiodsputils.CompareSignals(result, caseData.outputSignal) {
			t.Errorf("Infinite clipping of signal %v was incorrect, got: %v, want: %v.", caseData.inputSignal, result, caseData.outputSignal)
		}
	}
}

func TestHardClipping(t *testing.T) {
	testData := []struct {
		inputSignal  types.Signal
		threshold    float64
		outputSignal types.Signal
	}{
		{
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			0.4,
			types.Signal{
				Data:       []float64{0, 0.4, 0.4, 0.4, 0, -0.4, -0.4, -0.4, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
		},
		{
			types.Signal{
				Data:       []float64{0, -0.5, -2, -1.5, 0, 0.5, 2, 1.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			1.75,
			types.Signal{
				Data:       []float64{0, -0.5, -1.75, -1.5, 0, 0.5, 1.75, 1.5, 0},
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
				Data:       []float64{},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 0,
			},
		},
	}

	for _, caseData := range testData {
		result, err := HardClipping(caseData.inputSignal, caseData.threshold)

		if result.NumSamples == 0 && err == nil {
			t.Errorf("Error message informing of operation failure was expected and got: %v", err)
		}

		if result.NumSamples != 0 && err != nil {
			t.Errorf("No error message was expected, and got %v", err)
		}

		if !audiodsputils.CompareSignals(result, caseData.outputSignal) {
			t.Errorf("Hard clipping of signal %v was incorrect, got: %v, want: %v.", caseData.inputSignal, result, caseData.outputSignal)
		}
	}
}

func TestCubicDistortion(t *testing.T) {
	testData := []struct {
		inputSignal  types.Signal
		amplitude    float64
		outputSignal types.Signal
	}{
		{
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			0.5,
			types.Signal{
				Data:       []float64{0, 0.4791666666666667, 0.8333333333333334, 0.4791666666666667, 0, -0.4791666666666667, -0.8333333333333334, -0.4791666666666667, 0},
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
			1,
			types.Signal{
				Data:       []float64{0, 0.4583333333333333, 0.6666666666666667, 0.4583333333333333, 0, -0.4583333333333333, -0.6666666666666667, -0.4583333333333333, 0},
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
			0,
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
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
				Data:       []float64{},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 0,
			},
		},
		{
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			2,
			types.Signal{
				Data:       []float64{},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 0,
			},
		},
	}

	for _, caseData := range testData {
		result, err := CubicDistortion(caseData.inputSignal, caseData.amplitude)

		if result.NumSamples == 0 && err == nil {
			t.Errorf("Error message informing of operation failure was expected and got: %v", err)
		}

		if result.NumSamples != 0 && err != nil {
			t.Errorf("No error message was expected, and got %v", err)
		}

		if !audiodsputils.CompareSignals(result, caseData.outputSignal) {
			t.Errorf("Cubic distortion of signal %v was incorrect, got: %v, want: %v.", caseData.inputSignal, result, caseData.outputSignal)
		}
	}
}

func TestArctangentDistortion(t *testing.T) {
	testData := []struct {
		inputSignal  types.Signal
		alpha        float64
		outputSignal types.Signal
	}{
		{
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			10,
			types.Signal{
				Data:       []float64{0, 0.8743340836219976, 0.9365489651388929, 0.8743340836219976, 0, -0.8743340836219976, -0.9365489651388929, -0.8743340836219976, 0},
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
			1,
			types.Signal{
				Data:       []float64{0, 0.2951672353008665, 0.5, 0.2951672353008665, 0, -0.2951672353008665, -0.5, -0.2951672353008665, 0},
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
			0,
			types.Signal{
				Data:       []float64{},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 0,
			},
		},
		{
			types.Signal{
				Data:       []float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 9,
			},
			11,
			types.Signal{
				Data:       []float64{},
				Channels:   1,
				Samplerate: 44100.0,
				NumSamples: 0,
			},
		},
	}

	for _, caseData := range testData {
		result, err := ArctangentDistortion(caseData.inputSignal, caseData.alpha)

		if result.NumSamples == 0 && err == nil {
			t.Errorf("Error message informing of operation failure was expected and got: %v", err)
		}

		if result.NumSamples != 0 && err != nil {
			t.Errorf("No error message was expected, and got %v", err)
		}

		if !audiodsputils.CompareSignals(result, caseData.outputSignal) {
			t.Errorf("Arctangent distortion of signal %v was incorrect, got: %v, want: %v.", caseData.inputSignal, result, caseData.outputSignal)
		}
	}
}
