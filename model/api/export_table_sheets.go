package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 导出应用表格为 Excel
// https://open.shimo.im/docs/06API-document/interface-description/file-operation#export-table-sheets

type ExportTableSheetsParams struct {
	Auth
	FileId string
}

type ExportTableSheetsRespBody struct {
	RawResponse
	Status      int    `json:"status"`      // 导出状态，非零值表示异常
	Message     string `json:"message"`     // 导出异常时的提示信息
	DownloadUrl string `json:"downloadUrl"` // .xlsx 文件下载地址
}

func NewExportTableSheetsApi(params ExportTableSheetsParams) common.ApiConf {
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/api/files/export/table-sheets/%s", params.FileId),
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, nil),
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            ExportTableSheetsRespBody{},
		},
	}
}
