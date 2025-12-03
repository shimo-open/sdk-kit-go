package api

import (
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// CancelUserSeatReq contains parameters for canceling user seats.
// CancelUserSeatReq 包含取消用户席位的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E5%8F%96%E6%B6%88%E7%94%A8%E6%88%B7%E5%B8%AD%E4%BD%8D
type CancelUserSeatReq struct {
	Metadata
	CancelUserSeatReqBody
}

// CancelUserSeatReqBody contains the request body for canceling user seats.
// CancelUserSeatReqBody 包含取消用户席位的请求体。
type CancelUserSeatReqBody struct {
	// UserIds is the list of user IDs to cancel.
	// UserIds 是要取消的用户 ID 列表。
	UserIds []string `json:"userIds"`
}

// CancelUserSeatRes contains the response for canceling user seats.
// CancelUserSeatRes 包含取消用户席位的响应。
type CancelUserSeatRes struct {
	rawRes
}

// NewCancelUserSeatApi creates a new API config for canceling user seats.
// NewCancelUserSeatApi 创建用于取消用户席位的 API 配置。
func NewCancelUserSeatApi(cli *ehttp.Component, ss SignatureSigner, params CancelUserSeatReq) *APIConf {
	sign := ss.Sign(ExpireShort, ScopeSystem)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/license/users/deactivate",
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body:    params.CancelUserSeatReqBody,
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
		},
	}
}
