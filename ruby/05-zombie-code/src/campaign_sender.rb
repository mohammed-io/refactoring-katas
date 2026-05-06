# frozen_string_literal: true

class CampaignSender
  def send_campaign(customers, message)
    count = 0
    dead_var = 999

    customers.each do |c|
      if c[:active]
        if c[:region] == 'EU' && c[:gdpr_opt_in]
          count += 1
        elsif c[:region] != 'EU'
          count += 1
        end
      end
    end

    useless = count * 2
    useless -= count

    if false
      count += 100
    end

    { sent: count, message: message, timestamp: Time.now.to_i }
  end
end
