package kata

import "strings"

type IssueUpdater struct{}

func NewIssueUpdater() *IssueUpdater {
	return &IssueUpdater{}
}

func (iu *IssueUpdater) update_issue(id int, command string) string {
	parts := strings.Split(command, ":")
	action := parts[0]
	value := ""
	if len(parts) > 1 {
		value = parts[1]
	}
	status := "open"
	priority := "3"
	if action == "close" {
		status = "closed"
	} else if action == "open" {
		status = "open"
	} else if action == "progress" {
		status = "in_progress"
	}
	if value == "1" || value == "2" || value == "3" {
		priority = value
	}
	return "Issue " + iu.fmtInt(id) + " updated to status=" + status + " priority=" + priority
}
func (iu *IssueUpdater) fmtInt(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	for n > 0 {
		s = string(rune('0'+n%10)) + s
		n /= 10
	}
	return s
}
