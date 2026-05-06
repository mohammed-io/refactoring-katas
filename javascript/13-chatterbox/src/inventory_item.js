class InventoryItem {
  constructor(id, name, batchNumber, cacheTimestamp, rowId, quantity) {
    this.id = id;
    this.name = name;
    this.batchNumber = batchNumber;
    this.cacheTimestamp = cacheTimestamp;
    this.rowId = rowId;
    this.quantity = quantity;
  }

  get_id() {
    return this.id;
  }

  set_id(val) {
    this.id = val;
  }

  get_name() {
    return this.name;
  }

  set_name(val) {
    this.name = val;
  }

  get_batch_number() {
    return this.batchNumber;
  }

  set_batch_number(val) {
    this.batchNumber = val;
  }

  get_cache_timestamp() {
    return this.cacheTimestamp;
  }

  set_cache_timestamp(val) {
    this.cacheTimestamp = val;
  }

  get_row_id() {
    return this.rowId;
  }

  set_row_id(val) {
    this.rowId = val;
  }

  get_quantity() {
    return this.quantity;
  }

  set_quantity(val) {
    this.quantity = val;
  }
}

export { InventoryItem };
