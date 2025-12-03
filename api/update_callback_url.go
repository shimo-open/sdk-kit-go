package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// UpdateCallbackURLReq contains parameters for updating app callback URL.
// UpdateCallbackURLReq 包含更新应用回调地址的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E6%9B%B4%E6%96%B0app%E5%9B%9E%E8%B0%83%E5%9C%B0%E5%9D%80
type UpdateCallbackURLReq struct {
	Metadata
	// AppID is the application identifier.
	// AppID 是应用标识符。
	AppID string
	UpdateCallbackURLReqBody
}

// UpdateCallbackURLReqBody contains the request body for updating callback URL.
// UpdateCallbackURLReqBody 包含更新回调地址的请求体。
type UpdateCallbackURLReqBody struct {
	// URL is the new callback URL.
	// URL 是新的回调地址。
	URL string `json:"url"`
}

// UpdateCallbackURLRes contains the response for updating callback URL.
// UpdateCallbackURLRes 包含更新回调地址的响应。
type UpdateCallbackURLRes struct {
	rawRes
}

// NewUpdateCallbackURLApi creates a new API config for updating app callback URL.
// NewUpdateCallbackURLApi 创建用于更新应用回调地址的 API 配置。
func NewUpdateCallbackURLApi(cli *ehttp.Component, ss SignatureSigner, params UpdateCallbackURLReq) *APIConf {
	sign := ss.Sign(ExpireShort, ScopeSystem)
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
