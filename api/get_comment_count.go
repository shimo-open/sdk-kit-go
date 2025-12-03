package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// GetCommentCountReq contains parameters for getting the comment count of a file.
// GetCommentCountReq 包含获取文件评论数的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#get-comment-count
type GetCommentCountReq struct {
	Metadata
	// FileID is the unique identifier of the file.
	// FileID 是文件的唯一标识符。
	FileID string
}

// GetCommentCountRes contains the response for getting the comment count of a file.
// GetCommentCountRes 包含获取文件评论数的响应。
type GetCommentCountRes struct {
	rawRes
	// Count is the number of comments in the file.
	// Count 是文件中的评论数。
	Count int `json:"count"`
}

// NewGetCommentCountApi creates a new API config for getting the comment count of a file.
// NewGetCommentCountApi 创建用于获取文件评论数的 API 配置。
func NewGetCommentCountApi(cli *ehttp.Component, ss SignatureSigner, params GetCommentCountReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/collab-files/%s/comment-count", params.FileID),
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetCommentCountRes{},
		},
	}
}
