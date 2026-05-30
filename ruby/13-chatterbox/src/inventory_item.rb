# frozen_string_literal: true

class InventoryItem
  attr_accessor :id, :name, :batch_number, :cache_timestamp, :row_id, :quantity

  def initialize(id, name, batch_number, cache_timestamp, row_id, quantity)
    @id = id
    @name = name
    @batch_number = batch_number
    @cache_timestamp = cache_timestamp
    @row_id = row_id
    @quantity = quantity
  end

  def get_id
    @id
  end

  def set_id(v)
    @id = v
  end

  def get_name
    @name
  end

  def set_name(v)
    @name = v
  end

  def get_batch_number
    @batch_number
  end

  def set_batch_number(v)
    @batch_number = v
  end

  def get_cache_timestamp
    @cache_timestamp
  end

  def set_cache_timestamp(v)
    @cache_timestamp = v
  end

  def get_row_id
    @row_id
  end

  def set_row_id(v)
    @row_id = v
  end

  def get_quantity
    @quantity
  end

  def set_quantity(v)
    @quantity = v
  end

  def reserve(units)
    return { status: 'rejected', reason: 'invalid_quantity', remaining: @quantity } if units <= 0
    return { status: 'backorder', reserved: 0, remaining: @quantity } if units > @quantity

    @quantity -= units
    { status: 'reserved', reserved: units, remaining: @quantity, sku: "#{@id}-#{@batch_number}" }
  end

  def receive_stock(units)
    return @quantity if units <= 0

    @quantity += units
  end

  def public_snapshot
    {
      id: @id,
      name: @name,
      batch_number: @batch_number,
      quantity: @quantity,
      stock_status: stock_status
    }
  end

  def stock_status
    return 'out' if @quantity.zero?
    return 'low' if @quantity < 5

    'available'
  end
end
