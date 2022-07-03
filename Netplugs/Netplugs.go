package Netplugs

import (
	"encoding/json"
	"github.com/WWTeamMGC/Netplugs/Config"
	"github.com/WWTeamMGC/Netplugs/driver"
	"io/ioutil"
	"net/http"

	"github.com/Shopify/sarama"
	"github.com/WWTeamMGC/Netplugs/Producer/kafka"
	"github.com/WWTeamMGC/Netplugs/model"
	"github.com/gin-gonic/gin"
)

type ErrRsp struct {
	Msg string
}

func NetGinPlug(Config *Config.SetConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		if Config.SetIPFilter {
			s, _ := driver.BadIPServer.Match(c.ClientIP(), '*')
			if len(s) != 0 {
				c.JSON(http.StatusOK, &ErrRsp{Msg: "IP已被封禁"})
				c.Abort()
				return
			}
		}
		body, _ := ioutil.ReadAll(c.Request.Body)
		if Config.SetWordsFilter {
			s, _ := driver.BadWordsServer.Match(string(body), '*')
			if len(s) != 0 {
				c.JSON(http.StatusOK, &ErrRsp{Msg: "内容含有违禁词"})
				c.Abort()
				return
			}
		}
		go func() {
			var s []map[string]interface{}
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
