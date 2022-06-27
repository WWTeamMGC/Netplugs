package Netplugs

type Context struct {
	Method  string
	Header  []Header
	UrlPath string
	body    []byte
}
type Header struct {
	Name string
}

func GetConText(method string, url string, header []Header, body []byte) (bool, error) {
	return true, nil
}
