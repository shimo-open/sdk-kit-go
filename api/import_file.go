package api

import (
	"net/http"
	"os"

	"github.com/gotomicro/ego/client/ehttp"
)

// 导入文件
// https://open.shimo.im/docs/06API-document/interface-description/file-operation#import-v1

type ImportFileReq struct {
	Metadata
	ImportFileReqBody
}

type ImportFileReqBody struct {
	FileID   string   `json:"fileId"`
	Type     string   `json:"type"`
	File     *os.File `json:"file"`
	FileUrl  string   `json:"fileUrl"`  // 非必需
	FileName string   `json:"fileName"` // 非必需
}

type ImportFileRes struct {
	rawRes
	Status  int            `json:"status"`  // 导入状态，非零值表示异常
	Message string         `json:"message"` // 导入异常时的提示信息
	Data    ImportFileData `json:"data"`
}

type ImportFileData struct {
	TaskID string `json:"taskId"` // 导入任务的标识 ID，调用导入进度接口时，请带上该参数。导入失败时请提供此 ID 用于调试
}

func NewImportFileApi(cli *ehttp.Component, ss SignatureSigner, params ImportFileReq) (apiConf *APIConf) {
	sign := ss.Sign(expire4m, ScopeDefault)
	apiConf = &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/files/v1/import",
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
			Form: map[string][]string{
				"fileId":   {params.FileID},
				"type":     {params.Type},
				"fileUrl":  {params.FileUrl},
				"fileName": {params.FileName},
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            ImportFileRes{},
		},
	}
	return
}

func NewImportV2FileApi(cli *ehttp.Component, ss SignatureSigner, params ImportFileReq) (apiConf *APIConf) {
	sign := ss.Sign(expire4m, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	apiConf = &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/files/v2/import",
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body: map[string]string{
				"fileId":   params.FileID,
				"type":     params.Type,
				"fileUrl":  params.FileUrl,
				"fileName": params.FileName,
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            ImportFileRes{},
		},
	}
	return
}
