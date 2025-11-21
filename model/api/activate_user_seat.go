package api

import (
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 激活用户席位
// https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E6%BF%80%E6%B4%BB%E7%94%A8%E6%88%B7%E5%B8%AD%E4%BD%8D

type ActivateUserSeatParams struct {
	Auth
	ActivateUserSeatReqBody
}

type ActivateUserSeatReqBody struct {
	UserIds []string `json:"userIds"`
}

func NewActivateUserSeatApi(params ActivateUserSeatParams) common.ApiConf {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return common.ApiConf{
		Url:    Prefix + "/sdk/v2/api/license/users/activate",
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Body:    params.ActivateUserSeatReqBody,
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusNoContent,
		},
	}
}
