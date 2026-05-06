import time

class Account:
    def __init__(self, email, password):
        self.email = email
        self.password = password
        self.profile = {"name": "", "bio": ""}
        self.payment_methods = []
        self.notifications = {"email": True, "sms": False}
        self.audit_log = []
        self.subscription = "basic"

    def login(self, pw):
        if pw == self.password:
            self.audit_log.append("login:" + str(int(time.time() * 1000)))
            return True
        return False

    def logout(self):
        self.audit_log.append("logout:" + str(int(time.time() * 1000)))
        return True

    def update_profile(self, name, bio):
        self.profile["name"] = name
        self.profile["bio"] = bio
        return self.profile

    def change_password(self, old, new):
        if old == self.password:
            self.password = new
            return True
        return False

    def add_payment_method(self, card):
        self.payment_methods.append(card)
        return len(self.payment_methods)

    def remove_payment_method(self, index):
        self.payment_methods.pop(index)
        return self.payment_methods

    def set_notification_preference(self, typ, value):
        self.notifications[typ] = value
        return self.notifications

    def export_data(self):
        return {
            "email": self.email,
            "profile": self.profile,
            "paymentMethods": self.payment_methods,
            "notifications": self.notifications,
            "auditLog": self.audit_log,
            "subscription": self.subscription
        }

    def log_access(self, action):
        self.audit_log.append(action + ":" + str(int(time.time() * 1000)))
        return len(self.audit_log)

    def check_subscription(self):
        return self.subscription

    def upgrade_subscription(self, plan):
        self.subscription = plan
        return self.subscription
