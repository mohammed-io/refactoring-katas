import random

class TripBooking:
    def __init__(self):
        pass

    def book_trip(
        self,
        origin,
        destination,
        departure_date,
        return_date,
        travel_class,
        meal_preference,
        seat_preference,
        loyalty_number,
        insurance_flag,
        promo_code,
        flexible_dates_flag
    ):
        if not origin or not destination:
            return {"error": "Missing route"}

        if not departure_date:
            return {"error": "Missing departure"}

        base = 200
        if travel_class == "business":
            base = 800
        elif travel_class == "first":
            base = 2000

        discount = 0
        if promo_code == "SAVE20":
            discount = 0.2
        elif promo_code == "SAVE10":
            discount = 0.1

        total = base * (1 - discount)

        if insurance_flag:
            total += 50

        if flexible_dates_flag:
            total += 30

        if loyalty_number and loyalty_number.startswith("GOLD"):
            total -= 25

        return {
            "origin": origin,
            "destination": destination,
            "total": round(total, 2),
            "class": travel_class,
            "meal": meal_preference,
            "seat": seat_preference,
            "confirmation": "BK-" + str(random.randint(0, 999999))
        }
