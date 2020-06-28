package distortion

import (
	"testing"

	"github.com/felixrosatmetlla/audiodsplib/audiodsputils"
)

func TestBitReduction(t *testing.T) {
	testData := []struct {
		inputSignal  []float64
		numberBits   int
		outputSignal []float64
	}{
		{
			[]float64{0, 0.25, 0.5, 0.750, 1, 0.750, 0.5, 0.25, 0, -0.25, -0.5, -0.75, -1, -0.75, -0.5, -0.25, 0},
			1,
			[]float64{0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, -1, -1, -1, 0, 0, 0},
		},
		{
			[]float64{0, 0.25, 0.5, 0.750, 1, 0.750, 0.5, 0.25, 0, -0.25, -0.5, -0.75, -1, -0.75, -0.5, -0.25, 0},
			8,
			[]float64{0, 0.25, 0.5, 0.750, 1, 0.750, 0.5, 0.25, 0, -0.25, -0.5, -0.75, -1, -0.75, -0.5, -0.25, 0},
		},
		{
			[]float64{0, 0.25, 0.5, 0.750, 1, 0.750, 0.5, 0.25, 0, -0.25, -0.5, -0.75, -1, -0.75, -0.5, -0.25, 0},
			-1,
			[]float64{},
		},
	}

	for _, caseData := range testData {
		result, err := BitReduction(caseData.inputSignal, caseData.numberBits)

		if len(result) == 0 && err == nil {
			t.Errorf("Error message informing of operation failure was expected and got: %v", err)
		}

		if len(result) != 0 && err != nil {
			t.Errorf("No error message was expected, and got %v", err)
		}

		if !audiodsputils.CompareArrayValues(result, caseData.outputSignal) {
			t.Errorf("Bit reduction of signal %v was incorrect, got: %v, want: %v.", caseData.inputSignal, result, caseData.outputSignal)
		}
	}
}
