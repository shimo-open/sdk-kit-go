package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gotomicro/ego/client/ehttp"
)

// 删除表格行
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E5%88%A0%E9%99%A4%E8%A1%A8%E6%A0%BC%E8%A1%8C

type DeleteTableRowReq struct {
	Metadata
	FileID, SheetName string
	Index, Count      int
}

type DeleteTableRowRes struct {
	rawRes
}

func NewDeleteTableRowApi(cli *ehttp.Component, ss SignatureSigner, params DeleteTableRowReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
	extra := map[string]string{
		"Content-Type": "application/json",
	}
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/api/files/%s/sheets/%s/rows/%d", params.FileID, params.SheetName, params.Index),
		Method: http.MethodDelete,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, extra),
			Query: map[string][]string{
				"count": {strconv.Itoa(params.Count)},
			},
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusNoContent,
		},
	}
}
