package domain

import "time"

type CreditCard struct {
	ID              string
	Name            string
	Number          string
	ExpiretionMonth int32
	ExpiretionYear  int32
	CVV             int32
	Balance         float64
	limit           float64
	CreatedAt       time.Time
}
