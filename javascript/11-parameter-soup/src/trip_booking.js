class TripBooking {
  constructor() {}

  book_trip(origin, destination, departureDate, returnDate, travelClass, mealPreference, seatPreference, loyaltyNumber, insuranceFlag, promoCode, flexibleDatesFlag) {
    if (!origin || !destination) {
      return { error: 'Missing route' };
    }
    if (!departureDate) {
      return { error: 'Missing departure' };
    }

    let basePrice = 200;
    if (travelClass === 'business') {
      basePrice = 800;
    } else if (travelClass === 'first') {
      basePrice = 2000;
    }

    let discount = 0;
    if (promoCode === 'SAVE20') {
      discount = 0.2;
    } else if (promoCode === 'SAVE10') {
      discount = 0.1;
    }

    let total = basePrice * (1 - discount);
    if (insuranceFlag) {
      total += 50;
    }
    if (flexibleDatesFlag) {
      total += 30;
    }
    if (loyaltyNumber && loyaltyNumber.startsWith('GOLD')) {
      total -= 25;
    }

    return {
      origin: origin,
      destination: destination,
      total: Math.round(total * 100) / 100,
      class: travelClass,
      meal: mealPreference,
      seat: seatPreference,
      confirmation: 'BK-' + Math.floor(Math.random() * 1000000)
    };
  }
}

export { TripBooking };
