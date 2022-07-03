package Netplugs

import (
	"encoding/json"
	"github.com/WWTeamMGC/Netplugs/Config"
	"github.com/WWTeamMGC/Netplugs/driver"
	"io/ioutil"

	"github.com/Shopify/sarama"
	"github.com/WWTeamMGC/Netplugs/Producer/kafka"
	"github.com/WWTeamMGC/Netplugs/model"
	"github.com/gin-gonic/gin"
)

func NetGinPlug() gin.HandlerFunc {
	return func(c *gin.Context) {
		go func() {
			if Config.Config.SetIPFilter {
				s, _ := driver.BadIPServer.Match(c.ClientIP(), '*')
				if len(s) != 0 {
					c.Abort()
				}
			}
			var s []map[string]interface{}
			body, _ := ioutil.ReadAll(c.Request.Body)
			if Config.Config.SetWordsFilter {
				s, _ := driver.BadWordsServer.Match(string(body), '*')
				if len(s) != 0 {
					c.Abort()
				}
			}
			for k, v := range c.Request.Header {
				s = append(s, map[string]interface{}{k: v})
			}
			news, err := json.Marshal(s)
			//TODO fix err
			if err != nil {
				return
			}
			httpmsg := model.HttpInfo{
				ClientIP: c.ClientIP(),
				Method:   c.Request.Method,
				UrlPath:  c.Request.URL.Path,
				Header:   string(news),
				Body:     string(body),
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
		c.Next()
	}
}
