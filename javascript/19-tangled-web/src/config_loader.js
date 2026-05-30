import fs from 'fs';

class ConfigLoader {
  constructor() {}

  load_config() {
    const defaults = { retries: 3, theme: 'standard', discount: 0 };
    let local = {};
    try {
      const path = process.env.APP_CONFIG_PATH || '/tmp/config.json';
      let raw = fs.readFileSync(path);
      local = JSON.parse(raw);
    } catch (e) {
      local = {};
    }

    const envConfig = {};
    if (process.env.APP_THEME) {
      envConfig.theme = process.env.APP_THEME;
    }
    if (process.env.APP_RETRIES) {
      envConfig.retries = Number.parseInt(process.env.APP_RETRIES, 10);
    }

    let seasonal = {};
    let month = Number.parseInt(process.env.APP_CURRENT_MONTH || `${new Date().getMonth() + 1}`, 10);
    if (month === 12 || month === 1 || month === 2) {
      seasonal = { theme: 'winter', discount: 0.1 };
    } else if (month >= 6 && month <= 8) {
      seasonal = { theme: 'summer', discount: 0.05 };
    }

    return { ...defaults, ...local, ...envConfig, ...seasonal };
  }
}

export { ConfigLoader };
