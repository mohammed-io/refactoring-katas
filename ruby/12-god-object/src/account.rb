# frozen_string_literal: true

class Account
  attr_accessor :email, :password, :profile, :payment_methods, :notifications, :audit_log, :subscription

  def initialize(email, password)
    @email = email
    @password = password
    @profile = { name: '', bio: '' }
    @payment_methods = []
    @notifications = { email: true, sms: false }
    @audit_log = []
    @subscription = 'basic'
  end

  def login(pw)
    if pw == @password
      @audit_log << "login:#{Time.now.to_i}"
      true
    else
      false
    end
  end

  def logout
    @audit_log << "logout:#{Time.now.to_i}"
    true
  end

  def update_profile(name, bio)
    @profile[:name] = name
    @profile[:bio] = bio
    @profile
  end

  def change_password(old, newp)
    return false unless old == @password

    @password = newp
    true
  end

  def add_payment_method(card)
    @payment_methods << card
    @payment_methods.length
  end

  def remove_payment_method(index)
    @payment_methods.delete_at(index)
    @payment_methods
  end

  def set_notification_preference(type, value)
    @notifications[type] = value
    @notifications
  end

  def export_data
    { email: @email, profile: @profile, payment_methods: @payment_methods, notifications: @notifications,
      audit_log: @audit_log, subscription: @subscription }
  end

  def log_access(action)
    @audit_log << "#{action}:#{Time.now.to_i}"
    @audit_log.length
  end

  def check_subscription
    @subscription
  end

  def upgrade_subscription(plan)
    @subscription = plan
    @subscription
  end
end
