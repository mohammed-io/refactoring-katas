from src.trip_booking import TripBooking

def test_missing_origin():
    booking = TripBooking()
    result = booking.book_trip(None, "NYC", "2024-01-01", None, "economy", "vegan", "aisle", None, False, None, False)
    assert result["error"] == "Missing route"

def test_missing_destination():
    booking = TripBooking()
    result = booking.book_trip("LAX", None, "2024-01-01", None, "economy", "vegan", "aisle", None, False, None, False)
    assert result["error"] == "Missing route"

def test_missing_departure():
    booking = TripBooking()
    result = booking.book_trip("LAX", "NYC", None, None, "economy", "vegan", "aisle", None, False, None, False)
    assert result["error"] == "Missing departure"

def test_economy_price():
    booking = TripBooking()
    result = booking.book_trip("LAX", "NYC", "2024-01-01", None, "economy", "vegan", "aisle", None, False, None, False)
    assert result["total"] == 200
    assert result["class"] == "economy"

def test_business_price():
    booking = TripBooking()
    result = booking.book_trip("LAX", "NYC", "2024-01-01", None, "business", "kosher", "window", None, False, None, False)
    assert result["total"] == 800

def test_first_class_price():
    booking = TripBooking()
    result = booking.book_trip("LAX", "NYC", "2024-01-01", None, "first", "halal", "window", None, False, None, False)
    assert result["total"] == 2000

def test_save20_promo():
    booking = TripBooking()
    result = booking.book_trip("LAX", "NYC", "2024-01-01", None, "economy", "vegan", "aisle", None, False, "SAVE20", False)
    assert result["total"] == 160

def test_save10_promo():
    booking = TripBooking()
    result = booking.book_trip("LAX", "NYC", "2024-01-01", None, "economy", "vegan", "aisle", None, False, "SAVE10", False)
    assert result["total"] == 180

def test_insurance():
    booking = TripBooking()
    result = booking.book_trip("LAX", "NYC", "2024-01-01", None, "economy", "vegan", "aisle", None, True, None, False)
    assert result["total"] == 250

def test_flexible_dates():
    booking = TripBooking()
    result = booking.book_trip("LAX", "NYC", "2024-01-01", None, "economy", "vegan", "aisle", None, False, None, True)
    assert result["total"] == 230

def test_gold_loyalty():
    booking = TripBooking()
    result = booking.book_trip("LAX", "NYC", "2024-01-01", None, "economy", "vegan", "aisle", "GOLD123", False, None, False)
    assert result["total"] == 175

def test_includes_route():
    booking = TripBooking()
    result = booking.book_trip("LAX", "NYC", "2024-01-01", None, "economy", "vegan", "aisle", None, False, None, False)
    assert result["origin"] == "LAX"
    assert result["destination"] == "NYC"

def test_confirmation_code():
    booking = TripBooking()
    result = booking.book_trip("LAX", "NYC", "2024-01-01", None, "economy", "vegan", "aisle", None, False, None, False)
    assert result["confirmation"].startswith("BK-")
