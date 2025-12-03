package api

import (
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// ActivateUserSeatReq contains parameters for activating user seats.
// ActivateUserSeatReq 包含激活用户席位的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E6%BF%80%E6%B4%BB%E7%94%A8%E6%88%B7%E5%B8%AD%E4%BD%8D
type ActivateUserSeatReq struct {
	Metadata
	ActivateUserSeatReqBody
}

// ActivateUserSeatReqBody contains the request body for activating user seats.
// ActivateUserSeatReqBody 包含激活用户席位的请求体。
type ActivateUserSeatReqBody struct {
	// UserIds is the list of user IDs to activate.
	// UserIds 是要激活的用户 ID 列表。
	UserIds []string `json:"userIds"`
}

// ActivateUserSeatRes contains the response for activating user seats.
// ActivateUserSeatRes 包含激活用户席位的响应。
type ActivateUserSeatRes struct {
	rawRes
}

// NewActivateUserSeatApi creates a new API config for activating user seats.
// NewActivateUserSeatApi 创建用于激活用户席位的 API 配置。
func NewActivateUserSeatApi(cli *ehttp.Component, ss SignatureSigner, params ActivateUserSeatReq) *APIConf {
	sign := ss.Sign(ExpireShort, ScopeSystem)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/license/users/activate",
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body:    params.ActivateUserSeatReqBody,
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
		},
	}
}
