public class IssueUpdater
{
    public string update_issue(int id, string command)
    {
        var parts = command.Split(':');
        var action = parts[0];
        var value = parts.Length > 1 ? parts[1] : "";
        var status = "open";
        var priority = "3";

        if (action == "close")
        {
            status = "closed";
        }
        else if (action == "open")
        {
            status = "open";
        }
        else if (action == "progress")
        {
            status = "in_progress";
        }

        if (value == "1" || value == "2" || value == "3")
        {
            priority = value;
        }

        return $"Issue {id} updated to status={status} priority={priority}";
    }
}
