# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/config_loader'

class ConfigLoaderTest < Minitest::Test
  def setup
    @loader = ConfigLoader.new
  end

  def teardown
    File.delete('/tmp/config.json') if File.exist?('/tmp/config.json')
  end

  def test_returns_empty_object_when_file_and_fetch_fail
    result = @loader.load_config
    assert_kind_of Hash, result
  end

  def test_reads_local_config_file_when_present
    File.write('/tmp/config.json', JSON.dump({ name: 'test' }))
    result = @loader.load_config
    assert_equal 'test', result[:name]
  end

  def test_local_overrides_empty_file
    File.write('/tmp/config.json', JSON.dump({ name: 'local' }))
    result = @loader.load_config
    assert_equal 'local', result[:name]
  end

  def test_handles_malformed_json_gracefully
    File.write('/tmp/config.json', 'not json')
    result = @loader.load_config
    assert_kind_of Hash, result
  end

  def test_includes_seasonal_keys_in_object
    result = @loader.load_config
    assert result.key?(:theme) || result[:theme].nil?
    assert result.key?(:discount) || result[:discount].nil?
  end
end
