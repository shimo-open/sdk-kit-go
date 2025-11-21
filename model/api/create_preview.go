package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 创建预览
// https://open.shimo.im/docs/06API-document/interface-description/file-operation#%E5%88%9B%E5%BB%BA%E9%A2%84%E8%A7%88

type CreatePreviewParams struct {
	Auth
	FileId string
}

type CreatePreviewRespBody struct {
	RawResponse
	Code    string `json:"code"`    // 创建预览结果状态码，空字符串代表创建成功， 非空代表创建失败
	Message string `json:"message"` // 创建失败错误信息
}

func NewCreatePreviewApi(params CreatePreviewParams) common.ApiConf {
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/api/cloud-files/%s/create", params.FileId),
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, nil),
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            CreatePreviewRespBody{},
		},
	}
}
