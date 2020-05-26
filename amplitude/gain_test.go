package amplitude

import (
	"testing"

	"../audiodsputils"
)

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

		if !audiodsputils.CompareMonoSignals(result, caseData.invertedSignal) {
			t.Errorf("Inversion of signal %v was incorrect, got: %v, want: %v.", caseData.inputSignal, result, caseData.invertedSignal)
		}
	}
}
