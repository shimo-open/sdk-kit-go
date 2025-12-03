package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// GetPlainTextReq contains parameters for getting file plain text content.
// GetPlainTextReq 包含获取文件纯文本内容的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E8%8E%B7%E5%8F%96%E6%96%87%E4%BB%B6%E7%BA%AF%E6%96%87%E6%9C%AC%E5%86%85%E5%AE%B9
type GetPlainTextReq struct {
	Metadata
	// FileID is the unique identifier of the file.
	// FileID 是文件的唯一标识符。
	FileID string
}

// GetPlainTextRes contains the response for getting file plain text content.
// GetPlainTextRes 包含获取文件纯文本内容的响应。
type GetPlainTextRes struct {
	rawRes
	// Content is the plain text content of the file.
	// Content 是文件的纯文本内容。
	Content string `json:"content"`
}

// NewGetPlainTextApi creates a new API config for getting file plain text content.
// NewGetPlainTextApi 创建用于获取文件纯文本内容的 API 配置。
func NewGetPlainTextApi(cli *ehttp.Component, ss SignatureSigner, params GetPlainTextReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/collab-files/%s/plain-text", params.FileID),
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetPlainTextRes{},
		},
	}
}
