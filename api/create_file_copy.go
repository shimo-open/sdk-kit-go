package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 创建协同文档副本
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E5%88%9B%E5%BB%BA%E5%8D%8F%E5%90%8C%E6%96%87%E6%A1%A3%E5%89%AF%E6%9C%AC

type CreateFileCopyReq struct {
	Metadata
	OriginFileID string
	TargetFileID string
}

type CreateFileCopyRequestBody struct {
	FileID string `json:"fileId"`
}

type CreateFileCopyRes struct {
	rawRes
}

func NewCreateFileCopyApi(cli *ehttp.Component, ss SignatureSigner, params CreateFileCopyReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/collab-files/%s/copy", params.OriginFileID),
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body: CreateFileCopyRequestBody{
				FileID: params.TargetFileID,
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
		},
	}
}
