# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/notification_client'

class NotificationClientTest < Minitest::Test
  def setup
    @client = NotificationClient.new
  end

  def test_sends_notification_through_layers
    result = @client.send({ recipient: 'user-1', message: 'Hello', channel: 'sms' })

    assert_equal 'sent', result[:status]
    assert_equal 'sms-user-1', result[:delivery_id]
    assert_equal 'Hello', result[:payload][:message]
  end

  def test_defaults_channel_and_priority
    result = @client.send({ message: 'Test' })

    assert_equal 'sent', result[:status]
    assert_equal 'email', result[:payload][:channel]
    assert_equal 'normal', result[:payload][:priority]
  end

  def test_preserves_explicit_priority
    payload = { recipient: 'ops', message: 'Urgent', priority: 'high' }
    result = @client.send(payload)

    assert_equal 'high', result[:payload][:priority]
    assert_equal 'ops', result[:payload][:recipient]
  end

  def test_rejects_missing_message
    result = @client.send({ recipient: 'ops' })

    assert_equal 'rejected', result[:status]
    assert_equal 'missing_message', result[:reason]
  end

  def test_records_observable_audit_events
    result = @client.send({ message: 'Deploy complete', channel: 'push' })

    assert_equal ['queued:push', 'sent:push'], result[:audit]
  end

  def test_reports_failed_delivery_for_unsupported_channel
    result = @client.send({ recipient: 'ops', message: 'Legacy alert', channel: 'fax' })

    assert_equal 'failed', result[:status]
    assert_equal 'unsupported_channel', result[:reason]
    assert_equal ['queued:fax', 'failed:fax'], result[:audit]
  end

  def test_high_priority_failed_delivery_is_scheduled_for_retry
    result = @client.send({ recipient: 'ops', message: 'Legacy alert', channel: 'fax', priority: 'high' })

    assert_equal 'retrying', result[:status]
    assert_equal 'unsupported_channel', result[:reason]
    assert_equal ['queued:fax', 'failed:fax', 'retry_scheduled:fax'], result[:audit]
  end
end
