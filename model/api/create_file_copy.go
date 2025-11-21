package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 创建协同文档副本
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E5%88%9B%E5%BB%BA%E5%8D%8F%E5%90%8C%E6%96%87%E6%A1%A3%E5%89%AF%E6%9C%AC

type CreateFileCopyParams struct {
	Auth
	OriginFileId string
	TargetFileId string
}

type CreateFileCopyRequestBody struct {
	FileId string `json:"fileId"`
}

func NewCreateFileCopyApi(params CreateFileCopyParams) common.ApiConf {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/collab-files/%s/copy", params.OriginFileId),
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Body: CreateFileCopyRequestBody{
				FileId: params.TargetFileId,
			},
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusNoContent,
		},
	}
}
