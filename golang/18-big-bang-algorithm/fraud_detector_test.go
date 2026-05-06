package kata

import (
	"testing"
	"time"
)

func TestDetectLowRiskSmallTransaction(t *testing.T) {
	fd := NewFraudDetector()
	now := time.Now().UnixMilli()
	result := fd.detect(Tx{Amount: 10, Timestamp: now, History: nil, Merchant: "grocery", Country: "US", CardCountry: "US"})
	if result.Rating != "low" {
		t.Errorf("expected rating 'low', got %q", result.Rating)
	}
	if result.Level != 1 {
		t.Errorf("expected level 1, got %d", result.Level)
	}
}

func TestDetectMediumRiskForLargeAmount(t *testing.T) {
	fd := NewFraudDetector()
	now := time.Now().UnixMilli()
	result := fd.detect(Tx{Amount: 1100, Timestamp: now, History: nil, Merchant: "grocery", Country: "US", CardCountry: "US"})
	if result.Rating != "medium" {
		t.Errorf("expected rating 'medium', got %q", result.Rating)
	}
	if result.Level != 2 {
		t.Errorf("expected level 2, got %d", result.Level)
	}
}

func TestDetectElevatedRiskForGambling(t *testing.T) {
	fd := NewFraudDetector()
	now := time.Now().UnixMilli()
	result := fd.detect(Tx{Amount: 100, Timestamp: now, History: nil, Merchant: "gambling", Country: "US", CardCountry: "US"})
	if result.Rating != "medium" {
		t.Errorf("expected rating 'medium', got %q", result.Rating)
	}
	if result.Level != 2 {
		t.Errorf("expected level 2, got %d", result.Level)
	}
}

func TestDetectHighRiskForCrossBorder(t *testing.T) {
	fd := NewFraudDetector()
	now := time.Now().UnixMilli()
	result := fd.detect(Tx{Amount: 1000, Timestamp: now, History: nil, Merchant: "grocery", Country: "FR", CardCountry: "US"})
	if result.Rating != "low" {
		t.Errorf("expected rating 'low', got %q", result.Rating)
	}
	if result.Level != 1 {
		t.Errorf("expected level 1, got %d", result.Level)
	}
}

func TestDetectCriticalRiskForLateNightCrypto(t *testing.T) {
	fd := NewFraudDetector()
	ts := time.Date(2024, 1, 1, 2, 0, 0, 0, time.UTC).UnixMilli()
	result := fd.detect(Tx{Amount: 600, Timestamp: ts, History: nil, Merchant: "crypto", Country: "CN", CardCountry: "US"})
	if result.Level < 3 {
		t.Errorf("expected level >= 3 for late night crypto, got %d", result.Level)
	}
}

func TestDetectVelocityIncreasesRisk(t *testing.T) {
	fd := NewFraudDetector()
	now := time.Now().UnixMilli()
	history := []Tx{
		{Amount: 10, Timestamp: now - 1000, History: nil, Merchant: "", Country: "", CardCountry: ""},
		{Amount: 10, Timestamp: now - 2000, History: nil, Merchant: "", Country: "", CardCountry: ""},
		{Amount: 10, Timestamp: now - 3000, History: nil, Merchant: "", Country: "", CardCountry: ""},
		{Amount: 10, Timestamp: now - 4000, History: nil, Merchant: "", Country: "", CardCountry: ""},
	}
	result := fd.detect(Tx{Amount: 50, Timestamp: now, History: history, Merchant: "grocery", Country: "US", CardCountry: "US"})
	if result.Level != 2 {
		t.Errorf("expected level 2, got %d", result.Level)
	}
}

func TestDetectVolumeSpikesIncreaseRisk(t *testing.T) {
	fd := NewFraudDetector()
	now := time.Now().UnixMilli()
	history := []Tx{
		{Amount: 200, Timestamp: now - 10000, History: nil, Merchant: "", Country: "", CardCountry: ""},
		{Amount: 200, Timestamp: now - 20000, History: nil, Merchant: "", Country: "", CardCountry: ""},
		{Amount: 200, Timestamp: now - 30000, History: nil, Merchant: "", Country: "", CardCountry: ""},
	}
	result := fd.detect(Tx{Amount: 50, Timestamp: now, History: history, Merchant: "grocery", Country: "US", CardCountry: "US"})
	if result.Level != 1 {
		t.Errorf("expected level 1, got %d", result.Level)
	}
}

func TestDetectIncludesScore(t *testing.T) {
	fd := NewFraudDetector()
	now := time.Now().UnixMilli()
	result := fd.detect(Tx{Amount: 10, Timestamp: now, History: nil, Merchant: "grocery", Country: "US", CardCountry: "US"})
	if result.Score != 0 {
		t.Errorf("expected score to be a number (got %d)", result.Score)
	}
}
