package main

import (
	"github.com/gin-gonic/gin"
	"github.com/edouardhue/wmfr-adhesions/iraiser"
)

func addMember(c *gin.Context) {
	var member iraiser.Member
	if err := c.Bind(&member); err != nil {
		c.AbortWithError(500, err)
	} else {
		if err := recordMembership(member) ; err != nil {
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