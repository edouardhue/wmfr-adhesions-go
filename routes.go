package main

import (
	"github.com/gin-gonic/gin"
)

func addMember(c *gin.Context) {
	var member iRaiserMember
	if c.Bind(&member) == nil {
		if found, err := lookupMember(member) ; err != nil {
			c.Error(err)
		} else if found == 1 {
			c.Status(200)
		} else {
			c.Status(404)
		}
	}
}