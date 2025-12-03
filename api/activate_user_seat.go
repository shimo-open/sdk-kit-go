package api

import (
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 激活用户席位
// https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E6%BF%80%E6%B4%BB%E7%94%A8%E6%88%B7%E5%B8%AD%E4%BD%8D

type ActivateUserSeatReq struct {
	Metadata
	ActivateUserSeatReqBody
}

type ActivateUserSeatReqBody struct {
	UserIds []string `json:"userIds"`
}

type ActivateUserSeatRes struct {
	rawRes
}

func NewActivateUserSeatApi(cli *ehttp.Component, ss SignatureSigner, params ActivateUserSeatReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeSystem)
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
