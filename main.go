package main

import (
	"github.com/WWTeamMGC/Netplugs/Netplugs"
	_ "github.com/WWTeamMGC/Netplugs/Producer"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", Netplugs.NetGinPlug(), func(c *gin.Context) {
		params := c.Param("name")
		c.JSON(200, gin.H{"Hello": 200, "name": params})
	})
	r.Run(":8081")
}
