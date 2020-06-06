package distortion

import (
	"testing"

	"github.com/felixrosatmetlla/audiodsplib/audiodsputils"
)

func TestInfiniteClipping(t *testing.T) {
	testData := []struct {
		inputSignal   []float64
		clippedSignal []float64
	}{
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			[]float64{1, 1, 1, 1, 1, -1, -1, -1, 1},
		},
		{
			[]float64{0, -0.5, -2, -1.5, 0, 0.5, 2, 1.5, 0},
			[]float64{1, -1, -1, -1, 1, 1, 1, 1, 1},
		},
	}

	for _, caseData := range testData {
		result := InfiniteClipping(caseData.inputSignal)

		if !audiodsputils.CompareMonoSignals(result, caseData.clippedSignal) {
			t.Errorf("Infinite clipping of signal %v was incorrect, got: %v, want: %v.", caseData.inputSignal, result, caseData.clippedSignal)
		}
	}
}

func TestHardClipping(t *testing.T) {
	testData := []struct {
		inputSignal   []float64
		threshold     float64
		clippedSignal []float64
	}{
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			0.4,
			[]float64{0, 0.4, 0.4, 0.4, 0, -0.4, -0.4, -0.4, 0},
		},
		{
			[]float64{0, -0.5, -2, -1.5, 0, 0.5, 2, 1.5, 0},
			1.75,
			[]float64{0, -0.5, -1.75, -1.5, 0, 0.5, 1.75, 1.5, 0},
		},
	}

	for _, caseData := range testData {
		result := HardClipping(caseData.inputSignal, caseData.threshold)

		if !audiodsputils.CompareMonoSignals(result, caseData.clippedSignal) {
			t.Errorf("Hard clipping of signal %v was incorrect, got: %v, want: %v.", caseData.inputSignal, result, caseData.clippedSignal)
		}
	}
}

func TestCubicDistortion(t *testing.T) {
	testData := []struct {
		inputSignal  []float64
		amplitude    float64
		outputSignal []float64
	}{
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			0.5,
			[]float64{0, 0.4791666666666667, 0.8333333333333334, 0.4791666666666667, 0, -0.4791666666666667, -0.8333333333333334, -0.4791666666666667, 0},
		},
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			1,
			[]float64{0, 0.4583333333333333, 0.6666666666666667, 0.4583333333333333, 0, -0.4583333333333333, -0.6666666666666667, -0.4583333333333333, 0},
		},
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			0,
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
		},
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			-0.5,
			[]float64{},
		},
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			2,
			[]float64{},
		},
	}

	for _, caseData := range testData {
		result, err := CubicDistortion(caseData.inputSignal, caseData.amplitude)

		if len(result) == 0 && err == nil {
			t.Errorf("Error message informing of operation failure was expected and got: %v", err)
		}

		if len(result) != 0 && err != nil {
			t.Errorf("No error message was expected, and got %v", err)
		}

		if !audiodsputils.CompareMonoSignals(result, caseData.outputSignal) {
			t.Errorf("Cubic distortion of signal %v was incorrect, got: %v, want: %v.", caseData.inputSignal, result, caseData.outputSignal)
		}
	}
}

func TestArctangentDistortion(t *testing.T) {
	testData := []struct {
		inputSignal  []float64
		alpha        float64
		outputSignal []float64
	}{
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			10,
			[]float64{0, 0.8743340836219976, 0.9365489651388929, 0.8743340836219976, 0, -0.8743340836219976, -0.9365489651388929, -0.8743340836219976, 0},
		},
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			1,
			[]float64{0, 0.2951672353008665, 0.5, 0.2951672353008665, 0, -0.2951672353008665, -0.5, -0.2951672353008665, 0},
		},
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			0,
			[]float64{},
		},
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			11,
			[]float64{},
		},
	}

	for _, caseData := range testData {
		result, err := ArctangentDistortion(caseData.inputSignal, caseData.alpha)

		if len(result) == 0 && err == nil {
			t.Errorf("Error message informing of operation failure was expected and got: %v", err)
		}

		if len(result) != 0 && err != nil {
			t.Errorf("No error message was expected, and got %v", err)
		}

		if !audiodsputils.CompareMonoSignals(result, caseData.outputSignal) {
			t.Errorf("Arctangent distortion of signal %v was incorrect, got: %v, want: %v.", caseData.inputSignal, result, caseData.outputSignal)
		}
	}
}
