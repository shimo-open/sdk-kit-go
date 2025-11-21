package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 获取文件纯文本内容
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E6%96%87%E4%BB%B6%E7%BA%AF%E6%96%87%E6%9C%AC%E5%AD%97%E6%95%B0%E7%BB%9F%E8%AE%A1

type GetPlainTextWCParams struct {
	Auth
	FileId string
}

type GetPlainTextWCRespBody struct {
	RawResponse
	WordCount int         `json:"wordCount"` // 根据指定文件 ID 获取的石墨文件纯文本字数
	Keywords  interface{} `json:"keywords"`  // map[keyword]count，示例： {"foo":1,"bar":10}
}

func NewGetPlainTextWCApi(params GetPlainTextWCParams) common.ApiConf {
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/collab-files/%s/plain-text/wc", params.FileId),
		Method: http.MethodPost,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, nil),
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            GetPlainTextWCRespBody{},
		},
	}
}
