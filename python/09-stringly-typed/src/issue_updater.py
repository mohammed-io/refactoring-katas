class IssueUpdater:
    def __init__(self):
        pass

    def update_issue(self, id, command):
        parts = command.split(":")
        action = parts[0]
        value = parts[1] if len(parts) > 1 else ""
        status = "open"
        priority = "3"
        if action == "close": status = "closed"
        elif action == "open": status = "open"
        elif action == "progress": status = "in_progress"
        if value in ["1","2","3"]: priority = value
        return f"Issue {id} updated to status={status} priority={priority}"
