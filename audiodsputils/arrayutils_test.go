package audiodsputils

import (
	"testing"
)

func TestGetArrayMinMax(t *testing.T) {
	testData := []struct {
		inputArray   []float64
		minimumValue float64
		maximumValue float64
	}{
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			-1.0,
			1.0,
		},
		{
			[]float64{0, -0.5, -2, -1.5, 0, 0.5, 2, 1.5, 0},
			-2.0,
			2.0,
		},
	}

	for _, caseData := range testData {
		minResult, maxResult := GetArrayMinMax(caseData.inputArray)

		if minResult != caseData.minimumValue {
			t.Errorf("Minimum value of array %v was incorrect, got: %f, want: %f.", caseData.inputArray, minResult, caseData.minimumValue)
		}

		if maxResult != caseData.maximumValue {
			t.Errorf("Maximum value of array %v was incorrect, got: %f, want: %f.", caseData.inputArray, maxResult, caseData.maximumValue)
		}
	}
}

func TestCompareMonoSignals(t *testing.T) {
	testData := []struct {
		firstSignal     []float64
		secondSignal    []float64
		areSignalsEqual bool
	}{
		{
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			[]float64{0, 0.5, 1, 0.5, 0, -0.5, -1, -0.5, 0},
			true,
		},
		{
			[]float64{0, -0.5, -2, -1.5, 0, 0.5, 2, 1.5, 0},
			[]float64{0, -0.25, -1, -0.75, 0, 0.25, 1, 0.75, 0},
			false,
		},
		{
			[]float64{0, -0.5, -2, -1.5, 0, 0.5, 2, 1.5, 0},
			[]float64{0, -0.25, -1, -0.75},
			false,
		},
		{
			nil,
			[]float64{0, -0.25, -1, -0.75},
			false,
		},
		{
			nil,
			nil,
			true,
		},
	}

	for _, caseData := range testData {
		result := CompareMonoSignals(caseData.firstSignal, caseData.secondSignal)

		if result != caseData.areSignalsEqual {
			t.Errorf("Comparison of signals %v and %v was incorrect, got: %t, want: %t.", caseData.firstSignal, caseData.secondSignal, result, caseData.areSignalsEqual)
		}
	}
}
