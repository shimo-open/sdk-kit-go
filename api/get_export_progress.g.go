package api

import (
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 获取导出进度
// https://open.shimo.im/docs/06API-document/interface-description/file-operation#export-progress-v1

type GetExportProgReq struct {
	Metadata
	TaskId string
}

type GetExportProgRes struct {
	rawRes
	Status  int                   `json:"status"`  // 导出状态，非零值表示异常
	Message string                `json:"message"` // 导出异常时的提示信息
	Data    GetExportProgRespData `json:"data"`
}

type GetExportProgRespData struct {
	Progress    int    `json:"progress"`    // 导出进度百分比，为 100 时表示导入完成
	DownloadUrl string `json:"downloadUrl"` // 导出文件的下载地址
}

func NewGetExportProgressApi(cli *ehttp.Component, ss SignatureSigner, params GetExportProgReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/files/v1/export/progress",
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
			Query: map[string][]string{
				"taskId": {params.TaskId},
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetExportProgRes{},
		},
	}
}
