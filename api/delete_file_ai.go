package api

import (
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

type DeleteFileFromAiKnowledgeBaseReq struct {
	Metadata
	DeleteFileFromAiKnowledgeBaseReqBody
}

type DeleteFileFromAiKnowledgeBaseReqBody struct {
	KnowledgeBaseGuid string `json:"knowledgeBaseGuid"` // 接入方的知识库ID
	FileGuid          string `json:"fileGuid"`
}

type DeleteFileFromAiKnowledgeBaseRes struct {
	rawRes
}

func NewDeleteFileFromAiKnowledgeBaseApi(cli *ehttp.Component, ss SignatureSigner, params DeleteFileFromAiKnowledgeBaseReq) (apiConf *APIConf) {
	sign := ss.Sign(expire4m, ScopeDefault)
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
