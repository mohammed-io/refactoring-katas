# frozen_string_literal: true

require 'minitest/autorun'
require 'tmpdir'
require 'fileutils'
require_relative '../src/config_loader'

class ConfigLoaderTest < Minitest::Test
  def setup
    @loader = ConfigLoader.new
    @old_env = ENV.to_h
    @dir = Dir.mktmpdir
    ENV['APP_CONFIG_PATH'] = File.join(@dir, 'config.json')
    ENV.delete('APP_THEME')
    ENV.delete('APP_RETRIES')
    ENV.delete('APP_CURRENT_MONTH')
  end

  def teardown
    ENV.replace(@old_env)
    FileUtils.remove_entry(@dir)
  end

  def test_returns_defaults_when_file_is_missing
    result = @loader.load_config
    assert_equal 3, result[:retries]
    assert_equal 'standard', result[:theme]
    assert_equal 0, result[:discount]
  end

  def test_reads_local_config_file_when_present
    File.write(ENV.fetch('APP_CONFIG_PATH'), JSON.dump({ name: 'test', retries: 5 }))
    result = @loader.load_config
    assert_equal 'test', result[:name]
    assert_equal 5, result[:retries]
  end

  def test_environment_overrides_local_config
    File.write(ENV.fetch('APP_CONFIG_PATH'), JSON.dump({ theme: 'local', retries: 5 }))
    ENV['APP_THEME'] = 'env-theme'
    ENV['APP_RETRIES'] = '9'
    result = @loader.load_config
    assert_equal 'env-theme', result[:theme]
    assert_equal 9, result[:retries]
  end

  def test_handles_malformed_json_gracefully
    File.write(ENV.fetch('APP_CONFIG_PATH'), 'not json')
    result = @loader.load_config
    assert_equal 'standard', result[:theme]
  end

  def test_winter_seasonal_config_has_highest_precedence
    File.write(ENV.fetch('APP_CONFIG_PATH'), JSON.dump({ theme: 'local', discount: 0.25 }))
    ENV['APP_THEME'] = 'env-theme'
    ENV['APP_CURRENT_MONTH'] = '12'
    result = @loader.load_config
    assert_equal 'winter', result[:theme]
    assert_equal 0.1, result[:discount]
  end

  def test_summer_seasonal_config_is_deterministic
    ENV['APP_CURRENT_MONTH'] = '7'
    result = @loader.load_config
    assert_equal 'summer', result[:theme]
    assert_equal 0.05, result[:discount]
  end
end
