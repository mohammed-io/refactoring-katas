package kata

import (
	"strings"
	"testing"
)

func TestUpdateIssueClosesIssue(t *testing.T) {
	iu := NewIssueUpdater()
	result := iu.update_issue(42, "close")
	if !strings.Contains(result, "status=closed") {
		t.Errorf("expected result to contain 'status=closed', got %q", result)
	}
}

func TestUpdateIssueOpensIssue(t *testing.T) {
	iu := NewIssueUpdater()
	result := iu.update_issue(42, "open")
	if !strings.Contains(result, "status=open") {
		t.Errorf("expected result to contain 'status=open', got %q", result)
	}
}

func TestUpdateIssueSetsInProgress(t *testing.T) {
	iu := NewIssueUpdater()
	result := iu.update_issue(42, "progress")
	if !strings.Contains(result, "status=in_progress") {
		t.Errorf("expected result to contain 'status=in_progress', got %q", result)
	}
}

func TestUpdateIssueSetsPriority1(t *testing.T) {
	iu := NewIssueUpdater()
	result := iu.update_issue(42, "close:1")
	if !strings.Contains(result, "priority=1") {
		t.Errorf("expected result to contain 'priority=1', got %q", result)
	}
}

func TestUpdateIssueSetsPriority2(t *testing.T) {
	iu := NewIssueUpdater()
	result := iu.update_issue(42, "open:2")
	if !strings.Contains(result, "priority=2") {
		t.Errorf("expected result to contain 'priority=2', got %q", result)
	}
}

func TestUpdateIssueDefaultsToPriority3(t *testing.T) {
	iu := NewIssueUpdater()
	result := iu.update_issue(42, "progress")
	if !strings.Contains(result, "priority=3") {
		t.Errorf("expected result to contain 'priority=3', got %q", result)
	}
}

func TestUpdateIssueIncludesIssueId(t *testing.T) {
	iu := NewIssueUpdater()
	result := iu.update_issue(99, "close")
	if !strings.Contains(result, "Issue 99") {
		t.Errorf("expected result to contain 'Issue 99', got %q", result)
	}
}

func TestUpdateIssueIgnoresInvalidPriority(t *testing.T) {
	iu := NewIssueUpdater()
	result := iu.update_issue(42, "close:99")
	if !strings.Contains(result, "priority=3") {
		t.Errorf("expected result to contain 'priority=3', got %q", result)
	}
}
