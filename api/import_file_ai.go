package api

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/gotomicro/ego/client/ehttp"
)

// ImportFileToAiKnowledgeBaseReq contains parameters for importing a file to AI knowledge base.
// ImportFileToAiKnowledgeBaseReq 包含导入文件到 AI 知识库的参数。
type ImportFileToAiKnowledgeBaseReq struct {
	Metadata
	ImportFileToAiKnowledgeBaseReqBody
}

// ImportFileToAiKnowledgeBaseRes contains the response for importing a file to AI knowledge base.
// ImportFileToAiKnowledgeBaseRes 包含导入文件到 AI 知识库的响应。
type ImportFileToAiKnowledgeBaseRes struct {
	Res *resty.Response
	apiRes
}

// Response returns the raw HTTP response.
// Response 返回原始 HTTP 响应。
func (r *ImportFileToAiKnowledgeBaseRes) Response() *resty.Response {
	return r.Res
}

// SetResponse sets the raw HTTP response.
// SetResponse 设置原始 HTTP 响应。
func (r *ImportFileToAiKnowledgeBaseRes) SetResponse(res *resty.Response) {
	r.Res = res
}

// ImportFileToAiKnowledgeBaseReqBody contains the request body for importing a file to AI knowledge base.
// ImportFileToAiKnowledgeBaseReqBody 包含导入文件到 AI 知识库的请求体。
type ImportFileToAiKnowledgeBaseReqBody struct {
	// KnowledgeBaseGuid is the knowledge base ID.
	// KnowledgeBaseGuid 是知识库 ID。
	KnowledgeBaseGuid string `json:"knowledgeBaseGuid"`
	// ImportType is the import type ("file" or "url").
	// ImportType 是导入类型（"file" 或 "url"）。
	ImportType string `json:"importType"`
	// FileGuid is the file GUID (used when ImportType is "file").
	// FileGuid 是文件 GUID（当 ImportType 为 "file" 时使用）。
	FileGuid string `json:"fileGuid"`
	// FileType is the file type (document, documentPro, spreadsheet, presentation for "file"; pdf, rtf for "url").
	// FileType 是文件类型（"file" 时为 document, documentPro, spreadsheet, presentation；"url" 时为 pdf, rtf）。
	FileType string `json:"fileType"`
	// DownloadUrl is the file download URL (used when ImportType is "url").
	// DownloadUrl 是文件下载 URL（当 ImportType 为 "url" 时使用）。
	DownloadUrl string `json:"downloadUrl"`
}

// ImportFileToAiRes contains the response for importing a file to AI.
// ImportFileToAiRes 包含导入文件到 AI 的响应。
type ImportFileToAiRes struct {
	Res *resty.Response
	apiRes
}

// Response returns the raw HTTP response.
// Response 返回原始 HTTP 响应。
func (r *ImportFileToAiRes) Response() *resty.Response {
	return r.Res
}

// SetResponse sets the raw HTTP response.
// SetResponse 设置原始 HTTP 响应。
func (r *ImportFileToAiRes) SetResponse(res *resty.Response) {
	r.Res = res
}

// NewImportFileToAiKnowledgeBaseApi creates a new API config for importing a file to AI knowledge base.
// NewImportFileToAiKnowledgeBaseApi 创建用于导入文件到 AI 知识库的 API 配置。
func NewImportFileToAiKnowledgeBaseApi(cli *ehttp.Component, ss SignatureSigner, params ImportFileToAiKnowledgeBaseReq) (apiConf *APIConf) {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	apiConf = &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/ai/rag/import",
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body: ImportFileToAiKnowledgeBaseReqBody{
				KnowledgeBaseGuid: params.KnowledgeBaseGuid,
				ImportType:        params.ImportType,
				FileGuid:          params.FileGuid,
				FileType:          params.FileType,
				DownloadUrl:       params.DownloadUrl,
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
			Body:            ImportFileToAiRes{},
		},
	}
	return
}

// ImportFileToAiKnowledgeBaseV2Req contains parameters for importing a file to AI knowledge base using v2 API.
// ImportFileToAiKnowledgeBaseV2Req 包含使用 v2 接口导入文件到 AI 知识库的参数。
type ImportFileToAiKnowledgeBaseV2Req struct {
	Metadata
	ImportFileToAiKnowledgeBaseV2ReqBody
}

// ImportFileToAiKnowledgeBaseV2ReqBody contains the request body for importing a file to AI knowledge base using v2 API.
// ImportFileToAiKnowledgeBaseV2ReqBody 包含使用 v2 接口导入文件到 AI 知识库的请求体。
type ImportFileToAiKnowledgeBaseV2ReqBody struct {
	// KnowledgeBaseGuid is the knowledge base ID.
	// KnowledgeBaseGuid 是知识库 ID。
	KnowledgeBaseGuid string `json:"knowledgeBaseGuid"`
	// ImportType is the import type ("file" or "url").
	// ImportType 是导入类型（"file" 或 "url"）。
	ImportType string `json:"importType"`
	// FileGuid is the file GUID (used when ImportType is "file").
	// FileGuid 是文件 GUID（当 ImportType 为 "file" 时使用）。
	FileGuid string `json:"fileGuid"`
	// FileType is the file type (document, documentPro, spreadsheet, presentation for "file"; pdf, rtf for "url").
	// FileType 是文件类型（"file" 时为 document, documentPro, spreadsheet, presentation；"url" 时为 pdf, rtf）。
	FileType string `json:"fileType"`
	// DownloadUrl is the file download URL (used when ImportType is "url").
	// DownloadUrl 是文件下载 URL（当 ImportType 为 "url" 时使用）。
	DownloadUrl string `json:"downloadUrl"`
}

// ImportFileToAiV2Res contains the response for importing a file to AI using v2 API.
// ImportFileToAiV2Res 包含使用 v2 接口导入文件到 AI 的响应。
type ImportFileToAiV2Res struct {
	rawRes
	ImportFileToAiV2RespData
}

// ImportFileToAiV2RespData contains the response data for importing a file to AI using v2 API.
// ImportFileToAiV2RespData 包含使用 v2 接口导入文件到 AI 的响应数据。
type ImportFileToAiV2RespData struct {
	// TaskID is the task ID for progress tracking.
	// TaskID 是用于查询导入进度的任务 ID。
	TaskID string `json:"taskId"`
}

// NewImportFileToAiKnowledgeBaseV2Api creates a new API config for importing a file to AI knowledge base using v2 API.
// NewImportFileToAiKnowledgeBaseV2Api 创建用于使用 v2 接口导入文件到 AI 知识库的 API 配置。
func NewImportFileToAiKnowledgeBaseV2Api(cli *ehttp.Component, ss SignatureSigner, params ImportFileToAiKnowledgeBaseV2Req) (apiConf *APIConf) {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	apiConf = &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/ai/rag/v2/import",
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body: ImportFileToAiKnowledgeBaseV2ReqBody{
				KnowledgeBaseGuid: params.KnowledgeBaseGuid,
				ImportType:        params.ImportType,
				FileGuid:          params.FileGuid,
				FileType:          params.FileType,
				DownloadUrl:       params.DownloadUrl,
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            ImportFileToAiV2Res{},
		},
	}
	return
}

// GetImportFileToAiProgressV2Req contains parameters for getting import progress to AI knowledge base using v2 API.
// GetImportFileToAiProgressV2Req 包含使用 v2 接口获取导入到 AI 知识库进度的参数。
type GetImportFileToAiProgressV2Req struct {
	Metadata
	GetImportFileToAiProgressV2ReqBody
}

// GetImportFileToAiProgressV2ReqBody contains the request body for getting import progress.
// GetImportFileToAiProgressV2ReqBody 包含获取导入进度的请求体。
type GetImportFileToAiProgressV2ReqBody struct {
	// TaskID is the task ID.
	// TaskID 是任务 ID。
	TaskID string `json:"taskId"`
}

// GetImportFileToAiProgressV2Res contains the response for getting import progress to AI knowledge base using v2 API.
// GetImportFileToAiProgressV2Res 包含使用 v2 接口获取导入到 AI 知识库进度的响应。
type GetImportFileToAiProgressV2Res struct {
	Res *resty.Response
	apiRes
	GetImportFileToAiProgressV2RespData
}

// Response returns the raw HTTP response.
// Response 返回原始 HTTP 响应。
func (r *GetImportFileToAiProgressV2Res) Response() *resty.Response {
	return r.Res
}

// SetResponse sets the raw HTTP response.
// SetResponse 设置原始 HTTP 响应。
func (r *GetImportFileToAiProgressV2Res) SetResponse(res *resty.Response) {
	r.Res = res
}

// GetImportFileToAiProgressV2RespData contains the response data for getting import progress.
// GetImportFileToAiProgressV2RespData 包含获取导入进度的响应数据。
type GetImportFileToAiProgressV2RespData struct {
	// TaskId is the task ID.
	// TaskId 是任务 ID。
	TaskId string `json:"taskId"`
	// Status is the task status (pending, processing, completed, failed).
	// Status 是任务状态（pending, processing, completed, failed）。
	Status string `json:"status"`
	// Progress is the progress percentage (0-100).
	// Progress 是进度百分比（0-100）。
	Progress int `json:"progress"`
	// Message is the status description.
	// Message 是状态描述信息。
	Message string `json:"message"`
}

// NewGetImportFileToAiProgressV2Api creates a new API config for getting import progress to AI knowledge base using v2 API.
// NewGetImportFileToAiProgressV2Api 创建用于使用 v2 接口获取导入到 AI 知识库进度的 API 配置。
func NewGetImportFileToAiProgressV2Api(cli *ehttp.Component, ss SignatureSigner, params GetImportFileToAiProgressV2Req) (apiConf *APIConf) {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	apiConf = &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/ai/rag/v2/import/progress",
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body: GetImportFileToAiProgressV2ReqBody{
				TaskID: params.TaskID,
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetImportFileToAiProgressV2Res{},
		},
	}
	return
}
