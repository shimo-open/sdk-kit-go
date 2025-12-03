package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 获取表格内容
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#get-table-content

type GetTableContentReq struct {
	Metadata
	FileID string
	Rg     string
}

type GetTableContentRes struct {
	rawRes
	Values [][]interface{} `json:"values"`
	Lag    int             `json:"lag"`
}

func NewGetTableContentApi(cli *ehttp.Component, ss SignatureSigner, params GetTableContentReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/sheets/values", params.FileID),
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Query: map[string][]string{
				"range": {params.Rg},
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetTableContentRes{},
		},
	}
}
