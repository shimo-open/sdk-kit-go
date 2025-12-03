package api

import (
	"fmt"
	"net/http"

	"github.com/gotomicro/ego/client/ehttp"
)

// 获取文件纯文本内容
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E6%96%87%E4%BB%B6%E7%BA%AF%E6%96%87%E6%9C%AC%E5%AD%97%E6%95%B0%E7%BB%9F%E8%AE%A1

type GetPlainTextWCReq struct {
	Metadata
	FileID string
}

type GetPlainTextWCRes struct {
	rawRes
	WordCount int         `json:"wordCount"` // 根据指定文件 ID 获取的石墨文件纯文本字数
	Keywords  interface{} `json:"keywords"`  // map[keyword]count，示例： {"foo":1,"bar":10}
}

func NewGetPlainTextWCApi(cli *ehttp.Component, ss SignatureSigner, params GetPlainTextWCReq) *APIConf {
	sign := ss.Sign(expire4m, ScopeDefault)
	return &APIConf{
		ss:     ss,
		Client: cli,
		URL:    fmt.Sprintf("/sdk/v2/collab-files/%s/plain-text/wc", params.FileID),
		Method: http.MethodPost,
		ReqParams: APIRequestParams{
			Headers: addHeaders(params.Metadata, sign, nil),
		},
		ResParams: APIResParams{
			SuccessHTTPCode: http.StatusOK,
			Body:            GetPlainTextWCRes{},
		},
	}
}
