package common

import (
	"crypto/tls"

	"github.com/go-resty/resty/v2"
	"github.com/gotomicro/ego/core/econf"
)

type ApiConf struct {
	Url        string
	Method     string
	ReqParams  ApiRequestParams
	RespParams ApiResponseParams
}

type ApiRequestParams struct {
	Headers map[string][]string
	Query   map[string][]string
	Body    interface{}
	Form    map[string][]string
}

type ApiResponseParams struct {
	SuccessHttpCode int // 期望的 http 状态码
	// Headers         map[string][]string
	Body interface{}
}

func (ac *ApiConf) Request() *resty.Request {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: econf.GetBool("skipTlsVerify")})
	req := client.R()
	req.URL = ac.Url
	req.Method = ac.Method
	req.Header = ac.ReqParams.Headers
	req.QueryParam = ac.ReqParams.Query
	req.Body = ac.ReqParams.Body
	req.FormData = ac.ReqParams.Form
	return req
}
