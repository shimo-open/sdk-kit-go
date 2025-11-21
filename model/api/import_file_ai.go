package api

import (
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

type ImportFileToAiKnowledgeBaseParams struct {
	Auth
	ImportFileToAiKnowledgeBaseReqBody
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

type ImportFileToAiRespBody struct {
	RawResponse
}

func NewImportFileToAiKnowledgeBaseApi(params ImportFileToAiKnowledgeBaseParams) (apiConf common.ApiConf) {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	apiConf = common.ApiConf{
		Url:    Prefix + "/sdk/v2/api/ai/rag/import",
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Body: ImportFileToAiKnowledgeBaseReqBody{
				KnowledgeBaseGuid: params.KnowledgeBaseGuid,
				ImportType:        params.ImportType,
				FileGuid:          params.FileGuid,
				FileType:          params.FileType,
				DownloadUrl:       params.DownloadUrl,
			},
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusNoContent,
			Body:            ImportFileToAiRespBody{},
		},
	}
	return
}

// V2版本接口

type ImportFileToAiKnowledgeBaseV2Params struct {
	Auth
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

type ImportFileToAiV2RespBody struct {
	RawResponse
	ImportFileToAiV2RespData
}

type ImportFileToAiV2RespData struct {
	TaskId string `json:"taskId"` // 任务ID，用于查询导入进度
}

func NewImportFileToAiKnowledgeBaseV2Api(params ImportFileToAiKnowledgeBaseV2Params) (apiConf common.ApiConf) {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	apiConf = common.ApiConf{
		Url:    Prefix + "/sdk/v2/api/ai/rag/v2/import",
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Body: ImportFileToAiKnowledgeBaseV2ReqBody{
				KnowledgeBaseGuid: params.KnowledgeBaseGuid,
				ImportType:        params.ImportType,
				FileGuid:          params.FileGuid,
				FileType:          params.FileType,
				DownloadUrl:       params.DownloadUrl,
			},
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            ImportFileToAiV2RespBody{},
		},
	}
	return
}

// 导入进度查询接口

type GetImportFileToAiProgressV2Params struct {
	Auth
	GetImportFileToAiProgressV2ReqBody
}

type GetImportFileToAiProgressV2ReqBody struct {
	TaskId string `json:"taskId"` // 任务ID
}

type GetImportFileToAiProgressV2RespBody struct {
	RawResponse
	GetImportFileToAiProgressV2RespData
}

type GetImportFileToAiProgressV2RespData struct {
	TaskId   string `json:"taskId"`   // 任务ID
	Status   string `json:"status"`   // 任务状态：pending, processing, completed, failed
	Progress int    `json:"progress"` // 进度百分比 0-100
	Message  string `json:"message"`  // 状态描述信息
}

func NewGetImportFileToAiProgressV2Api(params GetImportFileToAiProgressV2Params) (apiConf common.ApiConf) {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	apiConf = common.ApiConf{
		Url:    Prefix + "/sdk/v2/api/ai/rag/v2/import/progress",
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Body: GetImportFileToAiProgressV2ReqBody{
				TaskId: params.TaskId,
			},
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            GetImportFileToAiProgressV2RespBody{},
		},
	}
	return
}
