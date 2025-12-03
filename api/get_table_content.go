package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// GetTableContentReq contains parameters for getting table content.
// GetTableContentReq 包含获取表格内容的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#get-table-content
type GetTableContentReq struct {
	Metadata
	// FileID is the unique identifier of the file.
	// FileID 是文件的唯一标识符。
	FileID string
	// Rg is the range of cells to retrieve (e.g., "Sheet1!A1:C3").
	// Rg 是要获取的单元格范围（例如："Sheet1!A1:C3"）。
	Rg string
}

// GetTableContentRes contains the response for getting table content.
// GetTableContentRes 包含获取表格内容的响应。
type GetTableContentRes struct {
	rawRes
	// Values is the 2D array of cell values.
	// Values 是单元格值的二维数组。
	Values [][]interface{} `json:"values"`
	// Lag is the lag value.
	// Lag 是延迟值。
	Lag int `json:"lag"`
}

// NewGetTableContentApi creates a new API config for getting table content.
// NewGetTableContentApi 创建用于获取表格内容的 API 配置。
func NewGetTableContentApi(cli *ehttp.Component, ss SignatureSigner, params GetTableContentReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/sheets/values", params.FileID),
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Query: map[string][]string{
				"range": {params.Rg},
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetTableContentRes{},
		},
	}
}
