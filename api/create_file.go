package api

import (
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// CreateFileReq contains parameters for creating a file.
// CreateFileReq 包含创建文件的参数。
type CreateFileReq struct {
	Metadata
	FileType   CollabFileType
	FileID     string
	Lang       Lang
	ContentKey string
}

type CreateFileRes struct {
	rawRes
}

// CreateFileRequestBody contains the request body for creating a file.
// CreateFileRequestBody 包含创建文件的请求体。
type CreateFileRequestBody struct {
	Type       CollabFileType `json:"type"`
	FileID     string         `json:"fileId"`
	ContentKey string         `json:"contentKey"`
}

// NewCreateFileApi creates a new API config for creating a file.
// NewCreateFileApi 创建用于创建文件的 API 配置。
func NewCreateFileApi(cli *ehttp.Component, ss SignatureSigner, params CreateFileReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
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
