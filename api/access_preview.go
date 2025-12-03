package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 访问预览
// https://open.shimo.im/docs/06API-document/interface-description/file-operation#%E8%AE%BF%E9%97%AE%E9%A2%84%E8%A7%88

type AccessPreviewReq struct {
	Metadata
	FileID string
}

type AccessPreviewRes struct {
	rawRes
}

func NewAccessPreviewApi(cli *ehttp.Component, ss SignatureSigner, params AccessPreviewReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
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
