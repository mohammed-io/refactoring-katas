using Xunit;

public class AccountTest
{
    [Fact]
    public void LogsInWithCorrectPassword()
    {
        var acc = new Account("a@b.com", "secret");
        Assert.True(acc.login("secret"));
    }

    [Fact]
    public void RejectsWrongPassword()
    {
        var acc = new Account("a@b.com", "secret");
        Assert.False(acc.login("wrong"));
    }

    [Fact]
    public void LogsOut()
    {
        var acc = new Account("a@b.com", "secret");
        Assert.True(acc.logout());
    }

    [Fact]
    public void UpdatesProfile()
    {
        var acc = new Account("a@b.com", "secret");
        var result = acc.update_profile("Alice", "Developer");
        Assert.Equal("Alice", result["name"]);
        Assert.Equal("Developer", result["bio"]);
    }

    [Fact]
    public void ChangesPassword()
    {
        var acc = new Account("a@b.com", "secret");
        Assert.True(acc.change_password("secret", "new"));
        Assert.True(acc.login("new"));
    }

    [Fact]
    public void RejectsBadOldPassword()
    {
        var acc = new Account("a@b.com", "secret");
        Assert.False(acc.change_password("wrong", "new"));
    }

    [Fact]
    public void AddsPaymentMethod()
    {
        var acc = new Account("a@b.com", "secret");
        var count = acc.add_payment_method("Visa-1234");
        Assert.Equal(1, count);
    }

    [Fact]
    public void RemovesPaymentMethod()
    {
        var acc = new Account("a@b.com", "secret");
        acc.add_payment_method("Visa-1234");
        var result = acc.remove_payment_method(0);
        Assert.Empty(result);
    }

    [Fact]
    public void SetsNotificationPreference()
    {
        var acc = new Account("a@b.com", "secret");
        var result = acc.set_notification_preference("sms", true);
        Assert.True(result["sms"]);
    }

    [Fact]
    public void ExportsAllData()
    {
        var acc = new Account("a@b.com", "secret");
        var result = acc.export_data();
        Assert.Equal("a@b.com", result.email);
        Assert.NotNull(result.auditLog);
    }

    [Fact]
    public void LogsAccess()
    {
        var acc = new Account("a@b.com", "secret");
        var count = acc.log_access("view");
        Assert.Equal(1, count);
    }

    [Fact]
    public void ChecksSubscription()
    {
        var acc = new Account("a@b.com", "secret");
        Assert.Equal("basic", acc.check_subscription());
    }

    [Fact]
    public void UpgradesSubscription()
    {
        var acc = new Account("a@b.com", "secret");
        var result = acc.upgrade_subscription("pro");
        Assert.Equal("pro", result);
    }
}
