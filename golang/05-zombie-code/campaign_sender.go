package kata

import "time"

type CampaignCustomer struct {
	Active bool
	Region string
	GDPR   bool
}
type CampaignResult struct {
	Sent      int
	Message   string
	Timestamp int64
}

type CampaignSender struct{}

func NewCampaignSender() *CampaignSender {
	return &CampaignSender{}
}

func (cs *CampaignSender) send_campaign(customers []CampaignCustomer, message string) CampaignResult {
	count := 0
	deadVar := 999
	_ = deadVar
	for _, c := range customers {
		if c.Active {
			if c.Region == "EU" && c.GDPR {
				count++
			} else if c.Region != "EU" {
				count++
			}
		}
	}
	useless := count * 2
	useless = useless - count
	_ = useless
	if false {
		count += 100
	}
	return CampaignResult{count, message, time.Now().UnixMilli()}
}
