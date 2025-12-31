package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// AppendTableContentReq contains parameters for appending content to a table.
// AppendTableContentReq 包含追加表格内容的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E8%BF%BD%E5%8A%A0%E8%A1%A8%E6%A0%BC%E5%86%85%E5%AE%B9
type AppendTableContentReq struct {
	Metadata
	// FileID is the unique identifier of the file.
	// FileID 是文件的唯一标识符。
	FileID string
	AppendTableContentReqBody
}

// AppendTableContentReqBody contains the request body for appending table content.
// AppendTableContentReqBody 包含追加表格内容的请求体。
type AppendTableContentReqBody struct {
	// Rg is the range of cells to append (e.g., "Sheet1!A1:C3").
	// Rg 是要追加的单元格范围（例如："Sheet1!A1:C3"）。
	Rg       string   `json:"range"`
	Resource Resource `json:"resource"` // 应该有 json tag
}

// Resource contains the values to append to the table.
// Resource 包含要追加到表格的值。
type Resource struct {
	// Values is the 2D array of cell values.
	// Values 是单元格值的二维数组。
	Values [][]interface{} `json:"values"`
}

// AppendTableContentRes contains the response for appending table content.
// AppendTableContentRes 包含追加表格内容的响应。
type AppendTableContentRes struct {
	rawRes
}

// NewAppendTableContentApi creates a new API config for appending table content.
// NewAppendTableContentApi 创建用于追加表格内容的 API 配置。
func NewAppendTableContentApi(cli *ehttp.Component, ss SignatureSigner, params AppendTableContentReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/sheets/values", params.FileID),
		Method: http.MethodPut,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body:    params.AppendTableContentReqBody,
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
		},
	}
}
