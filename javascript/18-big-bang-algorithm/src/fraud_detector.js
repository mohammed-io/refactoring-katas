class FraudDetector {
  constructor() {}

  detect(tx) {
    let s = 0;
    let v = 0;
    let m = 0;
    let h = 0;
    let d = new Date(tx.timestamp).getHours();

    for (let i = 0; i < tx.history.length; i++) {
      if (tx.history[i].amount > tx.amount * 2) {
        v++;
      }
      let diff = Math.abs(tx.timestamp - tx.history[i].timestamp);
      if (diff < 3600000) {
        s++;
      }
    }

    if (tx.amount > 500 && d >= 0 && d < 6) {
      m += 30;
    }
    if (tx.amount > 1000) {
      m += 20;
    }
    if (tx.merchant === 'gambling' || tx.merchant === 'crypto') {
      m += 25;
    }
    if (tx.country !== tx.cardCountry) {
      m += 15;
    }
    if (s > 3) {
      m += 20;
    }
    if (v > 2) {
      m += 15;
    }

    if (m < 20) {
      h = 1;
    } else if (m < 40) {
      h = 2;
    } else if (m < 60) {
      h = 3;
    } else if (m < 80) {
      h = 4;
    } else {
      h = 5;
    }

    let r = 'low';
    if (h === 2) {
      r = 'medium';
    } else if (h === 3) {
      r = 'elevated';
    } else if (h === 4) {
      r = 'high';
    } else if (h === 5) {
      r = 'critical';
    }

    return { score: m, level: h, rating: r };
  }
}

export { FraudDetector };
