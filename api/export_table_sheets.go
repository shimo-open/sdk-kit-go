package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 导出应用表格为 Excel
// https://open.shimo.im/docs/06API-document/interface-description/file-operation#export-table-sheets

type ExportTableSheetsReq struct {
	Metadata
	FileID string
}

type ExportTableSheetsRes struct {
	rawRes
	Status      int    `json:"status"`      // 导出状态，非零值表示异常
	Message     string `json:"message"`     // 导出异常时的提示信息
	DownloadUrl string `json:"downloadUrl"` // .xlsx 文件下载地址
}

func NewExportTableSheetsApi(cli *ehttp.Component, ss SignatureSigner, params ExportTableSheetsReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/export/table-sheets/%s", params.FileID),
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            ExportTableSheetsRes{},
		},
	}
}
