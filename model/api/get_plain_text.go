package api

import (
	"fmt"
	"net/http"

	"github.com/shimo-open/sdk-kit-go/model/common"
)

// 获取文件纯文本内容
// https://open.shimo.im/docs/06API-document/interface-description/collaborative-editing#%E8%8E%B7%E5%8F%96%E6%96%87%E4%BB%B6%E7%BA%AF%E6%96%87%E6%9C%AC%E5%86%85%E5%AE%B9

type GetPlainTextParams struct {
	Auth
	FileId string
}

type GetPlainTextRespBody struct {
	RawResponse
	Content string `json:"content"` // 根据指定文件 ID 获取的石墨文件纯文本内容
}

func NewGetPlainTextApi(params GetPlainTextParams) common.ApiConf {
	return common.ApiConf{
		Url:    fmt.Sprintf(Prefix+"/sdk/v2/collab-files/%s/plain-text", params.FileId),
		Method: http.MethodGet,
		ReqParams: common.ApiRequestParams{
			Headers: GetHeaderParams(params.Auth, nil),
		},
		RespParams: common.ApiResponseParams{
			SuccessHttpCode: http.StatusOK,
			Body:            GetPlainTextRespBody{},
		},
	}
}
