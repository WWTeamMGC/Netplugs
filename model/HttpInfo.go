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
	Badiplist []string `json:"badiplist"`
}
