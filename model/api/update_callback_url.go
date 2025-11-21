package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 更新App回调地址
// https://open.shimo.im/docs/06API-document/interface-description/system-interface#%E6%9B%B4%E6%96%B0app%E5%9B%9E%E8%B0%83%E5%9C%B0%E5%9D%80

type UpdateCallbackUrlParams struct {
	Auth
	AppId string
	UpdateCallbackUrlReqBody
}

type UpdateCallbackUrlReqBody struct {
	Url string `json:"url"`
}

func NewUpdateCallbackUrlApi(params UpdateCallbackUrlParams) common.ApiConf {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/api/license/apps/%s/endpoint-url", params.AppId),
		Method: http.MethodPut,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Body:    params.UpdateCallbackUrlReqBody,
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusNoContent,
		},
	}
}
