using Xunit;

public class IssueUpdaterTest
{
    [Fact]
    public void ClosesIssue()
    {
        var updater = new IssueUpdater();
        var result = updater.update_issue(42, "close");
        Assert.Contains("status=closed", result);
    }

    [Fact]
    public void OpensIssue()
    {
        var updater = new IssueUpdater();
        var result = updater.update_issue(42, "open");
        Assert.Contains("status=open", result);
    }

    [Fact]
    public void SetsInProgress()
    {
        var updater = new IssueUpdater();
        var result = updater.update_issue(42, "progress");
        Assert.Contains("status=in_progress", result);
    }

    [Fact]
    public void SetsPriority1()
    {
        var updater = new IssueUpdater();
        var result = updater.update_issue(42, "close:1");
        Assert.Contains("priority=1", result);
    }

    [Fact]
    public void SetsPriority2()
    {
        var updater = new IssueUpdater();
        var result = updater.update_issue(42, "open:2");
        Assert.Contains("priority=2", result);
    }

    [Fact]
    public void DefaultsToPriority3()
    {
        var updater = new IssueUpdater();
        var result = updater.update_issue(42, "progress");
        Assert.Contains("priority=3", result);
    }

    [Fact]
    public void IncludesIssueId()
    {
        var updater = new IssueUpdater();
        var result = updater.update_issue(99, "close");
        Assert.Contains("Issue 99", result);
    }

    [Fact]
    public void IgnoresInvalidPriority()
    {
        var updater = new IssueUpdater();
        var result = updater.update_issue(42, "close:99");
        Assert.Contains("priority=3", result);
    }
}
