# frozen_string_literal: true

require 'json'
class ConfigLoader
  def load_config
    begin
      local = JSON.parse(File.read('/tmp/config.json'), symbolize_names: true)
    rescue StandardError
      local = {}
    end
    remote = {}
    seasonal = {}
    month = Time.now.month - 1
    seasonal = { theme: 'winter', discount: 0.1 } if [11, 0, 1].include?(month)
    seasonal = { theme: 'summer', discount: 0.05 } if month >= 5 && month <= 7
    remote.merge(local).merge(seasonal)
  end
end
