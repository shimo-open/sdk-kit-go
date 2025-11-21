package api

import (
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 获取导入进度
// https://open.shimo.im/docs/06API-document/interface-description/file-operation#import-progress-v1

type GetImportProgParams struct {
	Auth
	TaskId string
}

type GetImportProgRespBody struct {
	RawResponse
	Status  int               `json:"status"`  // 导入状态，非零值表示异常
	Message string            `json:"message"` // 导入异常时的提示信息
	Data    GetImportProgData `json:"data"`
}

type GetImportProgData struct {
	Progress int `json:"progress"` // 导入进度百分比，为 100 时表示导入完成
}

func NewGetImportProgressApi(params GetImportProgParams) common.ApiConf {
	return common.ApiConf{
		Url:    Prefix + "/sdk/v2/api/files/v1/import/progress",
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: map[string][]string{
				"X-Shimo-Signature":     {params.Signature},
				"X-Shimo-Token":         {params.Token},
				"X-Weboffice-Token":     {params.Token},
				"X-Weboffice-User-Uuid": {params.UserUuid},
			},
			Query: map[string][]string{
				"taskId": {params.TaskId},
			},
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            GetImportProgRespBody{},
		},
	}
}

func NewGetImportV2ProgressApi(params GetImportProgParams) common.ApiConf {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return common.ApiConf{
		Url:    Prefix + "/sdk/v2/api/files/v2/import/progress",
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Body: map[string]string{
				"taskId": params.TaskId,
			},
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            GetImportProgRespBody{},
		},
	}
}
