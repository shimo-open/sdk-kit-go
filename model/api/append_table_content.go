package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 追加表格内容
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E8%BF%BD%E5%8A%A0%E8%A1%A8%E6%A0%BC%E5%86%85%E5%AE%B9

type AppendTableContentParams struct {
	Auth
	FileId string
	AppendTableContentReqBody
}

type AppendTableContentReqBody struct {
	Rg       string `json:"range"`
	Resource `json:"resource"`
}

type Resource struct {
	Values [][]interface{} `json:"values"`
}

func NewAppendTableContentApi(params AppendTableContentParams) common.ApiConf {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/api/files/%s/sheets/values", params.FileId),
		Method: http.MethodPut,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Body:    params.AppendTableContentReqBody,
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusNoContent,
		},
	}
}
