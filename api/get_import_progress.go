package api

import (
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// GetImportProgReq contains parameters for getting import progress.
// GetImportProgReq 包含获取导入进度的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/file-operation#import-progress-v1
type GetImportProgReq struct {
	Metadata
	// TaskId is the import task identifier.
	// TaskId 是导入任务标识符。
	TaskId string
}

// GetImportProgRes contains the response for getting import progress.
// GetImportProgRes 包含获取导入进度的响应。
type GetImportProgRes struct {
	rawRes
	// Status is the import status (non-zero indicates error).
	// Status 是导入状态（非零值表示异常）。
	Status int `json:"status"`
	// Message is the error message when import fails.
	// Message 是导入失败时的提示信息。
	Message string `json:"message"`
	// Data contains the import progress data.
	// Data 包含导入进度数据。
	Data GetImportProgData `json:"data"`
}

// GetImportProgData contains the import progress data.
// GetImportProgData 包含导入进度数据。
type GetImportProgData struct {
	// Progress is the import progress percentage (100 means complete).
	// Progress 是导入进度百分比（100 表示完成）。
	Progress int `json:"progress"`
}

// NewGetImportProgressApi creates a new API config for getting import progress.
// NewGetImportProgressApi 创建用于获取导入进度的 API 配置。
func NewGetImportProgressApi(cli *ehttp.Component, ss SignatureSigner, params GetImportProgReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/files/v1/import/progress",
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
			Query: map[string][]string{
				"taskId": {params.TaskId},
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetImportProgRes{},
		},
	}
}

// NewGetImportV2ProgressApi creates a new API config for getting import progress using v2 API.
// NewGetImportV2ProgressApi 创建用于使用 v2 接口获取导入进度的 API 配置。
func NewGetImportV2ProgressApi(cli *ehttp.Component, ss SignatureSigner, params GetImportProgReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/files/v2/import/progress",
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body: map[string]string{
				"taskId": params.TaskId,
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetImportProgRes{},
		},
	}
}
