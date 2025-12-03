package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 更新App回调地址
// https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E6%9B%B4%E6%96%B0app%E5%9B%9E%E8%B0%83%E5%9C%B0%E5%9D%80

type UpdateCallbackURLReq struct {
	Metadata
	AppID string
	UpdateCallbackURLReqBody
}

type UpdateCallbackURLReqBody struct {
	URL string `json:"url"`
}

type UpdateCallbackURLRes struct {
	rawRes
}

func NewUpdateCallbackURLApi(cli *ehttp.Component, ss SignatureSigner, params UpdateCallbackURLReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeSystem)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/license/apps/%s/endpoint-url", params.AppID),
		Method: http.MethodPut,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body:    params.UpdateCallbackURLReqBody,
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
		},
	}
}
