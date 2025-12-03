package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// CreateFileCopyReq contains parameters for creating a copy of a collaborative document.
// CreateFileCopyReq 包含创建协同文档副本的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E5%88%9B%E5%BB%BA%E5%8D%8F%E5%90%8C%E6%96%87%E6%A1%A3%E5%89%AF%E6%9C%AC
type CreateFileCopyReq struct {
	Metadata
	// OriginFileID is the unique identifier of the original file.
	// OriginFileID 是原始文件的唯一标识符。
	OriginFileID string
	// TargetFileID is the unique identifier for the new copy.
	// TargetFileID 是新副本的唯一标识符。
	TargetFileID string
}

// CreateFileCopyRequestBody contains the request body for creating a file copy.
// CreateFileCopyRequestBody 包含创建文件副本的请求体。
type CreateFileCopyRequestBody struct {
	// FileID is the target file ID for the copy.
	// FileID 是副本的目标文件 ID。
	FileID string `json:"fileId"`
}

// CreateFileCopyRes contains the response for creating a file copy.
// CreateFileCopyRes 包含创建文件副本的响应。
type CreateFileCopyRes struct {
	rawRes
}

// NewCreateFileCopyApi creates a new API config for creating a file copy.
// NewCreateFileCopyApi 创建用于创建文件副本的 API 配置。
func NewCreateFileCopyApi(cli *ehttp.Component, ss SignatureSigner, params CreateFileCopyReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/collab-files/%s/copy", params.OriginFileID),
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body: CreateFileCopyRequestBody{
				FileID: params.TargetFileID,
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
		},
	}
}
