package kata

import (
	"testing"
)

func TestAccountLogsInWithCorrectPassword(t *testing.T) {
	acc := NewAccount("a@b.com", "secret")
	if !acc.login("secret") {
		t.Error("expected login to succeed with correct password")
	}
}

func TestAccountRejectsWrongPassword(t *testing.T) {
	acc := NewAccount("a@b.com", "secret")
	if acc.login("wrong") {
		t.Error("expected login to fail with wrong password")
	}
}

func TestAccountLogsOut(t *testing.T) {
	acc := NewAccount("a@b.com", "secret")
	if !acc.logout() {
		t.Error("expected logout to succeed")
	}
}

func TestAccountUpdatesProfile(t *testing.T) {
	acc := NewAccount("a@b.com", "secret")
	result := acc.update_profile("Alice", "Developer")
	if result["name"] != "Alice" {
		t.Errorf("expected name 'Alice', got %q", result["name"])
	}
	if result["bio"] != "Developer" {
		t.Errorf("expected bio 'Developer', got %q", result["bio"])
	}
}

func TestAccountChangesPassword(t *testing.T) {
	acc := NewAccount("a@b.com", "secret")
	if !acc.change_password("secret", "new") {
		t.Error("expected password change to succeed")
	}
	if !acc.login("new") {
		t.Error("expected login to succeed with new password")
	}
}

func TestAccountRejectsBadOldPassword(t *testing.T) {
	acc := NewAccount("a@b.com", "secret")
	if acc.change_password("wrong", "new") {
		t.Error("expected password change to fail with wrong old password")
	}
}

func TestAccountAddsPaymentMethod(t *testing.T) {
	acc := NewAccount("a@b.com", "secret")
	count := acc.add_payment_method("Visa-1234")
	if count != 1 {
		t.Errorf("expected 1 payment method, got %d", count)
	}
}

func TestAccountRemovesPaymentMethod(t *testing.T) {
	acc := NewAccount("a@b.com", "secret")
	acc.add_payment_method("Visa-1234")
	result := acc.remove_payment_method(0)
	if len(result) != 0 {
		t.Errorf("expected 0 payment methods after removal, got %d", len(result))
	}
}

func TestAccountSetsNotificationPreference(t *testing.T) {
	acc := NewAccount("a@b.com", "secret")
	result := acc.set_notification_preference("sms", true)
	if !result["sms"] {
		t.Error("expected sms notification to be enabled")
	}
}

func TestAccountExportsAllData(t *testing.T) {
	acc := NewAccount("a@b.com", "secret")
	result := acc.export_data()
	if result["email"] != "a@b.com" {
		t.Errorf("expected email 'a@b.com', got %v", result["email"])
	}
}

func TestAccountLogsAccess(t *testing.T) {
	acc := NewAccount("a@b.com", "secret")
	length := acc.log_access("view_profile")
	if length != 1 {
		t.Errorf("expected audit log length 1, got %d", length)
	}
}

func TestAccountChecksSubscription(t *testing.T) {
	acc := NewAccount("a@b.com", "secret")
	result := acc.check_subscription()
	if result != "basic" {
		t.Errorf("expected subscription 'basic', got %q", result)
	}
}

func TestAccountUpgradesSubscription(t *testing.T) {
	acc := NewAccount("a@b.com", "secret")
	result := acc.upgrade_subscription("pro")
	if result != "pro" {
		t.Errorf("expected subscription 'pro', got %q", result)
	}
}
