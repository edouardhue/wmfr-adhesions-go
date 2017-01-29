package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wikimedia-france/wmfr-adhesions/iraiser"
	"github.com/wikimedia-france/wmfr-adhesions/memberships"
)

type Routes struct {
	adhesions memberships.Memberships
}

func MemberRoute(config *memberships.Config) gin.HandlerFunc {
	adhesions := memberships.NewMemberships(config)
	return func(c *gin.Context) {
		var donation iraiser.Donation
		err := c.Bind(&donation)
		if err != nil {
			c.AbortWithError(400, err)
		}
		err = adhesions.RecordMembership(&donation)
		if err != nil {
			c.AbortWithError(500, err)
		}
		c.Status(201)
	}
}

