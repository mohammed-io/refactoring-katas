# frozen_string_literal: true

class TripBooking
  def book_trip(origin, destination, departure_date, _return_date, travel_class, meal_preference, seat_preference,
                loyalty_number, insurance_flag, promo_code, flexible_dates_flag)
    return { error: 'Missing route' } if !origin || !destination

    return { error: 'Missing departure' } unless departure_date

    base = if travel_class == 'business'
             800
           else
             (travel_class == 'first' ? 2000 : 200)
           end

    discount = if promo_code == 'SAVE20'
                 0.2
               else
                 (promo_code == 'SAVE10' ? 0.1 : 0)
               end
    total = base * (1 - discount)

    total += 50 if insurance_flag
    total += 30 if flexible_dates_flag
    total -= 25 if loyalty_number&.start_with?('GOLD')
    { origin: origin, destination: destination, total: total.round(2), class: travel_class, meal: meal_preference,
      seat: seat_preference, confirmation: "BK-#{rand(1_000_000)}" }
  end
end
