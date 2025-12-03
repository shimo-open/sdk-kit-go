package api

import (
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// GetExportProgReq contains parameters for getting export progress.
// GetExportProgReq 包含获取导出进度的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/file-operation#export-progress-v1
type GetExportProgReq struct {
	Metadata
	// TaskId is the export task identifier.
	// TaskId 是导出任务标识符。
	TaskId string
}

// GetExportProgRes contains the response for getting export progress.
// GetExportProgRes 包含获取导出进度的响应。
type GetExportProgRes struct {
	rawRes
	// Status is the export status (non-zero indicates error).
	// Status 是导出状态（非零值表示异常）。
	Status int `json:"status"`
	// Message is the error message when export fails.
	// Message 是导出失败时的提示信息。
	Message string `json:"message"`
	// Data contains the export progress data.
	// Data 包含导出进度数据。
	Data GetExportProgRespData `json:"data"`
}

// GetExportProgRespData contains the export progress data.
// GetExportProgRespData 包含导出进度数据。
type GetExportProgRespData struct {
	// Progress is the export progress percentage (100 means complete).
	// Progress 是导出进度百分比（100 表示完成）。
	Progress int `json:"progress"`
	// DownloadUrl is the download URL of the exported file.
	// DownloadUrl 是导出文件的下载地址。
	DownloadUrl string `json:"downloadUrl"`
}

// NewGetExportProgressApi creates a new API config for getting export progress.
// NewGetExportProgressApi 创建用于获取导出进度的 API 配置。
func NewGetExportProgressApi(cli *ehttp.Component, ss SignatureSigner, params GetExportProgReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/files/v1/export/progress",
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
			Query: map[string][]string{
				"taskId": {params.TaskId},
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetExportProgRes{},
		},
	}
}
