package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gotomicro/ego/client/ehttp"
)

// DeleteTableRowReq contains parameters for deleting table rows.
// DeleteTableRowReq 包含删除表格行的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E5%88%A0%E9%99%A4%E8%A1%A8%E6%A0%BC%E8%A1%8C
type DeleteTableRowReq struct {
	Metadata
	// FileID is the unique identifier of the file.
	// FileID 是文件的唯一标识符。
	FileID string
	// SheetName is the name of the sheet.
	// SheetName 是工作表名称。
	SheetName string
	// Index is the starting row index to delete.
	// Index 是要删除的起始行索引。
	Index int
	// Count is the number of rows to delete.
	// Count 是要删除的行数。
	Count int
}

// DeleteTableRowRes contains the response for deleting table rows.
// DeleteTableRowRes 包含删除表格行的响应。
type DeleteTableRowRes struct {
	rawRes
}

// NewDeleteTableRowApi creates a new API config for deleting table rows.
// NewDeleteTableRowApi 创建用于删除表格行的 API 配置。
func NewDeleteTableRowApi(cli *ehttp.Component, ss SignatureSigner, params DeleteTableRowReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/sheets/%s/rows/%d", params.FileID, params.SheetName, params.Index),
		Method: http.MethodDelete,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Query: map[string][]string{
				"count": {strconv.Itoa(params.Count)},
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
		},
	}
}
