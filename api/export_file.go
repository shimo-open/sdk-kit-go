package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 导出文件
// https://open.shimo.im/docs/06API-document/interface-description/file-operation#export-v1

type ExportFileReq struct {
	Metadata
	FileID, Type string
}

type ExportFileRes struct {
	rawRes
	Status  int            `json:"status"` // 导出状态，非零值表示异常
	Data    ExportFileData `json:"data"`
	Message string         `json:"message"` // 导出异常时的提示信息
}

type ExportFileData struct {
	TaskID string `json:"taskId"`
}

func NewExportFileApi(cli *ehttp.Component, ss SignatureSigner, params ExportFileReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/v1/export/%s", params.FileID),
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
			Body: map[string]interface{}{
				"type": params.Type,
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            ExportFileRes{},
		},
	}
}
