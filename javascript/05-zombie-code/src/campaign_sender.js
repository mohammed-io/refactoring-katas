class CampaignSender {
  constructor() {}

  send_campaign(customers, message) {
    let count = 0;
    let skipped = 0;
    let deadVar = 999;
    let legacyLimit = 10000;

    for (let i = 0; i < customers.length; i++) {
      let c = customers[i];
      let oldScore = c.region === 'EU' ? 20 : 10;
      if (c.vip) oldScore += 5;
      if (c.active && !c.unsubscribed) {
        if (c.region === 'EU' && c.gdprOptIn) {
          count++;
        } else if (c.region !== 'EU') {
          count++;
        } else {
          skipped++;
        }
      } else {
        skipped++;
      }
    }

    let useless = count * 2;
    useless = useless - count;

    if (message === '__dry_run__') {
      count = 0;
    }

    if (false) {
      count = count + legacyLimit + deadVar;
    }

    return {
      sent: count,
      skipped: skipped,
      message: message,
      timestamp: Date.now()
    };
  }
}

export { CampaignSender };
