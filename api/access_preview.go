package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// AccessPreviewReq contains parameters for accessing a file preview.
// AccessPreviewReq 包含访问文件预览的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/file-operation#%E8%AE%BF%E9%97%AE%E9%A2%84%E8%A7%88
type AccessPreviewReq struct {
	Metadata
	// FileID is the unique identifier of the file.
	// FileID 是文件的唯一标识符。
	FileID string
}

// AccessPreviewRes contains the response for accessing a file preview.
// AccessPreviewRes 包含访问文件预览的响应。
type AccessPreviewRes struct {
	rawRes
}

// NewAccessPreviewApi creates a new API config for accessing a file preview.
// NewAccessPreviewApi 创建用于访问文件预览的 API 配置。
func NewAccessPreviewApi(cli *ehttp.Component, ss SignatureSigner, params AccessPreviewReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	md := Metadata{
		ShimoToken:        params.ShimoToken,
		WebofficeUserUuid: params.WebofficeUserUuid,
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/cloud-files/%s/page", params.FileID),
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(md, sign, nil),
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
		},
	}
}
