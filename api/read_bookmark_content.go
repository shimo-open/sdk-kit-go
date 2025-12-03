package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// ReadBookmarkContentReq contains parameters for reading bookmark content from a traditional document.
// ReadBookmarkContentReq 包含读取传统文档书签内容的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E8%AF%BB%E5%8F%96%E4%BC%A0%E7%BB%9F%E6%96%87%E6%A1%A3%E4%B9%A6%E7%AD%BE%E5%86%85%E5%AE%B9
type ReadBookmarkContentReq struct {
	Metadata
	// FileID is the unique identifier of the file.
	// FileID 是文件的唯一标识符。
	FileID string
	// Bookmarks is the list of bookmark names to read.
	// Bookmarks 是要读取的书签名称列表。
	Bookmarks []string
}

// ReadBookmarkContentRes contains the response for reading bookmark content.
// ReadBookmarkContentRes 包含读取书签内容的响应。
type ReadBookmarkContentRes struct {
	rawRes
	// Data is the list of bookmark data.
	// Data 是书签数据列表。
	Data []Data `json:"data"`
}

// Data represents a single bookmark data entry.
// Data 表示单个书签数据条目。
type Data struct {
	// Bookmark is the bookmark name.
	// Bookmark 是书签名称。
	Bookmark string `json:"bookmark"`
	// Content is the bookmark content.
	// Content 是书签内容。
	Content string `json:"content"`
}

// NewReadBookmarkContentApi creates a new API config for reading bookmark content.
// NewReadBookmarkContentApi 创建用于读取书签内容的 API 配置。
func NewReadBookmarkContentApi(cli *ehttp.Component, ss SignatureSigner, params ReadBookmarkContentReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/documentpro/bookmark_content", params.FileID),
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Query: map[string][]string{
				"bookmarks": params.Bookmarks,
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            ReadBookmarkContentRes{},
		},
	}
}
