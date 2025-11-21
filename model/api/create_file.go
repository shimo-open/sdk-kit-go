package api

import (
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 创建协同文档
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#create-collab-file

type Auth struct {
	Signature string
	Token     string
	UserUuid  string
}

type CreateFileParams struct {
	Auth
	FileType   common.CollabFileType
	FileId     string
	Lang       common.Lang
	ContentKey string
}

type CreateFileRequestBody struct {
	Type       common.CollabFileType `json:"type"`
	FileId     string                `json:"fileId"`
	ContentKey string                `json:"contentKey"`
}

func NewCreateFileApi(params CreateFileParams) common.ApiConf {
	extra := map[string]string{
		"Content-Type":    "application/json",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
	}
	return common.ApiConf{
		Url:    Prefix + "/sdk/v2/api/files",
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Query:   map[string][]string{"lang": {params.Lang.String()}},
			Body: CreateFileRequestBody{
				Type:       params.FileType,
				FileId:     params.FileId,
				ContentKey: params.ContentKey,
			},
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusNoContent,
		},
	}
}
