package api

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/gotomicro/ego/client/ehttp"
)

type ImportFileToAiKnowledgeBaseReq struct {
	Metadata
	ImportFileToAiKnowledgeBaseReqBody
}

type ImportFileToAiKnowledgeBaseRes struct {
	Res *resty.Response
	apiRes
}

func (r *ImportFileToAiKnowledgeBaseRes) Response() *resty.Response {
	return r.Res
}

func (r *ImportFileToAiKnowledgeBaseRes) SetResponse(res *resty.Response) {
	r.Res = res
}

type ImportFileToAiKnowledgeBaseReqBody struct {
	KnowledgeBaseGuid string `json:"knowledgeBaseGuid"` // 接入方的知识库ID
	// file 或 url 二选一
	// 导入一个石墨协同文档时，填 file，系统会通过 fileGuid 将其导入知识库
	// 导入一个外部文档时，填 url，系统会通过 downloadUrl 将其下载并导入知识库
	ImportType string `json:"importType"`
	FileGuid   string `json:"fileGuid"`
	// 文件类型，分为两种情况
	// 1.如果importType 为 file，传石墨文档类型（document, documentPro, spreadsheet, presentation）
	// 2.如果 importType 为 url，传外部文档类型（目前只支持 pdf, rtf）
	FileType string `json:"fileType"`
	// 只有在 importType 为 "file" 时才生效，云文件的下载链接，必须保证网络通畅，让石墨服务器可以下载到
	// 注意：该地址必须能被石墨服务所在的服务器访问 示例值: https://example.com/document.pdf
	DownloadUrl string `json:"downloadUrl"`
}

type ImportFileToAiRes struct {
	Res *resty.Response
	apiRes
}

func (r *ImportFileToAiRes) Response() *resty.Response {
	return r.Res
}

func (r *ImportFileToAiRes) SetResponse(res *resty.Response) {
	r.Res = res
}

func NewImportFileToAiKnowledgeBaseApi(cli *ehttp.Component, ss SignatureSigner, params ImportFileToAiKnowledgeBaseReq) (apiConf *APIConf) {
	sign := ss.Sign(expire4m, ScopeDefault)
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

// V2版本接口

type ImportFileToAiKnowledgeBaseV2Req struct {
	Metadata
	ImportFileToAiKnowledgeBaseV2ReqBody
}

type ImportFileToAiKnowledgeBaseV2ReqBody struct {
	KnowledgeBaseGuid string `json:"knowledgeBaseGuid"` // 接入方的知识库ID
	// file 或 url 二选一
	// 导入一个石墨协同文档时，填 file，系统会通过 fileGuid 将其导入知识库
	// 导入一个外部文档时，填 url，系统会通过 downloadUrl 将其下载并导入知识库
	ImportType string `json:"importType"`
	FileGuid   string `json:"fileGuid"`
	// 文件类型，分为两种情况
	// 1.如果importType 为 file，传石墨文档类型（document, documentPro, spreadsheet, presentation）
	// 2.如果 importType 为 url，传外部文档类型（目前只支持 pdf, rtf）
	FileType string `json:"fileType"`
	// 只有在 importType 为 "file" 时才生效，云文件的下载链接，必须保证网络通畅，让石墨服务器可以下载到
	// 注意：该地址必须能被石墨服务所在的服务器访问 示例值: https://example.com/document.pdf
	DownloadUrl string `json:"downloadUrl"`
}

type ImportFileToAiV2Res struct {
	rawRes
	ImportFileToAiV2RespData
}

type ImportFileToAiV2RespData struct {
	TaskID string `json:"taskId"` // 任务ID，用于查询导入进度
}

func NewImportFileToAiKnowledgeBaseV2Api(cli *ehttp.Component, ss SignatureSigner, params ImportFileToAiKnowledgeBaseV2Req) (apiConf *APIConf) {
	sign := ss.Sign(expire4m, ScopeDefault)
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

// 导入进度查询接口

type GetImportFileToAiProgressV2Req struct {
	Metadata
	GetImportFileToAiProgressV2ReqBody
}

type GetImportFileToAiProgressV2ReqBody struct {
	TaskID string `json:"taskId"` // 任务ID
}

type GetImportFileToAiProgressV2Res struct {
	Res *resty.Response
	apiRes
	GetImportFileToAiProgressV2RespData
}

func (r *GetImportFileToAiProgressV2Res) Response() *resty.Response {
	return r.Res
}

func (r *GetImportFileToAiProgressV2Res) SetResponse(res *resty.Response) {
	r.Res = res
}

type GetImportFileToAiProgressV2RespData struct {
	TaskId   string `json:"taskId"`   // 任务ID
	Status   string `json:"status"`   // 任务状态：pending, processing, completed, failed
	Progress int    `json:"progress"` // 进度百分比 0-100
	Message  string `json:"message"`  // 状态描述信息
}

func NewGetImportFileToAiProgressV2Api(cli *ehttp.Component, ss SignatureSigner, params GetImportFileToAiProgressV2Req) (apiConf *APIConf) {
	sign := ss.Sign(expire4m, ScopeDefault)
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
