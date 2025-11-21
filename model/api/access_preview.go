package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 访问预览
// https://open.shimo.im/docs/06API-document/interface-description/file-operation#%E8%AE%BF%E9%97%AE%E9%A2%84%E8%A7%88

type AccessPreviewParams struct {
	Auth
	FileId string
}

func NewAccessPreviewApi(params AccessPreviewParams) common.ApiConf {
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/api/cloud-files/%s/page", params.FileId),
		Method: http.MethodGet,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, nil),
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
		},
	}
}
