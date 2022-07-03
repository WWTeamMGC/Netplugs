//go:build

package main

import (
	_ "github.com/WWTeamMGC/Netplugs/driver"
)

func main() {
	//r := gin.Default()
	//r.GET("/hello", Netplugs.NetGinPlug(), func(c *gin.Context) {
	//	params := c.Param("name")
	//	c.JSON(200, gin.H{"Hello": 200, "name": params})
	//})
	//r.Run(":8081")
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
