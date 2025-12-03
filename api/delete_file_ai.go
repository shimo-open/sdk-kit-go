package api

import (
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// DeleteFileFromAiKnowledgeBaseReq contains parameters for deleting a file from AI knowledge base.
// DeleteFileFromAiKnowledgeBaseReq 包含从 AI 知识库删除文件的参数。
type DeleteFileFromAiKnowledgeBaseReq struct {
	Metadata
	DeleteFileFromAiKnowledgeBaseReqBody
}

// DeleteFileFromAiKnowledgeBaseReqBody contains the request body for deleting a file from AI knowledge base.
// DeleteFileFromAiKnowledgeBaseReqBody 包含从 AI 知识库删除文件的请求体。
type DeleteFileFromAiKnowledgeBaseReqBody struct {
	// KnowledgeBaseGuid is the knowledge base ID.
	// KnowledgeBaseGuid 是知识库 ID。
	KnowledgeBaseGuid string `json:"knowledgeBaseGuid"`
	// FileGuid is the file GUID to delete.
	// FileGuid 是要删除的文件 GUID。
	FileGuid string `json:"fileGuid"`
}

// DeleteFileFromAiKnowledgeBaseRes contains the response for deleting a file from AI knowledge base.
// DeleteFileFromAiKnowledgeBaseRes 包含从 AI 知识库删除文件的响应。
type DeleteFileFromAiKnowledgeBaseRes struct {
	rawRes
}

// NewDeleteFileFromAiKnowledgeBaseApi creates a new API config for deleting a file from AI knowledge base.
// NewDeleteFileFromAiKnowledgeBaseApi 创建用于从 AI 知识库删除文件的 API 配置。
func NewDeleteFileFromAiKnowledgeBaseApi(cli *ehttp.Component, ss SignatureSigner, params DeleteFileFromAiKnowledgeBaseReq) (apiConf *APIConf) {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	apiConf = &APIConf{
		ss:     ss,
		Client: cli,
		URL:    "/sdk/v2/api/ai/rag/delete",
		Method: http.MethodDelete,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body: DeleteFileFromAiKnowledgeBaseReqBody{
				KnowledgeBaseGuid: params.KnowledgeBaseGuid,
				FileGuid:          params.FileGuid,
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            DeleteFileFromAiKnowledgeBaseRes{},
		},
	}
	return
}
