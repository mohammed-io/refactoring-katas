package kata

import (
	"math"
	"testing"
)

func slicesEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.IsNaN(a[i]) && math.IsNaN(b[i]) {
			continue
		}
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestNormalizeNormalizesScoresTo0100Range(t *testing.T) {
	calc := NewCalculator()
	result := calc.normalize([]float64{10, 20, 30}, 0, 100)
	expected := []float64{0, 50, 100}
	if !slicesEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestNormalizeNormalizesScoresTo15Range(t *testing.T) {
	calc := NewCalculator()
	result := calc.normalize([]float64{10, 20, 30}, 1, 5)
	expected := []float64{1, 3, 5}
	if !slicesEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestNormalizeHandlesSingleValue(t *testing.T) {
	calc := NewCalculator()
	result := calc.normalize([]float64{50}, 0, 100)
	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}
	if !math.IsNaN(result[0]) {
		t.Errorf("expected NaN for single value, got %v", result[0])
	}
}

func TestNormalizeHandlesNegativeInputRange(t *testing.T) {
	calc := NewCalculator()
	result := calc.normalize([]float64{-10, 0, 10}, 0, 1)
	expected := []float64{0, 0.5, 1}
	if !slicesEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestNormalizeHandlesSameMinAndMax(t *testing.T) {
	calc := NewCalculator()
	result := calc.normalize([]float64{5, 5, 5}, 0, 100)
	if len(result) != 3 {
		t.Fatalf("expected 3 results, got %d", len(result))
	}
	if !math.IsNaN(result[0]) {
		t.Errorf("expected NaN at index 0, got %v", result[0])
	}
	if !math.IsNaN(result[1]) {
		t.Errorf("expected NaN at index 1, got %v", result[1])
	}
	if !math.IsNaN(result[2]) {
		t.Errorf("expected NaN at index 2, got %v", result[2])
	}
}
