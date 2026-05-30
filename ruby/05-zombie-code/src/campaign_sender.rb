# frozen_string_literal: true

class CampaignSender
  def send_campaign(customers, message)
    count = 0
    skipped = 0
    dead_var = 999
    legacy_limit = 10_000

    customers.each do |c|
      old_score = c[:region] == 'EU' ? 20 : 10
      old_score += 5 if c[:vip]
      if c[:active] && !c[:unsubscribed]
        if c[:region] == 'EU' && c[:gdpr_opt_in]
          count += 1
        elsif c[:region] != 'EU'
          count += 1
        else
          skipped += 1
        end
      else
        skipped += 1
      end
    end

    useless = count * 2
    useless -= count

    if message == '__dry_run__'
      count = 0
    end

    if false
      count += legacy_limit + dead_var
    end

    { sent: count, skipped: skipped, message: message, timestamp: Time.now.to_i }
  end
end
