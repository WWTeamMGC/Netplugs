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
		SetIPFilter:     true,
		SetWordsFilter:  true,
		SetImagesFilter: false,
	}
	driver.InitServer(config)
	r.GET("/hello", Netplugs.NetGinPlug(config), func(c *gin.Context) {
		params := c.Param("name")
		c.JSON(200, gin.H{"Hello": 200, "name": params})
	})
	r.Run(":8081")
	//url := "http://127.0.0.1:8080/BadApi/Ip"
	//http.Client{
	//	Transport:     nil,
	//	CheckRedirect: nil,
	//	Jar:           nil,
	//	Timeout:       0,
	//}
	//type T struct {
	//	Badiplist []string `json:"badiplist"`
	//}
	//BadIPList := &T{}
	//s, err := req.Post(url)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//json.Unmarshal(s.Bytes(), &BadIPList)
}
