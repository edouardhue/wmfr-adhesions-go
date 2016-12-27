package iraiser

import (
	"math/big"
	"time"
)

type Member struct {
	Mail string `json:"email" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	City string `json:"city" binding:"required"`
	Country string `json:"country" binding:"required"`
	Amount big.Float `json:"amount" binding:"required"`
	Date time.Time `json:"date" binding:"required"`
}
