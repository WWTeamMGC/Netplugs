package model

type HttpInfo struct {
	ClientIP string
	Method   string
	UrlPath  string
	Header   string
	Body     string
}
type Badwordslist struct {
	Badwordslist []string `json:"badwordslist"`
}
type Badiplist struct {
	Badiplist []BadIPListRsp `json:"badiplist"`
}
type BadIPListRsp struct {
	Ip      string `json:"ip"`
	PcMp    string `json:"pc_mp"`
	Address string `json:"address"`
}
