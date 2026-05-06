class Account {
  constructor(email, password) {
    this.email = email;
    this.password = password;
    this.profile = { name: '', bio: '' };
    this.paymentMethods = [];
    this.notifications = { email: true, sms: false };
    this.auditLog = [];
    this.subscription = 'basic';
  }

  login(pw) {
    if (pw === this.password) {
      this.auditLog.push('login:' + Date.now());
      return true;
    }
    return false;
  }

  logout() {
    this.auditLog.push('logout:' + Date.now());
    return true;
  }

  update_profile(name, bio) {
    this.profile.name = name;
    this.profile.bio = bio;
    return this.profile;
  }

  change_password(oldPw, newPw) {
    if (oldPw === this.password) {
      this.password = newPw;
      return true;
    }
    return false;
  }

  add_payment_method(card) {
    this.paymentMethods.push(card);
    return this.paymentMethods.length;
  }

  remove_payment_method(index) {
    this.paymentMethods.splice(index, 1);
    return this.paymentMethods;
  }

  set_notification_preference(type, value) {
    this.notifications[type] = value;
    return this.notifications;
  }

  export_data() {
    return {
      email: this.email,
      profile: this.profile,
      paymentMethods: this.paymentMethods,
      notifications: this.notifications,
      auditLog: this.auditLog,
      subscription: this.subscription
    };
  }

  log_access(action) {
    this.auditLog.push(action + ':' + Date.now());
    return this.auditLog.length;
  }

  check_subscription() {
    return this.subscription;
  }

  upgrade_subscription(newPlan) {
    this.subscription = newPlan;
    return this.subscription;
  }
}

export { Account };
