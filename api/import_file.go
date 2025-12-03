package api

import (
	"net/http"
	"os"

	"github.com/gotomicro/ego/client/ehttp"
)

// ImportFileReq contains parameters for importing a file.
// ImportFileReq 包含导入文件的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/file-operation#import-v1
type ImportFileReq struct {
	Metadata
	ImportFileReqBody
}

// ImportFileReqBody contains the request body for importing a file.
// ImportFileReqBody 包含导入文件的请求体。
type ImportFileReqBody struct {
	// FileID is the unique identifier for the imported file.
	// FileID 是导入文件的唯一标识符。
	FileID string `json:"fileId"`
	// Type is the file type.
	// Type 是文件类型。
	Type string `json:"type"`
	// File is the file to import (optional).
	// File 是要导入的文件（可选）。
	File *os.File `json:"file"`
	// FileUrl is the URL of the file to import (optional).
	// FileUrl 是要导入的文件 URL（可选）。
	FileUrl string `json:"fileUrl"`
	// FileName is the name of the file (optional).
	// FileName 是文件名称（可选）。
	FileName string `json:"fileName"`
}

// ImportFileRes contains the response for importing a file.
// ImportFileRes 包含导入文件的响应。
type ImportFileRes struct {
	rawRes
	// Status is the import status (non-zero indicates error).
	// Status 是导入状态（非零值表示异常）。
	Status int `json:"status"`
	// Message is the error message when import fails.
	// Message 是导入失败时的提示信息。
	Message string `json:"message"`
	// Data contains the import result data.
	// Data 包含导入结果数据。
	Data ImportFileData `json:"data"`
}

// ImportFileData contains the import result data.
// ImportFileData 包含导入结果数据。
type ImportFileData struct {
	// TaskID is the import task identifier for progress tracking.
	// TaskID 是导入任务的标识 ID，用于跟踪进度。
	TaskID string `json:"taskId"`
}

// NewImportFileApi creates a new API config for importing a file.
// NewImportFileApi 创建用于导入文件的 API 配置。
func NewImportFileApi(cli *ehttp.Component, ss SignatureSigner, params ImportFileReq) (apiConf *APIConf) {
	sign := ss.Sign(ExpireLong, ScopeDefault)
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

// NewImportV2FileApi creates a new API config for importing a file using v2 API.
// NewImportV2FileApi 创建用于使用 v2 接口导入文件的 API 配置。
func NewImportV2FileApi(cli *ehttp.Component, ss SignatureSigner, params ImportFileReq) (apiConf *APIConf) {
	sign := ss.Sign(ExpireLong, ScopeDefault)
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
