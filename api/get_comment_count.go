package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 获取文件中的评论数
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#get-comment-count

type GetCommentCountReq struct {
	Metadata
	FileID string
}

type GetCommentCountRes struct {
	rawRes
	Count int `json:"count"`
}

func NewGetCommentCountApi(cli *ehttp.Component, ss SignatureSigner, params GetCommentCountReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
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
