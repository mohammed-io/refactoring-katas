# frozen_string_literal: true

class Calculator
  def normalize(a, b, c)
    d = []
    e = Float::INFINITY
    f = -Float::INFINITY
    a.each do |x|
      e = x if x < e
      f = x if x > f
    end
    h = f - e
    i = c - b
    a.each do |x|
      k = b + ((x - e).to_f / h) * i
      d << k.round(2)
    end
    d
  end
end
