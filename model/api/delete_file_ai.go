package api

import (
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

type DeleteFileFromAiKnowledgeBaseParams struct {
	Auth
	DeleteFileFromAiKnowledgeBaseReqBody
}

type DeleteFileFromAiKnowledgeBaseReqBody struct {
	KnowledgeBaseGuid string `json:"knowledgeBaseGuid"` // 接入方的知识库ID
	FileGuid          string `json:"fileGuid"`
}

type DeleteFileFromAiKnowledgeBaseRespBody struct {
	RawResponse
}

func NewDeleteFileFromAiKnowledgeBaseApi(params DeleteFileFromAiKnowledgeBaseParams) (apiConf common.ApiConf) {
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	apiConf = common.ApiConf{
		Url:    Prefix + "/sdk/v2/api/ai/rag/delete",
		Method: http.MethodDelete,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, extra),
			Body: DeleteFileFromAiKnowledgeBaseReqBody{
				KnowledgeBaseGuid: params.KnowledgeBaseGuid,
				FileGuid:          params.FileGuid,
			},
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            DeleteFileFromAiKnowledgeBaseRespBody{},
		},
	}
	return
}
