package distortion

import (
	"audiodsplib/audiodsputils"
	"testing"
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
