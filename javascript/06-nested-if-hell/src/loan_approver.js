class LoanApprover {
  constructor() {}

  can_deliver(pkg) {
    if (pkg) {
      if (pkg.weight) {
        if (pkg.weight <= 50) {
          if (pkg.hazardous === false) {
            if (pkg.temperatureRequired) {
              if (pkg.temperatureRequired >= -20 && pkg.temperatureRequired <= 40) {
                if (pkg.weekend === false) {
                  if (pkg.remoteArea) {
                    if (pkg.remoteArea === true) {
                      if (pkg.weight <= 20) {
                        return { allowed: true, warning: 'Remote surcharge applies' };
                      } else {
                        return { allowed: false, warning: 'Too heavy for remote' };
                      }
                    } else {
                      return { allowed: true, warning: null };
                    }
                  } else {
                    return { allowed: true, warning: null };
                  }
                } else {
                  return { allowed: false, warning: 'No weekend delivery' };
                }
              } else {
                return { allowed: false, warning: 'Temperature out of range' };
              }
            } else {
              if (pkg.weekend === false) {
                if (pkg.remoteArea) {
                  if (pkg.remoteArea === true) {
                    if (pkg.weight <= 20) {
                      return { allowed: true, warning: 'Remote surcharge applies' };
                    } else {
                      return { allowed: false, warning: 'Too heavy for remote' };
                    }
                  } else {
                    return { allowed: true, warning: null };
                  }
                } else {
                  return { allowed: true, warning: null };
                }
              } else {
                return { allowed: false, warning: 'No weekend delivery' };
              }
            }
          } else {
            return { allowed: false, warning: 'Hazardous material' };
          }
        } else {
          return { allowed: false, warning: 'Weight exceeded' };
        }
      } else {
        return { allowed: false, warning: 'No weight specified' };
      }
    } else {
      return { allowed: false, warning: 'No package' };
    }
  }
}

export { LoanApprover };
