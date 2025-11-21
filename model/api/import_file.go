package api

import (
	"net/http"
	"os"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 导入文件
// https://open.shimo.im/docs/06API-document/interface-description/file-operation#import-v1

type ImportFileParams struct {
	Auth
	ImportFileReqBody
}

type ImportFileReqBody struct {
	FileId   string   `json:"fileId"`
	Type     string   `json:"type"`
	File     *os.File `json:"file"`
	FileUrl  string   `json:"fileUrl"`  // 非必需
	FileName string   `json:"fileName"` // 非必需
}

type ImportFileRespBody struct {
	RawResponse
	Status  int            `json:"status"`  // 导入状态，非零值表示异常
	Message string         `json:"message"` // 导入异常时的提示信息
	Data    ImportFileData `json:"data"`
}
type ImportFileData struct {
	TaskId string `json:"taskId"` // 导入任务的标识 ID，调用导入进度接口时，请带上该参数。导入失败时请提供此 ID 用于调试
}

func NewImportFileApi(params ImportFileParams) (apiConf common.ApiConf, err error) {
	apiConf = common.ApiConf{
		Url:    Prefix + "/sdk/v2/api/files/v1/import",
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, nil),
			Form: map[string][]string{
				"fileId":   {params.FileId},
				"type":     {params.Type},
				"fileUrl":  {params.FileUrl},
				"fileName": {params.FileName},
			},
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            ImportFileRespBody{},
		},
	}
	return
}

func NewImportV2FileApi(params ImportFileParams) (apiConf common.ApiConf, err error) {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	apiConf = common.ApiConf{
		Url:    Prefix + "/sdk/v2/api/files/v2/import",
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Body: map[string]string{
				"fileId":   params.FileId,
				"type":     params.Type,
				"fileUrl":  params.FileUrl,
				"fileName": params.FileName,
			},
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            ImportFileRespBody{},
		},
	}
	return
}
