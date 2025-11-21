package api

import (
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 取消用户席位
// https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E5%8F%96%E6%B6%88%E7%94%A8%E6%88%B7%E5%B8%AD%E4%BD%8D

type CancelUserSeatParams struct {
	Auth
	CancelUserSeatReqBody
}

type CancelUserSeatReqBody struct {
	UserIds []string `json:"userIds"`
}

func NewCancelUserSeatApi(params CancelUserSeatParams) common.ApiConf {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return common.ApiConf{
		Url:    Prefix + "/sdk/v2/api/license/users/deactivate",
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Body:    params.CancelUserSeatReqBody,
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusNoContent,
		},
	}
}
