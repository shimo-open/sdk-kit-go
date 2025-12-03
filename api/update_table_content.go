package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 更新表格内容
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E6%9B%B4%E6%96%B0%E8%A1%A8%E6%A0%BC%E5%86%85%E5%AE%B9

type UpdateTableContentReq struct {
	Metadata
	FileID string
	UpdateTableContentRequestBody
}

type UpdateTableContentRes struct {
	rawRes
}

type UpdateTableContentRequestBody struct {
	Rg       string `json:"range"`
	Resource struct {
		Values [][]interface{} `json:"values"`
	} `json:"resource"`
}

func NewUpdateTableContentApi(cli *ehttp.Component, ss SignatureSigner, params UpdateTableContentReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/sheets/values", params.FileID),
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body:    params.UpdateTableContentRequestBody,
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
		},
	}
}
