package main

import (
	"github.com/gin-gonic/gin"
	"github.com/edouardhue/wmfr-adhesions/iraiser"
	"encoding/hex"
	"os"
	"gopkg.in/yaml.v2"
	"github.com/edouardhue/wmfr-adhesions/memberships"
	"log"
)

func main() {
	if config, err := readConfigurationFile(); err != nil {
		panic(err)
	} else {
		log.Printf("%+v\n", config)

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
}

func iRaiserAuthentication(config *memberships.Config) gin.HandlerFunc {
	iRaiser := iraiser.NewIRaiser(&config.IRaiser)
	return func(c *gin.Context) {
		login := c.Request.Header.Get("secureLogin")
		timestamp :=  c.Request.Header.Get("secureTimestamp")
		token := c.Request.Header.Get("secureToken")
		if tokenBytes, err := hex.DecodeString(token); err != nil {
			c.AbortWithError(500, err)
		} else {
			var secureHeader = iraiser.SecureHeader{
				Login: login,
				Timestamp: timestamp,
				Token: tokenBytes,
			}
			if iRaiser.Verify(&secureHeader) {
				c.Next()
			} else {
				c.AbortWithStatus(401)
			}
		}
	}
}

func readConfigurationFile() (config *memberships.Config, _ error) {
	var location, exists = os.LookupEnv("CONFIG_LOCATION")
	if !exists {
		location = "./adhesions.yaml"
	}

	if fileinfo, err := os.Stat(location); err != nil {
		return nil, err
	} else {
		filesize := fileinfo.Size()

		if fp, err := os.Open(location); err != nil {
			return nil, err
		} else {
			defer fp.Close()

			buf := make([]byte, filesize)
			fp.Read(buf)
			if err = yaml.Unmarshal(buf, &config); err != nil {
				return nil, err
			} else {
				return config, nil
			}
		}
	}
}