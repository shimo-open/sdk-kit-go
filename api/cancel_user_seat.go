package api

import (
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 取消用户席位
// https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E5%8F%96%E6%B6%88%E7%94%A8%E6%88%B7%E5%B8%AD%E4%BD%8D

type CancelUserSeatReq struct {
	Metadata
	CancelUserSeatReqBody
}

type CancelUserSeatReqBody struct {
	UserIds []string `json:"userIds"`
}

type CancelUserSeatRes struct {
	rawRes
}

func NewCancelUserSeatApi(cli *ehttp.Component, ss SignatureSigner, params CancelUserSeatReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeSystem)
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
