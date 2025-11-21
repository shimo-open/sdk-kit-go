package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 获取文件中的评论数
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#get-comment-count

type GetCommentCountParams struct {
	Auth
	FileId string
}

type GetCommentCountRespBody struct {
	RawResponse
	Count int `json:"count"`
}

func NewGetCommentCountApi(params GetCommentCountParams) common.ApiConf {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/collab-files/%s/comment-count", params.FileId),
		Method: http.MethodGet,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            GetCommentCountRespBody{},
		},
	}
}
