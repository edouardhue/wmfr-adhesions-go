package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wikimedia-france/wmfr-adhesions/iraiser"
	"github.com/wikimedia-france/wmfr-adhesions/memberships"
)

func MemberRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var donation iraiser.Donation
		err := c.Bind(&donation)
		if err != nil {
			c.AbortWithError(400, err)
		}
		err = memberships.RecordMembership(&donation)
		if err != nil {
			c.AbortWithError(500, err)
		}
		c.Status(201)
	}
}
