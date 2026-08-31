package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// GetMarkdownContentReq contains parameters for retrieving Markdown content.
type GetMarkdownContentReq struct {
	Metadata
	// FileID is the provider file identifier.
	FileID string
}

// GetMarkdownContentRes contains the Markdown content returned by the API.
type GetMarkdownContentRes struct {
	rawRes
	Content string `json:"content"`
}

// NewGetMarkdownContentApi creates the Markdown content request.
func NewGetMarkdownContentApi(cli *ehttp.Component, ss SignatureSigner, p GetMarkdownContentReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/content/%s/r2m", p.FileID),
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(p.Metadata, sign, nil),
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetMarkdownContentRes{},
		},
	}
}
