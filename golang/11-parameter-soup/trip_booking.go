package kata

import "math/rand"

type Booking struct {
	Error        string
	Origin       string
	Destination  string
	Total        float64
	Class        string
	Confirmation string
}

type TripBooking struct{}

func NewTripBooking() *TripBooking {
	return &TripBooking{}
}

func (tb *TripBooking) book_trip(origin, destination, departureDate, returnDate, travelClass, meal, seat, loyalty string, insurance bool, promo string, flexible bool) Booking {
	_ = returnDate
	_ = meal
	_ = seat
	if origin == "" || destination == "" {
		return Booking{Error: "Missing route"}
	}
	if departureDate == "" {
		return Booking{Error: "Missing departure"}
	}
	base := 200.0
	if travelClass == "business" {
		base = 800
	} else if travelClass == "first" {
		base = 2000
	}
	discount := 0.0
	if promo == "SAVE20" {
		discount = 0.2
	} else if promo == "SAVE10" {
		discount = 0.1
	}
	total := base * (1 - discount)
	if insurance {
		total += 50
	}
	if flexible {
		total += 30
	}
	if len(loyalty) >= 4 && loyalty[:4] == "GOLD" {
		total -= 25
	}
	return Booking{Origin: origin, Destination: destination, Total: total, Class: travelClass, Confirmation: "BK-" + tb.fmtInt(rand.Intn(999999))}
}
func (tb *TripBooking) fmtInt(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	for n > 0 {
		s = string(rune('0'+n%10)) + s
		n /= 10
	}
	return s
}
