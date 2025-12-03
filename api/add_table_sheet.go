package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// AddTableSheetReq contains parameters for adding a new sheet to a table.
// AddTableSheetReq 包含新增表格工作表的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E6%96%B0%E5%A2%9E%E8%A1%A8%E6%A0%BC%E5%B7%A5%E4%BD%9C%E8%A1%A8
type AddTableSheetReq struct {
	Metadata
	// FileID is the unique identifier of the file.
	// FileID 是文件的唯一标识符。
	FileID string
	AddTableSheetReqBody
}

// AddTableSheetReqBody contains the request body for adding a new sheet.
// AddTableSheetReqBody 包含新增工作表的请求体。
type AddTableSheetReqBody struct {
	// Name is the name of the new sheet.
	// Name 是新工作表的名称。
	Name string `json:"name"`
}

// AddTableSheetRes contains the response for adding a new sheet.
// AddTableSheetRes 包含新增工作表的响应。
type AddTableSheetRes struct {
	rawRes
}

// NewAddTableSheetApi creates a new API config for adding a new sheet to a table.
// NewAddTableSheetApi 创建用于新增表格工作表的 API 配置。
func NewAddTableSheetApi(cli *ehttp.Component, ss SignatureSigner, params AddTableSheetReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
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
