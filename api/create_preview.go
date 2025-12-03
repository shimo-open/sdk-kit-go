package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// CreatePreviewReq contains parameters for creating a file preview.
// CreatePreviewReq 包含创建文件预览的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/file-operation#%E5%88%9B%E5%BB%BA%E9%A2%84%E8%A7%88
type CreatePreviewReq struct {
	Metadata
	// FileID is the unique identifier of the file.
	// FileID 是文件的唯一标识符。
	FileID string
}

// CreatePreviewRes contains the response for creating a file preview.
// CreatePreviewRes 包含创建文件预览的响应。
type CreatePreviewRes struct {
	rawRes
	// Code is the result status code (empty string means success).
	// Code 是结果状态码（空字符串表示成功）。
	Code string `json:"code"`
	// Message is the error message when creation fails.
	// Message 是创建失败时的错误信息。
	Message string `json:"message"`
}

// NewCreatePreviewApi creates a new API config for creating a file preview.
// NewCreatePreviewApi 创建用于创建文件预览的 API 配置。
func NewCreatePreviewApi(cli *ehttp.Component, ss SignatureSigner, params CreatePreviewReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/cloud-files/%s/create", params.FileID),
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            CreatePreviewRes{},
		},
	}
}
