class LoyaltyRules {
  constructor() {}

  get_discount_for_tier(tier) {
    if (tier === 'bronze') {
      return 0.05;
    } else if (tier === 'silver') {
      return 0.1;
    } else if (tier === 'gold') {
      return 0.15;
    } else if (tier === 'platinum') {
      return 0.2;
    }
    return 0;
  }

  get_label_for_tier(tier) {
    if (tier === 'bronze') {
      return 'Bronze Member';
    } else if (tier === 'silver') {
      return 'Silver Member';
    } else if (tier === 'gold') {
      return 'Gold Member';
    } else if (tier === 'platinum') {
      return 'Platinum Member';
    }
    return 'Standard';
  }

  get_threshold_for_tier(tier) {
    if (tier === 'bronze') {
      return 100;
    } else if (tier === 'silver') {
      return 500;
    } else if (tier === 'gold') {
      return 2000;
    } else if (tier === 'platinum') {
      return 10000;
    }
    return 0;
  }

  get_color_for_tier(tier) {
    if (tier === 'bronze') {
      return '#CD7F32';
    } else if (tier === 'silver') {
      return '#C0C0C0';
    } else if (tier === 'gold') {
      return '#FFD700';
    } else if (tier === 'platinum') {
      return '#E5E4E2';
    }
    return '#000000';
  }

  calculate_total(spending, tier) {
    let discount = 0;
    if (tier === 'bronze') {
      discount = 0.05;
    } else if (tier === 'silver') {
      discount = 0.1;
    } else if (tier === 'gold') {
      discount = 0.15;
    } else if (tier === 'platinum') {
      discount = 0.2;
    }
    return spending * (1 - discount);
  }
}

export { LoyaltyRules };
