package api

import (
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// BatchSetUserSeatReq contains parameters for batch setting user seats.
// BatchSetUserSeatReq 包含批量设置用户席位的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E6%89%B9%E9%87%8F%E8%AE%BE%E7%BD%AE%E7%94%A8%E6%88%B7%E5%B8%AD%E4%BD%8D
type BatchSetUserSeatReq struct {
	Metadata
	BatchSetUserSeatReqBody
	// Status is the target status to set.
	// Status 是要设置的目标状态。
	Status int
}

// BatchSetUserSeatReqBody contains the request body for batch setting user seats.
// BatchSetUserSeatReqBody 包含批量设置用户席位的请求体。
type BatchSetUserSeatReqBody struct {
	// UserIds is the list of user IDs to set.
	// UserIds 是要设置的用户 ID 列表。
	UserIds []string `json:"userIds"`
	// Status is the target status.
	// Status 是目标状态。
	Status int `json:"status"`
}

// BatchSetUserSeatRes contains the response for batch setting user seats.
// BatchSetUserSeatRes 包含批量设置用户席位的响应。
type BatchSetUserSeatRes struct {
	rawRes
}

// NewBatchSetUserSeatApi creates a new API config for batch setting user seats.
// NewBatchSetUserSeatApi 创建用于批量设置用户席位的 API 配置。
func NewBatchSetUserSeatApi(cli *ehttp.Component, ss SignatureSigner, params BatchSetUserSeatReq) *APIConf {
	sign := ss.Sign(ExpireShort, ScopeSystem)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/license/users/set-status",
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body:    params.BatchSetUserSeatReqBody,
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
		},
	}
}
