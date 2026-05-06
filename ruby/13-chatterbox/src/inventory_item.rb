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
end
