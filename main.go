package main

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wikimedia-france/wmfr-adhesions/internal"
	"github.com/wikimedia-france/wmfr-adhesions/iraiser"
	"io/ioutil"
	"os"
)

func main() {
	logFile, err := os.OpenFile(internal.Config.Log, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	r := gin.New()
	r.Use(gin.LoggerWithWriter(logFile), gin.Recovery())

	m := r.Group("/memberships")
	{
		m.GET("/status", func(c *gin.Context) {
			c.Status(200)
		})

		m.POST("/debug", func(c *gin.Context) {
			body, _ := ioutil.ReadAll(c.Request.Body)
			login := c.Request.Header.Get("secureLogin")
			timestamp := c.Request.Header.Get("secureTimestamp")
			token := c.Request.Header.Get("secureToken")
			fmt.Fprintf(logFile, "%s:%s:%s\n%s\n", login, timestamp, token, string(body))
			c.Status(202)
		})

		authorized := m.Group("/1")
		authorized.Use(iRaiserAuthentication())
		{
			authorized.POST("/members", MemberRoute())
		}
	}

	r.Run(":8000")
}

func iRaiserAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		login := c.Request.Header.Get("secureLogin")
		timestamp := c.Request.Header.Get("secureTimestamp")
		token := c.Request.Header.Get("secureToken")
		tokenBytes, err := hex.DecodeString(token)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		var secureHeader = iraiser.SecureHeader{
			Login:     login,
			Timestamp: timestamp,
			Token:     tokenBytes,
		}
		if !iraiser.Verify(&secureHeader) {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}
