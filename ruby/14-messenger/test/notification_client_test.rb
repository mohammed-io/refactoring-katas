# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/notification_client'

class NotificationClientTest < Minitest::Test
  def setup
    @client = NotificationClient.new
  end

  def test_sends_notification_through_layers
    result = @client.send({ message: 'Hello' })
    assert_equal 'sent', result[:status]
    assert_equal({ message: 'Hello' }, result[:payload])
  end

  def test_returns_sent_status
    result = @client.send({ message: 'Test' })
    assert_equal 'sent', result[:status]
  end

  def test_preserves_payload
    payload = { alert: 'Urgent', level: 3 }
    result = @client.send(payload)
    assert_equal payload, result[:payload]
  end
end
