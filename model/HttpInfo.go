package model

type HttpInfo struct {
	ClientIP string
	Method   string
	UrlPath  string
	Header   []byte
	Body     []byte
}
