package main

type iRaiserMember struct {
	Mail string `json:"email" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	City string `json:"city" binding:"required"`
	PostCode string `json:"postCode" binding:"required"`
	Amount uint `json:"amount" binding:"required"`
}