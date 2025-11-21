package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 读取传统文档书签内容
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E5%88%A0%E9%99%A4%E8%A1%A8%E6%A0%BC%E8%A1%8C

type ReadBookmarkContentParams struct {
	Auth
	FileId    string
	Bookmarks []string
}

type ReadBookmarkContentRespBody struct {
	RawResponse
	Data []Data `json:"data"`
}

type Data struct {
	Bookmark string `json:"bookmark"`
	Content  string `json:"content"`
}

func NewReadBookmarkContentApi(params ReadBookmarkContentParams) common.ApiConf {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/api/files/%s/documentpro/bookmark_content", params.FileId),
		Method: http.MethodGet,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Query: map[string][]string{
				"bookmarks": params.Bookmarks,
			},
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            ReadBookmarkContentRespBody{},
		},
	}
}
