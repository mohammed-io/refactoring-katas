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

    let t = (x - d) * 0.07;
    let total = x - d + t;

    let payment = { status: 'approved' };
    if (total > 5000) {
      payment.status = 'manual_review';
    }

    let ship = { carrier: 'USPS' };
    if (total > 50) {
      ship.carrier = 'UPS';
    }

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

    return {
      orderId: order.id,
      total: Math.round(total * 100) / 100,
      paymentStatus: payment.status,
      carrier: ship.carrier,
      email: email,
      log: log
    };
  }
}

export { LegacySystem };
