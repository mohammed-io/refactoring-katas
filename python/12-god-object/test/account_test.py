from src.account import Account

def test_login_correct_password():
    acc = Account("a@b.com", "secret")
    assert acc.login("secret") is True

def test_reject_wrong_password():
    acc = Account("a@b.com", "secret")
    assert acc.login("wrong") is False

def test_logout():
    acc = Account("a@b.com", "secret")
    assert acc.logout() is True

def test_update_profile():
    acc = Account("a@b.com", "secret")
    result = acc.update_profile("Alice", "Developer")
    assert result["name"] == "Alice"
    assert result["bio"] == "Developer"

def test_change_password():
    acc = Account("a@b.com", "secret")
    assert acc.change_password("secret", "new") is True
    assert acc.login("new") is True

def test_reject_bad_old_password():
    acc = Account("a@b.com", "secret")
    assert acc.change_password("wrong", "new") is False

def test_add_payment_method():
    acc = Account("a@b.com", "secret")
    count = acc.add_payment_method("Visa-1234")
    assert count == 1

def test_remove_payment_method():
    acc = Account("a@b.com", "secret")
    acc.add_payment_method("Visa-1234")
    result = acc.remove_payment_method(0)
    assert result == []

def test_set_notification_preference():
    acc = Account("a@b.com", "secret")
    result = acc.set_notification_preference("sms", True)
    assert result["sms"] is True

def test_export_all_data():
    acc = Account("a@b.com", "secret")
    result = acc.export_data()
    assert result["email"] == "a@b.com"
    assert "auditLog" in result

def test_log_access():
    acc = Account("a@b.com", "secret")
    count = acc.log_access("view")
    assert count == 1

def test_check_subscription():
    acc = Account("a@b.com", "secret")
    assert acc.check_subscription() == "basic"

def test_upgrade_subscription():
    acc = Account("a@b.com", "secret")
    result = acc.upgrade_subscription("pro")
    assert result == "pro"
