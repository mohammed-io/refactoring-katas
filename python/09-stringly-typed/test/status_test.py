from src.issue_updater import IssueUpdater

def test_close_issue():
    updater = IssueUpdater()
    result = updater.update_issue(42, "close")
    assert "status=closed" in result

def test_open_issue():
    updater = IssueUpdater()
    result = updater.update_issue(42, "open")
    assert "status=open" in result

def test_in_progress():
    updater = IssueUpdater()
    result = updater.update_issue(42, "progress")
    assert "status=in_progress" in result

def test_priority_1():
    updater = IssueUpdater()
    result = updater.update_issue(42, "close:1")
    assert "priority=1" in result

def test_priority_2():
    updater = IssueUpdater()
    result = updater.update_issue(42, "open:2")
    assert "priority=2" in result

def test_default_priority_3():
    updater = IssueUpdater()
    result = updater.update_issue(42, "progress")
    assert "priority=3" in result

def test_includes_issue_id():
    updater = IssueUpdater()
    result = updater.update_issue(99, "close")
    assert "Issue 99" in result

def test_invalid_priority():
    updater = IssueUpdater()
    result = updater.update_issue(42, "close:99")
    assert "priority=3" in result
