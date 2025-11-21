package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 更新表格内容
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E6%9B%B4%E6%96%B0%E8%A1%A8%E6%A0%BC%E5%86%85%E5%AE%B9

type UpdateTableContentParams struct {
	Auth
	FileId string
	UpdateTableContentRequestBody
}

type UpdateTableContentRequestBody struct {
	Rg       string `json:"range"`
	Resource struct {
		Values [][]interface{} `json:"values"`
	} `json:"resource"`
}

func NewUpdateTableContentApi(params UpdateTableContentParams) common.ApiConf {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/api/files/%s/sheets/values", params.FileId),
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Body:    params.UpdateTableContentRequestBody,
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusNoContent,
		},
	}
}
