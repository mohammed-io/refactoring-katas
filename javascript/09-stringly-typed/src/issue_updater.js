class IssueUpdater {
  constructor() {}

  update_issue(id, command) {
    let parts = command.split(':');
    let action = parts[0];
    let value = parts[1] || '';

    let status = 'open';
    let priority = '3';

    if (action === 'close') {
      status = 'closed';
    } else if (action === 'open') {
      status = 'open';
    } else if (action === 'progress') {
      status = 'in_progress';
    }

    if (value === '1' || value === '2' || value === '3') {
      priority = value;
    }

    let result = 'Issue ' + id + ' updated to status=' + status + ' priority=' + priority;
    return result;
  }
}

export { IssueUpdater };
