package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 创建预览
// https://open.shimo.im/docs/06API-document/interface-description/file-operation#%E5%88%9B%E5%BB%BA%E9%A2%84%E8%A7%88

type CreatePreviewReq struct {
	Metadata
	FileID string
}

type CreatePreviewRes struct {
	rawRes
	Code    string `json:"code"`    // 创建预览结果状态码，空字符串代表创建成功， 非空代表创建失败
	Message string `json:"message"` // 创建失败错误信息
}

func NewCreatePreviewApi(cli *ehttp.Component, ss SignatureSigner, params CreatePreviewReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/cloud-files/%s/create", params.FileID),
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            CreatePreviewRes{},
		},
	}
}
