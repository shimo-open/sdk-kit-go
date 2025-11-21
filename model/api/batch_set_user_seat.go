package api

import (
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 批量设置用户席位
// https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E6%89%B9%E9%87%8F%E8%AE%BE%E7%BD%AE%E7%94%A8%E6%88%B7%E5%B8%AD%E4%BD%8D

type BatchSetUserSeatParams struct {
	Auth
	BatchSetUserSeatReqBody
	Status int
}

type BatchSetUserSeatReqBody struct {
	UserIds []string `json:"userIds"`
	Status  int      `json:"status"`
}

func NewBatchSetUserSeatApi(params BatchSetUserSeatParams) common.ApiConf {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return common.ApiConf{
		Url:    Prefix + "/sdk/v2/api/license/users/set-status",
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Body:    params.BatchSetUserSeatReqBody,
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusNoContent,
		},
	}
}
