package Netplugs

import (
	"encoding/json"
	"io/ioutil"

	"github.com/WWTeamMGC/Netplugs/model"
	"github.com/gin-gonic/gin"
)

func NetGinPlug() gin.HandlerFunc {
	return func(c *gin.Context) {
		var s []map[string]interface{}
		body, _ := ioutil.ReadAll(c.Request.Body)
		for k, v := range c.Request.Header {
			s = append(s, map[string]interface{}{k: v})
		}
		b, err := json.Marshal(s)
		if err != nil {
			return
		}
		msg := &model.HttpInfo{
			ClientIP: c.ClientIP(),
			Method:   c.Request.Method,
			UrlPath:  c.Request.URL.Path,
			Header:   b,
			Body:     body,
		}
		//length := c.Request.ContentLength
		c.Next()
	}
}
