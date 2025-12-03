package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// GetPlainTextWCReq contains parameters for getting file plain text word count.
// GetPlainTextWCReq 包含获取文件纯文本字数统计的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E6%96%87%E4%BB%B6%E7%BA%AF%E6%96%87%E6%9C%AC%E5%AD%97%E6%95%B0%E7%BB%9F%E8%AE%A1
type GetPlainTextWCReq struct {
	Metadata
	// FileID is the unique identifier of the file.
	// FileID 是文件的唯一标识符。
	FileID string
}

// GetPlainTextWCRes contains the response for getting file plain text word count.
// GetPlainTextWCRes 包含获取文件纯文本字数统计的响应。
type GetPlainTextWCRes struct {
	rawRes
	// WordCount is the word count of the file's plain text content.
	// WordCount 是文件纯文本内容的字数。
	WordCount int `json:"wordCount"`
	// Keywords is the keyword frequency map (e.g., {"foo":1,"bar":10}).
	// Keywords 是关键词频率映射（例如：{"foo":1,"bar":10}）。
	Keywords interface{} `json:"keywords"`
}

// NewGetPlainTextWCApi creates a new API config for getting file plain text word count.
// NewGetPlainTextWCApi 创建用于获取文件纯文本字数统计的 API 配置。
func NewGetPlainTextWCApi(cli *ehttp.Component, ss SignatureSigner, params GetPlainTextWCReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/collab-files/%s/plain-text/wc", params.FileID),
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetPlainTextWCRes{},
		},
	}
}
