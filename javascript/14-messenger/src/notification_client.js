class NotificationBackend {
  send(payload) {
    if (payload.channel === 'fax') {
      return { status: 'failed', reason: 'unsupported_channel', payload };
    }
    return { status: 'sent', delivery_id: `${payload.channel}-${payload.recipient}`, payload };
  }
}

class NotificationGateway {
  constructor() {
    this.client = new NotificationBackend();
  }

  dispatch(payload) {
    return this.client.send(payload);
  }
}

class NotificationAudit {
  record(event, payload) {
    return `${event}:${payload.channel}`;
  }
}

class NotificationClient {
  constructor() {
    this.gateway = new NotificationGateway();
    this.audit = new NotificationAudit();
  }

  send(payload) {
    if (!payload.message) {
      return { status: 'rejected', reason: 'missing_message', payload };
    }

    const normalized = {
      recipient: payload.recipient ?? 'unknown',
      message: payload.message,
      channel: payload.channel ?? 'email',
      priority: payload.priority ?? 'normal',
    };
    const result = this.gateway.dispatch(normalized);
    result.audit = [this.audit.record('queued', normalized), this.audit.record(result.status, normalized)];
    if (result.status === 'failed' && normalized.priority === 'high') {
      result.status = 'retrying';
      result.audit.push(this.audit.record('retry_scheduled', normalized));
    }
    return result;
  }
}

export { NotificationClient };
