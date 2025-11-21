package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 获取表格内容
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#get-table-content

type GetTableContentParams struct {
	Auth
	FileId string
	Rg     string
}

type GetTableContentRespBody struct {
	RawResponse
	Values [][]interface{} `json:"values"`
	Lag    int             `json:"lag"`
}

func NewGetTableContentApi(params GetTableContentParams) common.ApiConf {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/api/files/%s/sheets/values", params.FileId),
		Method: http.MethodGet,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Query: map[string][]string{
				"range": {params.Rg},
			},
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            GetTableContentRespBody{},
		},
	}
}
