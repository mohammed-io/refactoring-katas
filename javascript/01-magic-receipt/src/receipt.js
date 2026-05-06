class Receipt {
  constructor() {}

  calculate_total(items, customerType) {
    let total = 0;
    for (let i = 0; i < items.length; i++) {
      total += items[i];
    }

    let discount = 0;
    if (customerType === "member") {
      discount = total * 0.05;
    } else if (customerType === "vip") {
      discount = total * 0.15;
    }

    if (total > 50) {
      discount += 5;
    }

    let afterDiscount = total - discount;
    let tax = afterDiscount * 0.08;
    let final = afterDiscount + tax;

    if (customerType === "vip") {
      final -= 2;
    }

    return Math.round(final * 100) / 100;
  }
}

export { Receipt };
