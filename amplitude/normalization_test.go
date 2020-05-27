package amplitude

import (
	"testing"

	"audiodsplib/audiodsputils"
)

//TODO: Check fixture packages to get more reliable testing results
func TestPeakNormalization(t *testing.T) {
	testData := []struct {
		inputSignal    []float64
		invertedSignal []float64
	}{
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
		},
		{
			[]float64{0, -0.5, -2, -1.5, 0, 0.5, 2, 1.5, 0},
			[]float64{0, -0.25, -1, -0.75, 0, 0.25, 1, 0.75, 0},
		},
	}

	for _, caseData := range testData {
		result := PeakNormalization(caseData.inputSignal)

		if !audiodsputils.CompareMonoSignals(result, caseData.invertedSignal) {
			t.Errorf("Peak normalization of signal %v was incorrect, got: %v, want: %v.", caseData.inputSignal, result, caseData.invertedSignal)
		}
	}

}
