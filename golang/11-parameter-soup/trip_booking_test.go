package kata

import (
	"strings"
	"testing"
)

func TestBookTripRejectsMissingOrigin(t *testing.T) {
	tb := NewTripBooking()
	result := tb.book_trip("", "NYC", "2024-01-01", "", "economy", "vegan", "aisle", "", false, "", false)
	if result.Error != "Missing route" {
		t.Errorf("expected error 'Missing route', got %q", result.Error)
	}
}

func TestBookTripRejectsMissingDestination(t *testing.T) {
	tb := NewTripBooking()
	result := tb.book_trip("LAX", "", "2024-01-01", "", "economy", "vegan", "aisle", "", false, "", false)
	if result.Error != "Missing route" {
		t.Errorf("expected error 'Missing route', got %q", result.Error)
	}
}

func TestBookTripRejectsMissingDepartureDate(t *testing.T) {
	tb := NewTripBooking()
	result := tb.book_trip("LAX", "NYC", "", "", "economy", "vegan", "aisle", "", false, "", false)
	if result.Error != "Missing departure" {
		t.Errorf("expected error 'Missing departure', got %q", result.Error)
	}
}

func TestBookTripCalculatesEconomyPrice(t *testing.T) {
	tb := NewTripBooking()
	result := tb.book_trip("LAX", "NYC", "2024-01-01", "", "economy", "vegan", "aisle", "", false, "", false)
	if result.Total != 200 {
		t.Errorf("expected total 200, got %v", result.Total)
	}
	if result.Class != "economy" {
		t.Errorf("expected class 'economy', got %q", result.Class)
	}
}

func TestBookTripCalculatesBusinessPrice(t *testing.T) {
	tb := NewTripBooking()
	result := tb.book_trip("LAX", "NYC", "2024-01-01", "", "business", "kosher", "window", "", false, "", false)
	if result.Total != 800 {
		t.Errorf("expected total 800, got %v", result.Total)
	}
}

func TestBookTripCalculatesFirstClassPrice(t *testing.T) {
	tb := NewTripBooking()
	result := tb.book_trip("LAX", "NYC", "2024-01-01", "", "first", "halal", "window", "", false, "", false)
	if result.Total != 2000 {
		t.Errorf("expected total 2000, got %v", result.Total)
	}
}

func TestBookTripAppliesSAVE20Promo(t *testing.T) {
	tb := NewTripBooking()
	result := tb.book_trip("LAX", "NYC", "2024-01-01", "", "economy", "vegan", "aisle", "", false, "SAVE20", false)
	if result.Total != 160 {
		t.Errorf("expected total 160, got %v", result.Total)
	}
}

func TestBookTripAppliesSAVE10Promo(t *testing.T) {
	tb := NewTripBooking()
	result := tb.book_trip("LAX", "NYC", "2024-01-01", "", "economy", "vegan", "aisle", "", false, "SAVE10", false)
	if result.Total != 180 {
		t.Errorf("expected total 180, got %v", result.Total)
	}
}

func TestBookTripAddsInsurance(t *testing.T) {
	tb := NewTripBooking()
	result := tb.book_trip("LAX", "NYC", "2024-01-01", "", "economy", "vegan", "aisle", "", true, "", false)
	if result.Total != 250 {
		t.Errorf("expected total 250, got %v", result.Total)
	}
}

func TestBookTripAddsFlexibleDates(t *testing.T) {
	tb := NewTripBooking()
	result := tb.book_trip("LAX", "NYC", "2024-01-01", "", "economy", "vegan", "aisle", "", false, "", true)
	if result.Total != 230 {
		t.Errorf("expected total 230, got %v", result.Total)
	}
}

func TestBookTripAppliesGoldLoyaltyDiscount(t *testing.T) {
	tb := NewTripBooking()
	result := tb.book_trip("LAX", "NYC", "2024-01-01", "", "economy", "vegan", "aisle", "GOLD123", false, "", false)
	if result.Total != 175 {
		t.Errorf("expected total 175, got %v", result.Total)
	}
}

func TestBookTripIncludesRouteInResult(t *testing.T) {
	tb := NewTripBooking()
	result := tb.book_trip("LAX", "NYC", "2024-01-01", "", "economy", "vegan", "aisle", "", false, "", false)
	if result.Origin != "LAX" {
		t.Errorf("expected origin 'LAX', got %q", result.Origin)
	}
	if result.Destination != "NYC" {
		t.Errorf("expected destination 'NYC', got %q", result.Destination)
	}
}

func TestBookTripIncludesConfirmationCode(t *testing.T) {
	tb := NewTripBooking()
	result := tb.book_trip("LAX", "NYC", "2024-01-01", "", "economy", "vegan", "aisle", "", false, "", false)
	if !strings.HasPrefix(result.Confirmation, "BK-") {
		t.Errorf("expected confirmation to start with 'BK-', got %q", result.Confirmation)
	}
}
