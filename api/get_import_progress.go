package api

import (
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 获取导入进度
// https://open.shimo.im/docs/06API-document/interface-description/file-operation#import-progress-v1

type GetImportProgReq struct {
	Metadata
	TaskId string
}

type GetImportProgRes struct {
	rawRes
	Status  int               `json:"status"`  // 导入状态，非零值表示异常
	Message string            `json:"message"` // 导入异常时的提示信息
	Data    GetImportProgData `json:"data"`
}

type GetImportProgData struct {
	Progress int `json:"progress"` // 导入进度百分比，为 100 时表示导入完成
}

func NewGetImportProgressApi(cli *ehttp.Component, ss SignatureSigner, params GetImportProgReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/files/v1/import/progress",
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
			Query: map[string][]string{
				"taskId": {params.TaskId},
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetImportProgRes{},
		},
	}
}

func NewGetImportV2ProgressApi(cli *ehttp.Component, ss SignatureSigner, params GetImportProgReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/files/v2/import/progress",
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body: map[string]string{
				"taskId": params.TaskId,
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetImportProgRes{},
		},
	}
}
