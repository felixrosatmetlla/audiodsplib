package amplitude

import (
	"math"
	"testing"

	"github.com/felixrosatmetlla/audiodsplib/audiodsputils"
)

//TODO: Check fixture packages to get more reliable testing results
func TestPeakNormalization(t *testing.T) {
	testData := []struct {
		inputSignal      []float64
		normalizedSignal []float64
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

		if !audiodsputils.CompareMonoSignals(result, caseData.normalizedSignal) {
			t.Errorf("Peak normalization of signal %v was incorrect, got: %v, want: %v.", caseData.inputSignal, result, caseData.normalizedSignal)
		}
	}
}

func TestRMSNormalization(t *testing.T) {
	testData := []struct {
		inputSignal      []float64
		rmsAmplitude     float64
		normalizedSignal []float64
	}{
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			1,
			[]float64{0, 0.8660254037844386, 1.7320508075688772, 0.8660254037844386, 0, -0.8660254037844386, -1.7320508075688772, -0.8660254037844386, 0},
		},
		{
			[]float64{0, -0.5, -2, -1.5, 0, 0.5, 2, 1.5, 0},
			math.Sqrt(2) / 2,
			[]float64{0, -0.3498340549097454, -1.3993362196389816, -1.0495021647292362, 0, 0.3498340549097454, 1.3993362196389816, 1.0495021647292362, 0},
		},
	}

	for _, caseData := range testData {
		result := RMSNormalization(caseData.inputSignal, caseData.rmsAmplitude)

		if !audiodsputils.CompareMonoSignals(result, caseData.normalizedSignal) {
			t.Errorf("Peak normalization of signal %v was incorrect, got: %v, want: %v.", caseData.inputSignal, result, caseData.normalizedSignal)
		}
	}
}
