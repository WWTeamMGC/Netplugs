// go:build

package main

import (
	"github.com/WWTeamMGC/Netplugs/Config"
	"github.com/WWTeamMGC/Netplugs/Netplugs"
	"github.com/WWTeamMGC/Netplugs/driver"
	_ "github.com/WWTeamMGC/Netplugs/driver"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := &Config.SetConfig{
		SetIPFilter:     false,
		SetWordsFilter:  false,
		SetImagesFilter: false,
	}
	driver.InitServer(config)
	r.Use(Netplugs.NetGinPlug(config))
	r.GET("/hello", func(c *gin.Context) {
		params := c.Param("name")
		c.JSON(200, gin.H{"Hello": 200, "name": params})
	})
	r.GET("/shop/cat", func(c *gin.Context) {
		params := c.Param("name")
		c.JSON(200, gin.H{"Hello": 200, "name": params})
	})
	r.GET("/shop/dog", func(c *gin.Context) {
		params := c.Param("name")
		c.JSON(200, gin.H{"Hello": 200, "name": params})
	})

	r.GET("/pay", func(c *gin.Context) {
		params := c.Param("name")
		c.JSON(200, gin.H{"Hello": 200, "name": params})
	})
	r.GET("/study/english", func(c *gin.Context) {
		params := c.Param("name")
		c.JSON(200, gin.H{"Hello": 200, "name": params})
	})
	r.GET("/study/math", func(c *gin.Context) {
		params := c.Param("name")
		c.JSON(200, gin.H{"Hello": 200, "name": params})
	})

	r.Run(":8085")
}
