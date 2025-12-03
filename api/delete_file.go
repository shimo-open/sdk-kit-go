package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 删除协同文档
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E5%88%A0%E9%99%A4%E5%8D%8F%E5%90%8C%E6%96%87%E6%A1%A3

type DeleteFileReq struct {
	Metadata
	FileID string
}

type DeleteFileRes struct {
	rawRes
}

func NewDeleteFileApi(cli *ehttp.Component, ss SignatureSigner, params DeleteFileReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
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
