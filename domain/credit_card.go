package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CreditCard struct {
	ID              string
	Name            string
	Number          string
	ExpiretionMonth int32
	ExpiretionYear  int32
	CVV             int32
	Balance         float64
	Limit           float64
	CreatedAt       time.Time
}

func NewCreditCard() *CreditCard {
	card := &CreditCard{}
	card.ID = uuid.NewV4().String()
	card.CreatedAt = time.Now()
	return card
}
