package api

import (
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 批量设置用户席位
// https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E6%89%B9%E9%87%8F%E8%AE%BE%E7%BD%AE%E7%94%A8%E6%88%B7%E5%B8%AD%E4%BD%8D

type BatchSetUserSeatReq struct {
	Metadata
	BatchSetUserSeatReqBody
	Status int
}

type BatchSetUserSeatReqBody struct {
	UserIds []string `json:"userIds"`
	Status  int      `json:"status"`
}

type BatchSetUserSeatRes struct {
	rawRes
}

func NewBatchSetUserSeatApi(cli *ehttp.Component, ss SignatureSigner, params BatchSetUserSeatReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeSystem)
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
