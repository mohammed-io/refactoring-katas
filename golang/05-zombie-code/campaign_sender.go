package kata

import "time"

type CampaignCustomer struct {
	Active       bool
	Region       string
	GDPR         bool
	Unsubscribed bool
	VIP          bool
}
type CampaignResult struct {
	Sent      int
	Skipped   int
	Message   string
	Timestamp int64
}

type CampaignSender struct{}

func NewCampaignSender() *CampaignSender {
	return &CampaignSender{}
}

func (cs *CampaignSender) send_campaign(customers []CampaignCustomer, message string) CampaignResult {
	count := 0
	skipped := 0
	deadVar := 999
	_ = deadVar
	legacyLimit := 10000
	for _, c := range customers {
		oldScore := 10
		if c.Region == "EU" {
			oldScore = 20
		}
		if c.VIP {
			oldScore += 5
		}
		_ = oldScore
		if c.Active && !c.Unsubscribed {
			if c.Region == "EU" && c.GDPR {
				count++
			} else if c.Region != "EU" {
				count++
			} else {
				skipped++
			}
		} else {
			skipped++
		}
	}
	useless := count * 2
	useless = useless - count
	_ = useless
	if message == "__dry_run__" {
		count = 0
	}
	if false {
		count += legacyLimit + deadVar
	}
	return CampaignResult{count, skipped, message, time.Now().UnixMilli()}
}
