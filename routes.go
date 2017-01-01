package main

import (
	"github.com/gin-gonic/gin"
	"github.com/edouardhue/wmfr-adhesions/iraiser"
	"github.com/edouardhue/wmfr-adhesions/memberships"
)

type Routes struct {
	adhesions memberships.Memberships
}

func MemberRoute(config *memberships.Config) gin.HandlerFunc {
	adhesions := memberships.NewMemberships(config)
	return func(c *gin.Context) {
		var donation iraiser.Donation
		if err := c.Bind(&donation); err != nil {
			c.AbortWithError(400, err)
		} else {
			if err := adhesions.RecordMembership(donation) ; err != nil {
				switch err.(type) {
				case *memberships.NoSuchContactError:
					c.AbortWithError(404, err)
				case *memberships.NoCommonMembershipError:
					c.AbortWithError(404, err)
				default:
					c.AbortWithError(500, err)
				}
			} else {
				c.Status(201)
			}
		}
	}
}

