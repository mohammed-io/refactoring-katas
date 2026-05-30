class InventoryItem:
    def __init__(self, id, name, batch_number, cache_timestamp, row_id, quantity):
        self.id = id
        self.name = name
        self.batch_number = batch_number
        self.cache_timestamp = cache_timestamp
        self.row_id = row_id
        self.quantity = quantity

    def get_id(self):
        return self.id

    def set_id(self, v):
        self.id = v

    def get_name(self):
        return self.name

    def set_name(self, v):
        self.name = v

    def get_batch_number(self):
        return self.batch_number

    def set_batch_number(self, v):
        self.batch_number = v

    def get_cache_timestamp(self):
        return self.cache_timestamp

    def set_cache_timestamp(self, v):
        self.cache_timestamp = v

    def get_row_id(self):
        return self.row_id

    def set_row_id(self, v):
        self.row_id = v

    def get_quantity(self):
        return self.quantity

    def set_quantity(self, v):
        self.quantity = v

    def reserve(self, units):
        if units <= 0:
            return {"status": "rejected", "reason": "invalid_quantity", "remaining": self.quantity}
        if units > self.quantity:
            return {"status": "backorder", "reserved": 0, "remaining": self.quantity}

        self.quantity -= units
        return {
            "status": "reserved",
            "reserved": units,
            "remaining": self.quantity,
            "sku": f"{self.id}-{self.batch_number}",
        }

    def receive_stock(self, units):
        if units <= 0:
            return self.quantity

        self.quantity += units
        return self.quantity

    def public_snapshot(self):
        return {
            "id": self.id,
            "name": self.name,
            "batch_number": self.batch_number,
            "quantity": self.quantity,
            "stock_status": self.stock_status(),
        }

    def stock_status(self):
        if self.quantity == 0:
            return "out"
        if self.quantity < 5:
            return "low"
        return "available"
