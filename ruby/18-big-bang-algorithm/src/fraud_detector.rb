# frozen_string_literal: true

class FraudDetector
  def detect(tx)
    s = 0
    v = 0
    m = 0
    d = Time.at(tx[:timestamp] / 1000).utc.hour
    tx[:history].each do |h|
      v += 1 if h[:amount] > tx[:amount] * 2
      s += 1 if (tx[:timestamp] - h[:timestamp]).abs < 3_600_000
    end
    m += 30 if tx[:amount] > 500 && d >= 0 && d < 6
    m += 20 if tx[:amount] > 1000
    m += 25 if %w[gambling crypto].include?(tx[:merchant])
    m += 15 if tx[:country] != tx[:card_country]
    m += 20 if s > 3
    m += 15 if v > 2
    level = if m < 20
              1
            elsif m < 40
              2
            elsif m < 60
              3
            else
              m < 80 ? 4 : 5
            end
    rating = { 1 => 'low', 2 => 'medium', 3 => 'elevated', 4 => 'high', 5 => 'critical' }[level]
    { score: m, level: level, rating: rating }
  end
end
