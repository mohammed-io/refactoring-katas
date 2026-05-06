class OrderProcessor {
  constructor() {}

  process_order(order) {
    if (!order.items || order.items.length === 0) {
      return { error: 'No items' };
    }
    if (!order.customer || !order.customer.email) {
      return { error: 'Invalid customer' };
    }
    if (!order.address || !order.address.zip) {
      return { error: 'Invalid address' };
    }

    let inventory = true;
    for (let i = 0; i < order.items.length; i++) {
      if (order.items[i].quantity > 100) {
        inventory = false;
      }
    }
    if (!inventory) {
      return { error: 'Out of stock' };
    }

    let subtotal = 0;
    for (let i = 0; i < order.items.length; i++) {
      subtotal += order.items[i].price * order.items[i].quantity;
    }

    let shipping = 0;
    if (subtotal < 25) {
      shipping = 5.99;
    } else if (subtotal < 50) {
      shipping = 3.99;
    }

    let tax = subtotal * 0.07;
    let total = subtotal + tax + shipping;

    let paymentId = 'PAY-' + Math.floor(Math.random() * 1000000);
    let paymentStatus = 'approved';
    if (total > 1000) {
      paymentStatus = 'pending_review';
    }

    let label = {
      to: order.address,
      weight: order.items.reduce((sum, item) => sum + item.quantity, 0),
      carrier: 'USPS'
    };
    if (total > 100) {
      label.carrier = 'UPS';
    }

    let email = {
      to: order.customer.email,
      subject: 'Order Confirmation',
      body: 'Your order total is $' + total.toFixed(2)
    };

    return {
      orderId: 'ORD-' + Date.now(),
      paymentId: paymentId,
      paymentStatus: paymentStatus,
      total: Math.round(total * 100) / 100,
      shippingLabel: label,
      email: email
    };
  }
}

export { OrderProcessor };
