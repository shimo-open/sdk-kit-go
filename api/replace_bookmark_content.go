package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// RepBookmarkContentReq contains parameters for replacing bookmark content in a traditional document.
// RepBookmarkContentReq 包含替换传统文档书签内容的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#replace-bookmark
type RepBookmarkContentReq struct {
	Metadata
	// FileID is the unique identifier of the file.
	// FileID 是文件的唯一标识符。
	FileID string
	RepBookmarkContentReqBody
}

// RepBookmarkContentReqBody contains the request body for replacing bookmark content.
// RepBookmarkContentReqBody 包含替换书签内容的请求体。
type RepBookmarkContentReqBody struct {
	// Replacements is the list of bookmark replacements.
	// Replacements 是书签替换列表。
	Replacements []Replacement `json:"replacements"`
}

// Replacement represents a single bookmark replacement.
// Replacement 表示单个书签替换。
type Replacement struct {
	// Bookmark is the bookmark name.
	// Bookmark 是书签名称。
	Bookmark string `json:"bookmark"`
	// Type is the replacement type (e.g., "text").
	// Type 是替换类型（例如："text"）。
	Type string `json:"type"`
	// Value is the replacement value.
	// Value 是替换值。
	Value string `json:"value"`
}

// RepBookmarkContentRes contains the response for replacing bookmark content.
// RepBookmarkContentRes 包含替换书签内容的响应。
type RepBookmarkContentRes struct {
	rawRes
}

// NewReplaceBookmarkContentApi creates a new API config for replacing bookmark content.
// NewReplaceBookmarkContentApi 创建用于替换书签内容的 API 配置。
func NewReplaceBookmarkContentApi(cli *ehttp.Component, ss SignatureSigner, params RepBookmarkContentReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/documentpro/bookmark_content", params.FileID),
		Method: http.MethodPut,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body:    params.RepBookmarkContentReqBody,
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
		},
	}
}
