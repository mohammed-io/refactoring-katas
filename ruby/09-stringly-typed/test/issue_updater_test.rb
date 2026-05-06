# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/issue_updater'

class IssueUpdaterTest < Minitest::Test
  def setup
    @updater = IssueUpdater.new
  end

  def test_closes_issue
    result = @updater.update_issue(42, 'close')
    assert_includes result, 'status=closed'
  end

  def test_opens_issue
    result = @updater.update_issue(42, 'open')
    assert_includes result, 'status=open'
  end

  def test_sets_in_progress
    result = @updater.update_issue(42, 'progress')
    assert_includes result, 'status=in_progress'
  end

  def test_sets_priority_1
    result = @updater.update_issue(42, 'close:1')
    assert_includes result, 'priority=1'
  end

  def test_sets_priority_2
    result = @updater.update_issue(42, 'open:2')
    assert_includes result, 'priority=2'
  end

  def test_defaults_to_priority_3
    result = @updater.update_issue(42, 'progress')
    assert_includes result, 'priority=3'
  end

  def test_includes_issue_id
    result = @updater.update_issue(99, 'close')
    assert_includes result, 'Issue 99'
  end

  def test_ignores_invalid_priority
    result = @updater.update_issue(42, 'close:99')
    assert_includes result, 'priority=3'
  end
end
