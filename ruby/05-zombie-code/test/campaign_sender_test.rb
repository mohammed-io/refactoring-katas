# frozen_string_literal: true

require 'minitest/autorun'
require_relative '../src/campaign_sender'

class CampaignSenderTest < Minitest::Test
  def setup
    @sender = CampaignSender.new
  end

  def test_counts_eu_customers_with_opt_in
    result = @sender.send_campaign([{ active: true, region: 'EU', gdpr_opt_in: true }], 'Hi')
    assert_equal 1, result[:sent]
  end

  def test_skips_eu_customers_without_opt_in
    result = @sender.send_campaign([{ active: true, region: 'EU', gdpr_opt_in: false }], 'Hi')
    assert_equal 0, result[:sent]
  end

  def test_counts_non_eu_active_customers
    result = @sender.send_campaign([{ active: true, region: 'US' }], 'Hi')
    assert_equal 1, result[:sent]
  end

  def test_skips_inactive_customers
    result = @sender.send_campaign([{ active: false, region: 'US' }], 'Hi')
    assert_equal 0, result[:sent]
  end

  def test_handles_mixed_customers
    result = @sender.send_campaign([
      { active: true, region: 'EU', gdpr_opt_in: true },
      { active: true, region: 'EU', gdpr_opt_in: false },
      { active: true, region: 'US' },
      { active: false, region: 'US' }
    ], 'Hi')
    assert_equal 2, result[:sent]
  end

  def test_returns_message_in_result
    result = @sender.send_campaign([], 'Hello World')
    assert_equal 'Hello World', result[:message]
  end

  def test_returns_timestamp_in_result
    result = @sender.send_campaign([], 'Hi')
    assert_kind_of Integer, result[:timestamp]
  end
end
