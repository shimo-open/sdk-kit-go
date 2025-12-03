package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// ExportFileReq contains parameters for exporting a file.
// ExportFileReq 包含导出文件的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/file-operation#export-v1
type ExportFileReq struct {
	Metadata
	// FileID is the unique identifier of the file to export.
	// FileID 是要导出的文件的唯一标识符。
	FileID string
	// Type is the export format type.
	// Type 是导出格式类型。
	Type string
}

// ExportFileRes contains the response for exporting a file.
// ExportFileRes 包含导出文件的响应。
type ExportFileRes struct {
	rawRes
	// Status is the export status (non-zero indicates error).
	// Status 是导出状态（非零值表示异常）。
	Status int `json:"status"`
	// Data contains the export result data.
	// Data 包含导出结果数据。
	Data ExportFileData `json:"data"`
	// Message is the error message when export fails.
	// Message 是导出失败时的提示信息。
	Message string `json:"message"`
}

// ExportFileData contains the export result data.
// ExportFileData 包含导出结果数据。
type ExportFileData struct {
	// TaskID is the export task identifier for progress tracking.
	// TaskID 是导出任务的标识 ID，用于跟踪进度。
	TaskID string `json:"taskId"`
}

// NewExportFileApi creates a new API config for exporting a file.
// NewExportFileApi 创建用于导出文件的 API 配置。
func NewExportFileApi(cli *ehttp.Component, ss SignatureSigner, params ExportFileReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/v1/export/%s", params.FileID),
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
			Body: map[string]interface{}{
				"type": params.Type,
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            ExportFileRes{},
		},
	}
}
