# frozen_string_literal: true

require 'json'
class ConfigLoader
  def load_config
    defaults = { retries: 3, theme: 'standard', discount: 0 }
    begin
      path = ENV['APP_CONFIG_PATH'] || '/tmp/config.json'
      local = JSON.parse(File.read(path), symbolize_names: true)
    rescue StandardError
      local = {}
    end
    env_config = {}
    env_config[:theme] = ENV['APP_THEME'] if ENV['APP_THEME']
    env_config[:retries] = ENV['APP_RETRIES'].to_i if ENV['APP_RETRIES']
    seasonal = {}
    month = (ENV['APP_CURRENT_MONTH'] || Time.now.month).to_i
    seasonal = { theme: 'winter', discount: 0.1 } if [12, 1, 2].include?(month)
    seasonal = { theme: 'summer', discount: 0.05 } if month >= 6 && month <= 8
    defaults.merge(local).merge(env_config).merge(seasonal)
  end
end
