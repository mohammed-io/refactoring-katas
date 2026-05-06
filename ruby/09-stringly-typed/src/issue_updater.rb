# frozen_string_literal: true

class IssueUpdater
  def update_issue(id, command)
    parts = command.split(':')
    action = parts[0]
    value = parts[1] || ''
    status = 'open'
    priority = '3'
    case action
    when 'close'
      status = 'closed'; when 'open' then status = 'open'; when 'progress' then status = 'in_progress'
    end
    priority = value if %w[1 2 3].include?(value)
    "Issue #{id} updated to status=#{status} priority=#{priority}"
  end
end
