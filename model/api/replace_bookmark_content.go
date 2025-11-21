package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 替换传统文档书签内容
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#replace-bookmark

type RepBookmarkContentParams struct {
	Auth
	FileId string
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

func NewReplaceBookmarkContentApi(params RepBookmarkContentParams) common.ApiConf {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/api/files/%s/documentpro/bookmark_content", params.FileId),
		Method: http.MethodPut,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Body:    params.RepBookmarkContentReqBody,
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusNoContent,
		},
	}
}
