class Order {
  constructor(items, customer, table, special, rush = false) {
    this.items = items;
    this.customer = customer;
    this.table = table;
    this.special = special;
    this.rush = rush;
  }
}

class KitchenTicket {
  print_ticket(order) {
    let lines = [];
    lines.push('Table: ' + order.table.number);
    lines.push('Zone: ' + (order.table.zone || 'main'));
    lines.push('Server: ' + (order.table.server || 'unassigned'));
    lines.push('Customer: ' + order.customer.name);
    if (order.customer.vip) lines.push('VIP');
    if (order.rush) lines.push('RUSH');
    let totalItems = 0;
    for (let i = 0; i < order.items.length; i++) {
      totalItems += order.items[i].qty;
      let line = order.items[i].name + ' x' + order.items[i].qty;
      if (order.items[i].modifiers && order.items[i].modifiers.length > 0) {
        line += ' [' + order.items[i].modifiers.join(', ') + ']';
      }
      if (order.items[i].allergy) {
        line += ' ALLERGY:' + order.items[i].allergy;
      }
      lines.push(line);
    }
    lines.push('Items: ' + totalItems);
    if (order.special && order.special.length > 0) {
      lines.push('Special: ' + order.special);
    }
    lines.push('---');
    return lines.join('\n');
  }
}

export { KitchenTicket, Order };
