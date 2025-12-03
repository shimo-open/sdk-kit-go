package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 追加表格内容
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E8%BF%BD%E5%8A%A0%E8%A1%A8%E6%A0%BC%E5%86%85%E5%AE%B9

type AppendTableContentReq struct {
	Metadata
	FileID string
	AppendTableContentReqBody
}

type AppendTableContentReqBody struct {
	Rg       string `json:"range"`
	Resource `json:"resource"`
}

type Resource struct {
	Values [][]interface{} `json:"values"`
}

type AppendTableContentRes struct {
	rawRes
}

func NewAppendTableContentApi(cli *ehttp.Component, ss SignatureSigner, params AppendTableContentReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/sheets/values", params.FileID),
		Method: http.MethodPut,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Body:    params.AppendTableContentReqBody,
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
		},
	}
}
