# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/account'

class AccountTest < Minitest::Test
  def setup
    @account = Account.new('a@b.com', 'secret')
  end

  def test_logs_in_with_correct_password
    assert_equal true, @account.login('secret')
  end

  def test_rejects_wrong_password
    assert_equal false, @account.login('wrong')
  end

  def test_logs_out
    assert_equal true, @account.logout
  end

  def test_updates_profile
    result = @account.update_profile('Alice', 'Developer')
    assert_equal 'Alice', result[:name]
    assert_equal 'Developer', result[:bio]
  end

  def test_changes_password
    assert_equal true, @account.change_password('secret', 'new')
    assert_equal true, @account.login('new')
  end

  def test_rejects_bad_old_password
    assert_equal false, @account.change_password('wrong', 'new')
  end

  def test_adds_payment_method
    count = @account.add_payment_method('Visa-1234')
    assert_equal 1, count
  end

  def test_removes_payment_method
    @account.add_payment_method('Visa-1234')
    result = @account.remove_payment_method(0)
    assert_equal [], result
  end

  def test_sets_notification_preference
    result = @account.set_notification_preference(:sms, true)
    assert_equal true, result[:sms]
  end

  def test_exports_all_data
    result = @account.export_data
    assert_equal 'a@b.com', result[:email]
    assert result[:audit_log]
  end

  def test_logs_access
    count = @account.log_access('view')
    assert_equal 1, count
  end

  def test_checks_subscription
    assert_equal 'basic', @account.check_subscription
  end

  def test_upgrades_subscription
    result = @account.upgrade_subscription('pro')
    assert_equal 'pro', result
  end
end
