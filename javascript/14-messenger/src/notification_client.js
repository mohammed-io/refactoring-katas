class NotificationClient {
  send(payload) {
    return { status: 'sent', payload: payload };
  }
}

export { NotificationClient };
