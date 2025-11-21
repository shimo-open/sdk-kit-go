package api

import (
	"github.com/go-resty/resty/v2"
)

var Prefix string

type RawResponse struct {
	Resp *resty.Response `json:"-"`
}

func GetHeaderParams(auth Auth, extra map[string]string) map[string][]string {
	// 基础 header
	params := map[string][]string{
		"X-Shimo-Signature":     {auth.Signature},
		"X-Shimo-Token":         {auth.Token},
		"X-Weboffice-Token":     {auth.Token},
		"X-Weboffice-User-Uuid": {auth.UserUuid},
	}
	// 添加额外 header
	for k, v := range extra {
		params[k] = []string{v}
	}
	return params
}
