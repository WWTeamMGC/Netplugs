package model

type HttpInfo struct {
	ClientIP string
	Method   string
	UrlPath  string
	Header   []map[string]interface{}
	Body     []byte
}
