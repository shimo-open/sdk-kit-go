package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// AppendMarkdownContentReq contains parameters for appending Markdown content.
type AppendMarkdownContentReq struct {
	Metadata
	// FileID is the provider file identifier.
	FileID string
	// Content is the Markdown text appended to the document.
	Content string
}

// AppendMarkdownContentReqBody contains the request body for appending Markdown content.
// AppendMarkdownContentReqBody 包含追加 Markdown 内容的请求体。
type AppendMarkdownContentReqBody struct {
	// Type is the appended content type. Markdown content uses "md".
	// Type 是追加内容类型，Markdown 内容固定传 "md"。
	Type string `json:"typ"`
	// Content is the Markdown text appended to the document.
	// Content 是追加到文档末尾的 Markdown 文本。
	Content string `json:"content"`
}

// AppendMarkdownContentRes contains the append operation result.
type AppendMarkdownContentRes struct {
	rawRes
	Status string `json:"status"`
	TID    string `json:"tid"`
}

// NewAppendMarkdownContentApi creates the append Markdown content request.
func NewAppendMarkdownContentApi(cli *ehttp.Component, ss SignatureSigner, p AppendMarkdownContentReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	body := AppendMarkdownContentReqBody{
		Type:    "md",
		Content: p.Content,
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/edit/%s/compose-custom", p.FileID),
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(p.Metadata, sign, extra),
			Body:    body,
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            AppendMarkdownContentRes{},
		},
	}
}
