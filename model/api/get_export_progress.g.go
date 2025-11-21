package api

import (
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 获取导出进度
// https://open.shimo.im/docs/06API-document/interface-description/file-operation#export-progress-v1

type GetExportProgParams struct {
	Auth
	TaskId string
}

type GetExportProgRespBody struct {
	RawResponse
	Status  int                   `json:"status"`  // 导出状态，非零值表示异常
	Message string                `json:"message"` // 导出异常时的提示信息
	Data    GetExportProgRespData `json:"data"`
}

type GetExportProgRespData struct {
	Progress    int    `json:"progress"`    // 导出进度百分比，为 100 时表示导入完成
	DownloadUrl string `json:"downloadUrl"` // 导出文件的下载地址
}

func NewGetExportProgressApi(params GetExportProgParams) common.ApiConf {
	return common.ApiConf{
		Url:    Prefix + "/sdk/v2/api/files/v1/export/progress",
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, nil),
			Query: map[string][]string{
				"taskId": {params.TaskId},
			},
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            GetExportProgRespBody{},
		},
	}
}
