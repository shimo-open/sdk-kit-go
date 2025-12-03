package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// DeleteFileReq contains parameters for deleting a collaborative document.
// DeleteFileReq 包含删除协同文档的参数。
// API Doc: https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E5%88%A0%E9%99%A4%E5%8D%8F%E5%90%8C%E6%96%87%E6%A1%A3
type DeleteFileReq struct {
	Metadata
	// FileID is the unique identifier of the file to delete.
	// FileID 是要删除的文件的唯一标识符。
	FileID string
}

// DeleteFileRes contains the response for deleting a file.
// DeleteFileRes 包含删除文件的响应。
type DeleteFileRes struct {
	rawRes
}

// NewDeleteFileApi creates a new API config for deleting a file.
// NewDeleteFileApi 创建用于删除文件的 API 配置。
func NewDeleteFileApi(cli *ehttp.Component, ss SignatureSigner, params DeleteFileReq) *APIConf {
	sign := ss.Sign(ExpireLong, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s", params.FileID),
		Method: http.MethodDelete,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
		},
	}
}
