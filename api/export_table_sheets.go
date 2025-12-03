package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// ExportTableSheetsReq contains parameters for exporting table sheets to Excel.
// ExportTableSheetsReq 包含导出应用表格为 Excel 的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/file-operation#export-table-sheets
type ExportTableSheetsReq struct {
	Metadata
	// FileID is the unique identifier of the file to export.
	// FileID 是要导出的文件的唯一标识符。
	FileID string
}

// ExportTableSheetsRes contains the response for exporting table sheets.
// ExportTableSheetsRes 包含导出应用表格的响应。
type ExportTableSheetsRes struct {
	rawRes
	// Status is the export status (non-zero indicates error).
	// Status 是导出状态（非零值表示异常）。
	Status int `json:"status"`
	// Message is the error message when export fails.
	// Message 是导出失败时的提示信息。
	Message string `json:"message"`
	// DownloadUrl is the download URL of the exported .xlsx file.
	// DownloadUrl 是导出的 .xlsx 文件下载地址。
	DownloadUrl string `json:"downloadUrl"`
}

// NewExportTableSheetsApi creates a new API config for exporting table sheets to Excel.
// NewExportTableSheetsApi 创建用于导出应用表格为 Excel 的 API 配置。
func NewExportTableSheetsApi(cli *ehttp.Component, ss SignatureSigner, params ExportTableSheetsReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/export/table-sheets/%s", params.FileID),
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            ExportTableSheetsRes{},
		},
	}
}
