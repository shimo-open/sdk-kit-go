package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 读取传统文档书签内容
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E5%88%A0%E9%99%A4%E8%A1%A8%E6%A0%BC%E8%A1%8C

type ReadBookmarkContentReq struct {
	Metadata
	FileID    string
	Bookmarks []string
}

type ReadBookmarkContentRes struct {
	rawRes
	Data []Data `json:"data"`
}

type Data struct {
	Bookmark string `json:"bookmark"`
	Content  string `json:"content"`
}

func NewReadBookmarkContentApi(cli *ehttp.Component, ss SignatureSigner, params ReadBookmarkContentReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
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
