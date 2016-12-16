package main

import "github.com/gin-gonic/gin"

func addMember(c *gin.Context) {
	var member iRaiserMember

	if c.Bind(&member) == nil {
		c.Status(201)
	}
}