package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wikimedia-france/wmfr-adhesions/iraiser"
	"encoding/hex"
	"os"
	"gopkg.in/yaml.v2"
	"github.com/wikimedia-france/wmfr-adhesions/memberships"
)

func main() {
	config, err := readConfigurationFile()
	if err != nil {
		panic(err)
	}
	runServer(config)
}

func runServer(config *memberships.Config) {
	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.Status(200)
	})

	authorized := r.Group("/1")
	authorized.Use(iRaiserAuthentication(config))
	{
		authorized.POST("/members", MemberRoute(config))
	}

	r.Run(":8000")
}

func iRaiserAuthentication(config *memberships.Config) gin.HandlerFunc {
	iRaiser := iraiser.NewIRaiser(&config.IRaiser)
	return func(c *gin.Context) {
		login := c.Request.Header.Get("secureLogin")
		timestamp :=  c.Request.Header.Get("secureTimestamp")
		token := c.Request.Header.Get("secureToken")
		tokenBytes, err := hex.DecodeString(token)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		var secureHeader = iraiser.SecureHeader{
			Login: login,
			Timestamp: timestamp,
			Token: tokenBytes,
		}
		if !iRaiser.Verify(&secureHeader) {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}

func readConfigurationFile() (config *memberships.Config, _ error) {
	var location, exists = os.LookupEnv("CONFIG_LOCATION")
	if !exists {
		location = "./adhesions.yaml"
	}
	fileinfo, err := os.Stat(location)
	if err != nil {
		return nil, err
	}
	filesize := fileinfo.Size()
	fp, err := os.Open(location)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	buf := make([]byte, filesize)
	fp.Read(buf)
	if err = yaml.Unmarshal(buf, &config); err != nil {
		return nil, err
	}
	return config, nil
}