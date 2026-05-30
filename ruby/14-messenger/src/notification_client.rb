# frozen_string_literal: true

class NotificationBackend
  def send(payload)
    return { status: 'failed', reason: 'unsupported_channel', payload: payload } if payload[:channel] == 'fax'

    { status: 'sent', delivery_id: "#{payload[:channel]}-#{payload[:recipient]}", payload: payload }
  end
end

class NotificationGateway
  def initialize
    @client = NotificationBackend.new
  end

  def dispatch(payload)
    @client.send(payload)
  end
end

class NotificationAudit
  def record(event, payload)
    "#{event}:#{payload[:channel]}"
  end
end

class NotificationClient
  def initialize
    @gateway = NotificationGateway.new
    @audit = NotificationAudit.new
  end

  def send(payload)
    return { status: 'rejected', reason: 'missing_message', payload: payload } if payload[:message].to_s.empty?

    normalized = {
      recipient: payload[:recipient] || 'unknown',
      message: payload[:message],
      channel: payload[:channel] || 'email',
      priority: payload[:priority] || 'normal'
    }
    result = @gateway.dispatch(normalized)
    result[:audit] = [@audit.record('queued', normalized), @audit.record(result[:status], normalized)]
    if result[:status] == 'failed' && normalized[:priority] == 'high'
      result[:status] = 'retrying'
      result[:audit] << @audit.record('retry_scheduled', normalized)
    end
    result
  end
end
