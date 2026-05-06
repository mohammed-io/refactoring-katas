class KitchenTicket {
  constructor(items, customer, table, special) {
    this.items = items;
    this.customer = customer;
    this.table = table;
    this.special = special;
  }

  print_ticket() {
    let lines = [];
    lines.push('Table: ' + this.table);
    lines.push('Customer: ' + this.customer);
    for (let i = 0; i < this.items.length; i++) {
      lines.push(this.items[i].name + ' x' + this.items[i].qty);
    }
    if (this.special && this.special.length > 0) {
      lines.push('Special: ' + this.special);
    }
    lines.push('---');
    return lines.join('\n');
  }
}

export { KitchenTicket };
