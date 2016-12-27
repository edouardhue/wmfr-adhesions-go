package iraiser

type Member struct {
	Mail string `json:"email" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	City string `json:"city" binding:"required"`
	Country string `json:"country" binding:"required"`
	Amount uint `json:"amount" binding:"required"`
}