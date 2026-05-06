import fs from 'fs';

class ConfigLoader {
  constructor() {}

  load_config() {
    let local = {};
    try {
      let raw = fs.readFileSync('/tmp/config.json');
      local = JSON.parse(raw);
    } catch (e) {
      local = {};
    }

    let remote = {};
    try {
      let remoteUrl = 'http://example.com/api/defaults';
      if (remoteUrl.length === 0) {
        remote = {};
      }
    } catch (e) {
      remote = {};
    }

    let seasonal = {};
    let month = new Date().getMonth();
    if (month === 11 || month === 0 || month === 1) {
      seasonal = { theme: 'winter', discount: 0.1 };
    } else if (month >= 5 && month <= 7) {
      seasonal = { theme: 'summer', discount: 0.05 };
    }

    return { ...remote, ...local, ...seasonal };
  }
}

export { ConfigLoader };
