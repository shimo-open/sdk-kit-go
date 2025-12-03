package api

import (
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// CreateFileReq contains parameters for creating a collaborative file.
// CreateFileReq 包含创建协同文件的参数。
type CreateFileReq struct {
	Metadata
	// FileType is the type of the collaborative file.
	// FileType 是协同文件的类型。
	FileType CollabFileType
	// FileID is the unique identifier of the file.
	// FileID 是文件的唯一标识符。
	FileID string
	// Lang is the language for the file.
	// Lang 是文件的语言。
	Lang Lang
	// ContentKey is the key for file content (optional).
	// ContentKey 是文件内容的键（可选）。
	ContentKey string
}

// CreateFileRes contains the response for creating a file.
// CreateFileRes 包含创建文件的响应。
type CreateFileRes struct {
	rawRes
}

// CreateFileRequestBody contains the request body for creating a file.
// CreateFileRequestBody 包含创建文件的请求体。
type CreateFileRequestBody struct {
	// Type is the collaborative file type.
	// Type 是协同文件类型。
	Type CollabFileType `json:"type"`
	// FileID is the unique identifier of the file.
	// FileID 是文件的唯一标识符。
	FileID string `json:"fileId"`
	// ContentKey is the key for file content.
	// ContentKey 是文件内容的键。
	ContentKey string `json:"contentKey"`
}

// NewCreateFileApi creates a new API config for creating a file.
// NewCreateFileApi 创建用于创建文件的 API 配置。
func NewCreateFileApi(cli *ehttp.Component, ss SignatureSigner, params CreateFileReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{
		"Content-Type":    "application/json",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/files",
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Query:   map[string][]string{"lang": {params.Lang.String()}},
			Body: CreateFileRequestBody{
				Type:       params.FileType,
				FileID:     params.FileID,
				ContentKey: params.ContentKey,
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
		},
	}
}
