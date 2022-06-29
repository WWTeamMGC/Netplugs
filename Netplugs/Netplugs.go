package Netplugs

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Shopify/sarama"
	"github.com/WWTeamMGC/Netplugs/Producer/kafka"
	"github.com/WWTeamMGC/Netplugs/model"
	"github.com/gin-gonic/gin"
)

func NetGinPlug() gin.HandlerFunc {
	return func(c *gin.Context) {
		go func() {
			var s []map[string]interface{}
			body, _ := ioutil.ReadAll(c.Request.Body)
			for k, v := range c.Request.Header {
				s = append(s, map[string]interface{}{k: v})
			}
			// headerjson, err := json.Marshal(s)
			// if err != nil {
			// 	return
			// }
			httpmsg := model.HttpInfo{
				ClientIP: c.ClientIP(),
				Method:   c.Request.Method,
				UrlPath:  c.Request.URL.Path,
				Header:   s,
				Body:     body,
			}
			msgjson, err := json.Marshal(httpmsg)
			if err != nil {
				return
			}
			msg := &sarama.ProducerMessage{
				Topic: "test",
				Value: sarama.ByteEncoder(msgjson),
			}
			kafka.ToMsgChan(msg)
		}()
		//length := c.Request.ContentLength
		c.Next()
	}
}
