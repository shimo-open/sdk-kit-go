package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 新增表格工作表
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E6%96%B0%E5%A2%9E%E8%A1%A8%E6%A0%BC%E5%B7%A5%E4%BD%9C%E8%A1%A8

type AddTableSheetReq struct {
	Metadata
	FileID string
	AddTableSheetReqBody
}

type AddTableSheetReqBody struct {
	Name string `json:"name"` // 新增表格工作表名称
}

type AddTableSheetRes struct {
	rawRes
}

func NewAddTableSheetApi(cli *ehttp.Component, ss SignatureSigner, params AddTableSheetReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/sheets", params.FileID),
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body:    params.AddTableSheetReqBody,
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
		},
	}
}
