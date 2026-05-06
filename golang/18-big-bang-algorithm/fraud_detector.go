package kata

import "time"

type Tx struct {
	Amount      float64
	Timestamp   int64
	History     []Tx
	Merchant    string
	Country     string
	CardCountry string
}
type Risk struct {
	Score  int
	Level  int
	Rating string
}

type FraudDetector struct{}

func NewFraudDetector() *FraudDetector {
	return &FraudDetector{}
}

func (fd *FraudDetector) detect(tx Tx) Risk {
	s := 0
	v := 0
	m := 0
	d := time.UnixMilli(tx.Timestamp).UTC().Hour()
	for _, h := range tx.History {
		if h.Amount > tx.Amount*2 {
			v++
		}
		if fd.abs(tx.Timestamp-h.Timestamp) < 3600000 {
			s++
		}
	}
	if tx.Amount > 500 && d >= 0 && d < 6 {
		m += 30
	}
	if tx.Amount > 1000 {
		m += 20
	}
	if tx.Merchant == "gambling" || tx.Merchant == "crypto" {
		m += 25
	}
	if tx.Country != tx.CardCountry {
		m += 15
	}
	if s > 3 {
		m += 20
	}
	if v > 2 {
		m += 15
	}
	level := 1
	if m >= 80 {
		level = 5
	} else if m >= 60 {
		level = 4
	} else if m >= 40 {
		level = 3
	} else if m >= 20 {
		level = 2
	}
	rating := []string{"", "low", "medium", "elevated", "high", "critical"}[level]
	return Risk{m, level, rating}
}
func (fd *FraudDetector) abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}
