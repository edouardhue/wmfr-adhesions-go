package main

import (
	"github.com/gin-gonic/gin"
)

func addMember(c *gin.Context) {
	var member iRaiserMember
	if c.Bind(&member) == nil {
		if err := updateOrCreateMembership(member) ; err != nil {
			switch err.(type) {
				case *NoSuchContactError:
					c.AbortWithError(404, err)
				case *NoCommonMembershipError:
					c.AbortWithError(404, err)
				default:
					c.AbortWithError(500, err)
			}
		} else {
			c.Status(201)
		}
	}
}