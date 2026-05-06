# frozen_string_literal: true

class NotificationBackend
  def send(payload)
    { status: 'sent', payload: payload }
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

class NotificationClient
  def initialize
    @gateway = NotificationGateway.new
  end

  def send(payload)
    @gateway.dispatch(payload)
  end
end
