package driver

import (
	"encoding/json"
	"fmt"
	"github.com/WWTeamMGC/Netplugs/Config"
	"github.com/WWTeamMGC/Netplugs/Producer/kafka"
	"github.com/WWTeamMGC/Netplugs/Producer/setting"
	"github.com/WWTeamMGC/Netplugs/match"
	"github.com/WWTeamMGC/Netplugs/model"
	"github.com/imroc/req"
	"time"
)

var (
	Badip          = model.Badiplist{}
	Badwords       = model.Badwordslist{}
	BadIPServer    *match.MatchService
	BadWordsServer *match.MatchService
)

func init() {
	err := setting.Init()
	if err != nil {
		panic(err)
	}
	err = kafka.Init([]string{setting.GetConf().KafkaConfig.Address}, setting.GetConf().KafkaConfig.ChanSize)
	if err != nil {
		panic(err)
	}
}
func InitServer(config *Config.SetConfig) {
	if config.SetIPFilter {
		go ReFlushIP()
	}
	if config.SetIPFilter {
		go ReFlushWords()
	}
	time.Sleep(5 * time.Second)
}
func FlushIPList() model.Badiplist {
	url := "http://127.0.0.1:8080/BadApi/Ip"
	BadIPList := model.Badiplist{}
	s, err := req.Post(url)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(s.Bytes(), &BadIPList)
	return BadIPList
}
func FlushWordsList() model.Badwordslist {
	url := "http://127.0.0.1:8080/BadApi/Words"
	BadWordsList := model.Badwordslist{}
	s, err := req.Post(url)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(s.Bytes(), &BadWordsList)
	return BadWordsList
}

//定时30秒刷新IP

func ReFlushIP() {
	for {
		Badip = FlushIPList()
		var Iplist []string
		BadIPServer = match.NewMatchService()
		for _, v := range Badip.Badiplist {
			Iplist = append(Iplist, v.Ip)
		}
		BadIPServer.Build(Iplist)
		time.Sleep(30 * time.Second)
	}
}
func ReFlushWords() {
	for {
		Badwords = FlushWordsList()
		BadWordsServer = match.NewMatchService()
		BadWordsServer.Build(Badwords.Badwordslist)
		time.Sleep(30 * time.Second)
	}
}
