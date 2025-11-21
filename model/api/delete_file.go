package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 删除协同文档
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E5%88%A0%E9%99%A4%E5%8D%8F%E5%90%8C%E6%96%87%E6%A1%A3

type DeleteFileParams struct {
	Auth
	FileId string
}

func NewDeleteFileApi(params DeleteFileParams) common.ApiConf {
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/api/files/%s", params.FileId),
		Method: http.MethodDelete,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, nil),
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusNoContent,
		},
	}
}
