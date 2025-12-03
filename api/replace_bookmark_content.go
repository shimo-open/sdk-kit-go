package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 替换传统文档书签内容
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#replace-bookmark

type RepBookmarkContentReq struct {
	Metadata
	FileID string
	RepBookmarkContentReqBody
}

type RepBookmarkContentReqBody struct {
	Replacements []Replacement `json:"replacements"`
}

type Replacement struct {
	Bookmark string `json:"bookmark"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}

type RepBookmarkContentRes struct {
	rawRes
}

func NewReplaceBookmarkContentApi(cli *ehttp.Component, ss SignatureSigner, params RepBookmarkContentReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
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
