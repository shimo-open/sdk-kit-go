package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 新增表格工作表
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E6%96%B0%E5%A2%9E%E8%A1%A8%E6%A0%BC%E5%B7%A5%E4%BD%9C%E8%A1%A8

type AddTableSheetParams struct {
	Auth
	FileId string
	AddTableSheetReqBody
}

type AddTableSheetReqBody struct {
	Name string `json:"name"` // 新增表格工作表名称
}

func NewAddTableSheetApi(params AddTableSheetParams) common.ApiConf {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/api/files/%s/sheets", params.FileId),
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Body:    params.AddTableSheetReqBody,
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusNoContent,
		},
	}
}
