package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 导出文件
// https://open.shimo.im/docs/06API-document/interface-description/file-operation#export-v1

type ExportFileParams struct {
	Auth
	FileId, Type string
}

type ExportFileRespBody struct {
	RawResponse
	Status  int            `json:"status"` // 导出状态，非零值表示异常
	Data    ExportFileData `json:"data"`
	Message string         `json:"message"` // 导出异常时的提示信息
}

type ExportFileData struct {
	TaskId string `json:"taskId"`
}

func NewExportFileApi(params ExportFileParams) common.ApiConf {
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/api/files/v1/export/%s", params.FileId),
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, nil),
			Body: map[string]interface{}{
				"type": params.Type,
			},
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            ExportFileRespBody{},
		},
	}
}
