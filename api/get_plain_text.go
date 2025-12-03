package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 获取文件纯文本内容
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E8%8E%B7%E5%8F%96%E6%96%87%E4%BB%B6%E7%BA%AF%E6%96%87%E6%9C%AC%E5%86%85%E5%AE%B9

type GetPlainTextReq struct {
	Metadata
	FileID string
}

type GetPlainTextRes struct {
	rawRes
	Content string `json:"content"` // 根据指定文件 ID 获取的石墨文件纯文本内容
}

func NewGetPlainTextApi(cli *ehttp.Component, ss SignatureSigner, params GetPlainTextReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/collab-files/%s/plain-text", params.FileID),
		Method: http.MethodGet,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetPlainTextRes{},
		},
	}
}
