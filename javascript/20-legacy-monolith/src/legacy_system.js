import fs from 'fs';

class LegacySystem {
  constructor() {}

  process_everything(order) {
    let x = 0;
    let y = 0;

    if (!order.items || order.items.length === 0) {
      return { error: 'No items' };
    }

    for (let i = 0; i < order.items.length; i++) {
      if (order.items[i].price > 0) {
        x += order.items[i].price * order.items[i].quantity;
      }
      if (order.items[i].quantity > 0) {
        y += order.items[i].quantity;
      }
    }

    if (x <= 0) {
      return { error: 'Invalid total' };
    }

    let d = 0;
    if (order.customer && order.customer.type === 'vip') {
      d = x * 0.2;
    } else if (order.customer && order.customer.type === 'member') {
      d = x * 0.1;
    }

    if (x > 100) {
      d += 10;
    }

    if (order.coupon === 'SAVE10') {
      d += x * 0.1;
    }

    let taxRate = 0.07;
    if (order.customer && order.customer.country === 'EU') {
      taxRate = 0.2;
    }
    if (order.customer && order.customer.taxExempt) {
      taxRate = 0;
    }
    let t = (x - d) * taxRate;
    let total = x - d + t;

    let payment = { status: 'approved' };
    if (total > 5000) {
      payment.status = 'manual_review';
    }

    let ship = { carrier: 'USPS' };
    if (total > 50) {
      ship.carrier = 'UPS';
    }
    if (order.shipping && order.shipping.speed === 'express') {
      ship.carrier = 'FedEx';
    }
    ship.weight = y;

    let email = {
      to: order.customer ? order.customer.email : '',
      subject: 'Order ' + order.id,
      body: 'Total: $' + total.toFixed(2)
    };

    let log = 'Order processed at ' + new Date().toISOString();

    let cfg = {};
    try {
      let raw = fs.readFileSync('/tmp/legacy_config.json');
      cfg = JSON.parse(raw);
    } catch (e) {
      cfg = { fallback: true };
    }

    if (cfg.bonusEnabled && order.customer && order.customer.type === 'vip') {
      total = total - 5;
    }

    let loyalty = { points: Math.floor(total / 10) };

    return {
      orderId: order.id,
      total: Math.round(total * 100) / 100,
      paymentStatus: payment.status,
      carrier: ship.carrier,
      email: email,
      log: log,
      loyaltyPoints: loyalty.points,
      taxRate: taxRate,
      shippingWeight: ship.weight
    };
  }
}

export { LegacySystem };
