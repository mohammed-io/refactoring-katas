class CampaignSender {
  constructor() {}

  send_campaign(customers, message) {
    let count = 0;
    let deadVar = 999;

    for (let i = 0; i < customers.length; i++) {
      let c = customers[i];
      if (c.active) {
        if (c.region === 'EU' && c.gdprOptIn) {
          count++;
        } else if (c.region !== 'EU') {
          count++;
        }
      }
    }

    let useless = count * 2;
    useless = useless - count;

    if (false) {
      count = count + 100;
    }

    return {
      sent: count,
      message: message,
      timestamp: Date.now()
    };
  }
}

export { CampaignSender };
