package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// UpdateTableContentReq contains parameters for updating table content.
// UpdateTableContentReq 包含更新表格内容的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E6%9B%B4%E6%96%B0%E8%A1%A8%E6%A0%BC%E5%86%85%E5%AE%B9
type UpdateTableContentReq struct {
	Metadata
	// FileID is the unique identifier of the file.
	// FileID 是文件的唯一标识符。
	FileID string
	UpdateTableContentRequestBody
}

// UpdateTableContentRes contains the response for updating table content.
// UpdateTableContentRes 包含更新表格内容的响应。
type UpdateTableContentRes struct {
	rawRes
}

// UpdateTableContentRequestBody contains the request body for updating table content.
// UpdateTableContentRequestBody 包含更新表格内容的请求体。
type UpdateTableContentRequestBody struct {
	// Rg is the range of cells to update (e.g., "Sheet1!A1:C3").
	// Rg 是要更新的单元格范围（例如："Sheet1!A1:C3"）。
	Rg string `json:"range"`
	// Resource contains the values to update.
	// Resource 包含要更新的值。
	Resource struct {
		// Values is the 2D array of cell values.
		// Values 是单元格值的二维数组。
		Values [][]interface{} `json:"values"`
	} `json:"resource"`
}

// NewUpdateTableContentApi creates a new API config for updating table content.
// NewUpdateTableContentApi 创建用于更新表格内容的 API 配置。
func NewUpdateTableContentApi(cli *ehttp.Component, ss SignatureSigner, params UpdateTableContentReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
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
