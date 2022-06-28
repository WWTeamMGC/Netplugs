package Producer

import (
	"github.com/WWTeamMGC/Netplugs/Producer/kafka"
	"github.com/WWTeamMGC/Netplugs/Producer/setting"
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
	//msg := &sarama.Producer{}
	//msg.Topic = "aaaaaaa"
	//msg.Value = sarama.StringEncoder("ffffffffff")
	//kafka.ToMsgChan(msg)  //test

}
